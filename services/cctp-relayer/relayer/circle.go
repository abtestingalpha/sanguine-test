package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/circlecctp"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"gorm.io/gorm"
)

type circleCCTPHandler struct {
	cfg              config.Config
	db               db2.CCTPRelayerDB
	omniRPCClient    omniClient.RPCClient
	boundCircleCCTPs map[uint32]*circlecctp.MessageTransmitter
	txSubmitter      submitter.TransactionSubmitter
	handler          metrics.Handler
}

// NewCircleCCTPHandler creates a new CircleCCTPHandler.
func NewCircleCCTPHandler(ctx context.Context, cfg config.Config, db db2.CCTPRelayerDB, omniRPCClient omniClient.RPCClient, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) (CCTPHandler, error) {
	boundCircleCCTPs := make(map[uint32]*circlecctp.MessageTransmitter)
	for _, chain := range cfg.Chains {
		cl, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundCircleCCTPs[chain.ChainID], err = circlecctp.NewMessageTransmitter(chain.GetCircleCCTPAddress(), cl)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
	}
	return &circleCCTPHandler{
		cfg:              cfg,
		db:               db,
		omniRPCClient:    omniRPCClient,
		boundCircleCCTPs: boundCircleCCTPs,
		txSubmitter:      txSubmitter,
		handler:          handler,
	}, nil
}

func (c *circleCCTPHandler) HandleLog(ctx context.Context, log *types.Log, chainID uint32) (processQueue bool, err error) {
	if log == nil {
		return false, fmt.Errorf("log is nil")
	}

	// shouldn't be possible: maybe remove?
	if len(log.Topics) == 0 {
		return false, fmt.Errorf("not enough topics")
	}

	switch log.Topics[0] {
	case circlecctp.MessageSentTopic:
		msg, err := c.handleMessageSent(ctx, log, chainID)
		if err != nil {
			return false, fmt.Errorf("could not store message sent: %w", err)
		}

		if msg != nil {
			processQueue = true
		}

		return processQueue, nil

	case circlecctp.MessageReceivedTopic:
		err = c.handleMessageReceived(ctx, log, chainID)
		return false, nil
	default:
		logger.Warnf("unknown topic %s", log.Topics[0])
		return false, nil
	}
}

func (c *circleCCTPHandler) FetchAndProcessSentEvent(ctx context.Context, txHash common.Hash, chainID uint32) (msg *relayTypes.Message, err error) {
	// check if message already exist before we do anything
	msg, err = c.db.GetMessageByOriginHash(ctx, txHash)
	// if we already have the message, we can just return it
	if err == nil {
		return msg, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("could not get message by origin hash: %w", err)
	}

	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(chainID), 1)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	// TODO: consider pulling from scribe?
	// TODO: this function is particularly prone to silent errors
	receipt, err := ethClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	for _, log := range receipt.Logs {
		// this should never happen
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] == circlecctp.MessageSentTopic {
			msg, err = c.handleMessageSent(ctx, log, chainID)
			if err != nil {
				return nil, fmt.Errorf("could not handle message sent: %w", err)
			}
			return msg, nil
		}
	}
	return nil, fmt.Errorf("no message sent event found")
}

func (c *circleCCTPHandler) SubmitReceiveMessage(ctx context.Context, msg *relayTypes.Message) (err error) {
	contract, ok := c.boundCircleCCTPs[msg.DestChainID]
	if !ok {
		return fmt.Errorf("no contract found for chain %d", msg.DestChainID)
	}

	//TODO: check if already executed on destination?

	var nonce uint64
	var destTxHash common.Hash
	nonce, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(msg.DestChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.ReceiveMessage(transactor, msg.Message, msg.Attestation)
		if err != nil {
			return nil, fmt.Errorf("could not submit transaction: %w", err)
		}

		destTxHash = tx.Hash()
		return tx, nil
	})
	if err != nil {
		err = fmt.Errorf("could not submit transaction: %w", err)
		return err
	}

	// Store the completed message.
	// Note: this can cause double submission sometimes
	msg.State = relayTypes.Submitted
	msg.DestNonce = int(nonce)
	msg.DestTxHash = destTxHash.String()
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store completed message: %w", err)
	}

	return nil
}

func (c *circleCCTPHandler) handleMessageSent(ctx context.Context, log *types.Log, chainID uint32) (*relayTypes.Message, error) {
	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(chainID), 1)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	eventParser, err := circlecctp.NewMessageTransmitterFilterer(log.Address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("could not create event parser: %w", err)
	}

	messageSentEvent, err := eventParser.ParseMessageSent(*log)
	if err != nil {
		return nil, fmt.Errorf("could not parse message sent: %w", err)
	}

	destChainID, err := parseDestChainID(messageSentEvent.Message)
	if err != nil {
		return nil, fmt.Errorf("could not parse destination chain from raw message")
	}

	rawMsg := relayTypes.Message{
		OriginTxHash:  log.TxHash.Hex(),
		OriginChainID: chainID,
		DestChainID:   destChainID,
		Message:       messageSentEvent.Message,
		MessageHash:   crypto.Keccak256Hash(messageSentEvent.Message).String(),

		//Attestation: //comes from the api
		BlockNumber: log.BlockNumber,
	}

	// Store the requested message.
	rawMsg.State = relayTypes.Pending
	err = c.db.StoreMessage(ctx, rawMsg)
	if err != nil {
		return nil, fmt.Errorf("could not store pending message: %w", err)
	}

	return &rawMsg, nil
}

func (c *circleCCTPHandler) handleMessageReceived(ctx context.Context, log *types.Log, destChain uint32) (err error) {
	if len(log.Topics) == 0 {
		return fmt.Errorf("no topics found")
	}

	// Parse the request id from the log.
	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(destChain), 1)
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}

	eventParser, err := circlecctp.NewMessageTransmitterFilterer(log.Address, ethClient)
	if err != nil {
		return fmt.Errorf("could not create event parser: %w", err)
	}

	event, err := eventParser.ParseMessageReceived(*log)
	if err != nil {
		return fmt.Errorf("could not parse circle request fulfilled: %w", err)
	}

	messageHash := crypto.Keccak256Hash(event.MessageBody)
	msg, err := c.db.GetMessageByHash(ctx, messageHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Reconstruct what we can from the given log.
			msg = &relayTypes.Message{
				OriginChainID: event.SourceDomain,
				DestChainID:   destChain,
				BlockNumber:   log.BlockNumber,
				MessageHash:   messageHash.String(),
			}
		} else {
			return fmt.Errorf("could not get message by hash: %w", err)
		}
	}

	if msg == nil {
		return fmt.Errorf("no message found")
	}

	// Mark as Complete and store the message.
	msg.State = relayTypes.Complete
	msg.DestTxHash = log.TxHash.String()
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store complete message: %w", err)
	}
	return nil
}
