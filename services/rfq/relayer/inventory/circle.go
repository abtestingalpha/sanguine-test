package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	messagetransmitter "github.com/synapsecns/sanguine/services/cctp-relayer/contracts/messagetransmitter"
	tokenmessenger "github.com/synapsecns/sanguine/services/cctp-relayer/contracts/tokenmessenger"
	cctpRelay "github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type rebalanceManagerCircleCCTP struct {
	// cfg is the config
	cfg relconfig.Config
	// handler is the metrics handler
	handler metrics.Handler
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// boundTokenMessengers is the map of TokenMessenger contracts (used for rebalancing)
	boundTokenMessengers map[int]*tokenmessenger.TokenMessenger
	// boundMessageTransmitters is the map of MessageTransmitter contracts (used for rebalancing)
	boundMessageTransmitters map[int]*messagetransmitter.MessageTransmitter
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// chainListeners is the map of chain listeners for CCTP events
	chainListeners map[int]listener.ContractListener
	// db is the database
	db reldb.Service
}

func newRebalanceManagerCircleCCTP(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerCircleCCTP {
	return &rebalanceManagerCircleCCTP{
		cfg:                      cfg,
		handler:                  handler,
		chainClient:              chainClient,
		txSubmitter:              txSubmitter,
		boundTokenMessengers:     make(map[int]*tokenmessenger.TokenMessenger),
		boundMessageTransmitters: make(map[int]*messagetransmitter.MessageTransmitter),
		relayerAddress:           relayerAddress,
		chainListeners:           make(map[int]listener.ContractListener),
		db:                       db,
	}
}

func (c *rebalanceManagerCircleCCTP) Start(ctx context.Context) (err error) {
	err = c.initContracts(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize contracts: %w", err)
	}

	err = c.initListeners(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize listeners: %w", err)
	}

	g, _ := errgroup.WithContext(ctx)
	for cid := range c.cfg.Chains {
		// capture func literal
		chainID := cid
		g.Go(func() error {
			return c.listen(ctx, chainID)
		})
	}

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("error listening to contract: %w", err)
	}
	return nil
}

func (c *rebalanceManagerCircleCCTP) initContracts(ctx context.Context) (err error) {
	for chainID := range c.cfg.Chains {
		contractAddr, err := c.cfg.GetTokenMessengerAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		messengerContract, err := tokenmessenger.NewTokenMessenger(common.HexToAddress(contractAddr), chainClient)
		if err != nil {
			return fmt.Errorf("could not get cctp: %w", err)
		}
		transmitterAddr, err := messengerContract.LocalMessageTransmitter(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get message transmitter")
		}
		transmitterContract, err := messagetransmitter.NewMessageTransmitter(transmitterAddr, chainClient)
		if err != nil {
			return fmt.Errorf("could not get message transmitter: %w", err)
		}
		c.boundTokenMessengers[chainID] = messengerContract
		c.boundMessageTransmitters[chainID] = transmitterContract
	}
	return nil
}

func (c *rebalanceManagerCircleCCTP) initListeners(ctx context.Context) (err error) {
	for chainID := range c.cfg.GetChains() {
		cctpAddr, err := c.cfg.GetTokenMessengerAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		initialBlock, err := c.cfg.GetCCTPStartBlock(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp start block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, c.db, common.HexToAddress(cctpAddr), initialBlock, c.handler)
		if err != nil {
			return fmt.Errorf("could not get chain listener: %w", err)
		}
		c.chainListeners[chainID] = chainListener
	}
	return nil
}

func (c *rebalanceManagerCircleCCTP) Execute(parentCtx context.Context, rebalance *RebalanceData) (err error) {
	contract, ok := c.boundTokenMessengers[rebalance.OriginMetadata.ChainID]
	if !ok {
		return fmt.Errorf("could not find cctp contract for chain %d", rebalance.OriginMetadata.ChainID)
	}
	ctx, span := c.handler.Tracer().Start(parentCtx, "rebalance.Execute", trace.WithAttributes(
		attribute.Int("rebalance_origin", rebalance.OriginMetadata.ChainID),
		attribute.Int("rebalance_dest", rebalance.DestMetadata.ChainID),
		attribute.String("rebalance_amount", rebalance.Amount.String()),
		attribute.String("burn_token_addr", rebalance.OriginMetadata.Addr.String()),
		attribute.String("burn_token_name", rebalance.OriginMetadata.Name),
		attribute.String("burn_token_balance", rebalance.OriginMetadata.Balance.String()),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	destChainID := rebalance.DestMetadata.ChainID
	destDomain, err := cctpRelay.ChainIDToCircleDomain(uint32(destChainID), cctpRelay.IsTestnetChainID(uint32(destChainID)))
	if err != nil {
		return fmt.Errorf("could not convert chain ID to domain: %w", err)
	}

	// convert our address to bytes32
	addrBytes32 := cctpRelay.AddressToBytes32(c.relayerAddress)

	span.SetAttributes(
		attribute.Int("dest_domain", int(destDomain)),
		attribute.String("addr_bytes32", hexutil.Encode(addrBytes32[:])),
		attribute.String("relayer_address", c.relayerAddress.String()),
	)

	// perform rebalance by calling depositForBurn()
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		tx, err = contract.DepositForBurn(
			transactor,
			rebalance.Amount,
			destDomain,
			addrBytes32,
			rebalance.OriginMetadata.Addr,
		)
		if err != nil {
			return nil, fmt.Errorf("could not send circle token: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit CCTP rebalance: %w", err)
	}

	// store the rebalance in the db
	model := reldb.Rebalance{
		Origin:       uint64(rebalance.OriginMetadata.ChainID),
		Destination:  uint64(rebalance.DestMetadata.ChainID),
		OriginAmount: rebalance.Amount,
		Status:       reldb.RebalanceInitiated,
	}
	err = c.db.StoreRebalance(ctx, model)
	if err != nil {
		return fmt.Errorf("could not store rebalance: %w", err)
	}
	return nil
}

// nolint:cyclop
func (c *rebalanceManagerCircleCCTP) listen(parentCtx context.Context, chainID int) (err error) {
	listener, ok := c.chainListeners[chainID]
	if !ok {
		return fmt.Errorf("could not find listener for chain %d", chainID)
	}
	ethClient, err := c.chainClient.GetClient(parentCtx, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}

	err = listener.Listen(parentCtx, func(parentCtx context.Context, log types.Log) (err error) {
		ctx, span := c.handler.Tracer().Start(parentCtx, "rebalance.Listen", trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID),
		))
		defer func(err error) {
			metrics.EndSpanWithErr(span, err)
		}(err)

		switch log.Topics[0] {
		case tokenmessenger.DepositForBurnTopic:
			err = c.handleDepositForBurn(ctx, log, chainID, ethClient)
			if err != nil {
				return fmt.Errorf("could not handle depositForBurn: %w", err)
			}
			return nil
		case messagetransmitter.MessageReceivedTopic:
			err = c.handleMessageReceived(ctx, log, chainID, ethClient)
			if err != nil {
				return fmt.Errorf("could not handle message received: %w", err)
			}
			return nil
		default:
			logger.Warnf("unknown event %s", log.Topics[0])
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen to contract: %w", err)
	}
	return nil
}

func (c *rebalanceManagerCircleCCTP) handleDepositForBurn(ctx context.Context, log types.Log, chainID int, ethClient client.EVM) (err error) {
	ctx, span := c.handler.Tracer().Start(ctx, "rebalance.handleDepositForBurn", trace.WithAttributes(
		attribute.Int(metrics.ChainID, chainID),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	messengerAddr := common.HexToAddress(c.cfg.Chains[chainID].TokenMessengerAddress)
	parser, err := tokenmessenger.NewTokenMessengerFilterer(messengerAddr, ethClient)
	if err != nil {
		return fmt.Errorf("could not get cctp events: %w", err)
	}

	event, err := parser.ParseDepositForBurn(log)
	if err != nil {
		logger.Warnf("could not parse circle request sent: %w", err)
		return nil
	}

	// check that we sent the tx
	if event.Depositor != c.relayerAddress {
		span.AddEvent(fmt.Sprintf("depositor %s does not match relayer address %s", event.Depositor.String(), c.relayerAddress.String()))
		return nil
	}

	// use message hash as requestID
	sourceDomain, err := c.boundMessageTransmitters[chainID].LocalDomain(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get local domain: %w", err)
	}
	requestID := cctpRelay.GetCircleRequestID(sourceDomain, event.Nonce)
	span.SetAttributes(
		attribute.String("log_type", "DepositForBurn"),
		attribute.String("request_id", requestID),
	)
	origin := uint64(chainID)
	err = c.db.UpdateRebalanceStatus(ctx, requestID, &origin, reldb.RebalancePending)
	if err != nil {
		logger.Warnf("could not update rebalance status: %w", err)
		return nil
	}
	return nil
}

func (c *rebalanceManagerCircleCCTP) handleMessageReceived(ctx context.Context, log types.Log, chainID int, ethClient client.EVM) (err error) {
	ctx, span := c.handler.Tracer().Start(ctx, "rebalance.handleDepositForBurn")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	transmitterAddr, err := c.boundTokenMessengers[chainID].LocalMessageTransmitter(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get message transmitter: %w", err)
	}
	parser, err := messagetransmitter.NewMessageTransmitterFilterer(transmitterAddr, ethClient)
	if err != nil {
		return fmt.Errorf("could not get message transmitter events: %w", err)
	}

	parsedEvent, err := parser.ParseMessageReceived(log)
	if err != nil {
		logger.Warnf("could not parse circle request fulfilled: %w", err)
		return nil
	}

	// check that we sent the tx
	sender := cctpRelay.Bytes32ToAddress(parsedEvent.Sender)
	if sender != c.relayerAddress {
		span.AddEvent(fmt.Sprintf("sender %s does not match relayer address %s", sender.String(), c.relayerAddress.String()))
		return nil
	}

	requestID := cctpRelay.GetCircleRequestID(parsedEvent.SourceDomain, parsedEvent.Nonce)
	span.SetAttributes(
		attribute.String("log_type", "MessageReceived"),
		attribute.String("request_id", requestID),
	)
	err = c.db.UpdateRebalanceStatus(ctx, requestID, nil, reldb.RebalanceCompleted)
	if err != nil {
		logger.Warnf("could not update rebalance status: %w", err)
		return nil
	}
	return nil
}
