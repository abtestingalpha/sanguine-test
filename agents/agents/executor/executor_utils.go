package executor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (types.Message, error) {
	message, ok := e.chainExecutors[chainID].originParser.ParseSent(log)
	if !ok {
		return nil, fmt.Errorf("could not parse committed message")
	}

	if message == nil {
		//nolint:nilnil
		return nil, nil
	}

	return message, nil
}

// logToAttestation converts the log to an attestation.
func (e Executor) logToAttestation(log ethTypes.Log, chainID uint32, summitAttestation bool) (types.Attestation, error) {
	fmt.Printf("logToAttestation on chain %d and log tx hash %s, summitAttestation %v\n", chainID, log.TxHash.Hex(), summitAttestation)
	var attestation types.Attestation
	var ok bool

	if summitAttestation {
		attestation, ok = e.chainExecutors[chainID].summitParser.ParseAttestationSaved(log)
		if !ok {
			return nil, fmt.Errorf("could not parse attestation")
		}
	} else {
		attestationMetadata, err := e.chainExecutors[chainID].lightInboxParser.ParseAttestationAccepted(log)
		if err != nil {
			return nil, fmt.Errorf("could not parse attestation: %w", err)
		}

		attestation = attestationMetadata.Attestation
	}

	if attestation == nil {
		//nolint:nilnil
		return nil, nil
	}

	return attestation, nil
}

// logToSnapshot converts the log to a snapshot.
func (e Executor) logToSnapshot(log ethTypes.Log, chainID uint32) (types.Snapshot, error) {
	snapshotMetadata, err := e.chainExecutors[chainID].inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse snapshot: %w", err)
	}

	if snapshotMetadata.Snapshot == nil || snapshotMetadata.AgentDomain() == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return snapshotMetadata.Snapshot, nil
}

func (e Executor) logToInterface(log ethTypes.Log, chainID uint32) (any, error) {
	switch {
	case e.isSnapshotAcceptedEvent(log, chainID):
		return e.logToSnapshot(log, chainID)
	case e.isSentEvent(log, chainID):
		return e.logToMessage(log, chainID)
	case e.isAttestationAcceptedEvent(log, chainID):
		return e.logToAttestation(log, chainID, false)
	case e.isAttestationSavedEvent(log, chainID):
		return e.logToAttestation(log, chainID, true)
	default:
		fmt.Printf("logToInterface: unknown event type on chain %d with log tx hash %s\n", chainID, log.TxHash.Hex())
		//nolint:nilnil
		return nil, nil
	}
}

func (e Executor) isSnapshotAcceptedEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].inboxParser == nil {
		return false
	}

	inboxEvent, ok := e.chainExecutors[chainID].inboxParser.EventType(log)
	return ok && inboxEvent == inbox.SnapshotAcceptedEvent
}

func (e Executor) isSentEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].originParser == nil {
		return false
	}

	originEvent, ok := e.chainExecutors[chainID].originParser.EventType(log)
	return ok && originEvent == origin.SentEvent
}

func (e Executor) isAttestationAcceptedEvent(log ethTypes.Log, chainID uint32) bool {
	fmt.Printf("isAttestationAcceptedEvent on chain %d with log tx hash %s\n", chainID, log.TxHash.Hex())
	fmt.Printf("lightinboxparser: %v\n", e.chainExecutors[chainID].lightInboxParser)
	if e.chainExecutors[chainID].lightInboxParser == nil {
		return false
	}

	lightManagerEvent, ok := e.chainExecutors[chainID].lightInboxParser.EventType(log)
	fmt.Printf("ok: %v, lightManagerEvent: %v\n", ok, lightManagerEvent)
	return ok && lightManagerEvent == lightinbox.AttestationAcceptedEvent
}

func (e Executor) isAttestationSavedEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].summitParser == nil {
		return false
	}

	summitEvent, ok := e.chainExecutors[chainID].summitParser.EventType(log)
	return ok && summitEvent == summit.AttestationSavedEvent
}

// processMessage processes and stores a message.
func (e Executor) processMessage(ctx context.Context, message types.Message, log ethTypes.Log) (err error) {
	types.LogTx("EXECUTOR", fmt.Sprintf("Processing message: %s", types.MessageToString(message)), message.OriginDomain(), nil)
	ctx, span := e.handler.Tracer().Start(ctx, "processMessage", trace.WithAttributes(
		attribute.String("txHash", log.TxHash.String()),
		attribute.Int("logBlockNumber", int(log.BlockNumber)),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	merkleIndex := e.chainExecutors[message.OriginDomain()].merkleTree.NumOfItems()
	leaf, err := message.ToLeaf()
	if err != nil {
		return fmt.Errorf("could not convert message to leaf: %w", err)
	}

	// Make sure the nonce of the message is being inserted at the right index.
	switch {
	case merkleIndex+1 > message.Nonce():
		return nil
	case merkleIndex+1 < message.Nonce():
		span.AddEvent("nonce is not correct", trace.WithAttributes(
			attribute.Int("nonce", int(message.Nonce())),
			attribute.Int("merkle_index", int(merkleIndex)),
		))
		logger.Warnf("nonce is not correct. expected: %d, got: %d", merkleIndex+1, message.Nonce())
		return fmt.Errorf("nonce is not correct. expected: %d, got: %d", merkleIndex+1, message.Nonce())
	default:
	}

	e.chainExecutors[message.OriginDomain()].merkleTree.Insert(leaf[:])

	span.AddEvent("storing message", trace.WithAttributes(attribute.Int("nonce", int(message.Nonce()))))
	err = e.executorDB.StoreMessage(ctx, message, log.BlockNumber, false, 0, log.TxHash)
	if err != nil {
		return fmt.Errorf("could not store message: %w", err)
	}

	return nil
}

// processAttestation processes and stores an attestation.
func (e Executor) processSnapshot(ctx context.Context, snapshot types.Snapshot, log ethTypes.Log) (err error) {
	snapshotRoot, proofs, err := snapshot.SnapshotRootAndProofs()
	if err != nil {
		return fmt.Errorf("could not get snapshot root and proofs: %w", err)
	}

	ctx, span := e.handler.Tracer().Start(ctx, "processSnapshot", trace.WithAttributes(
		attribute.Int("logBlockNumber", int(log.BlockNumber)),
		attribute.String("snapRoot", common.BytesToHash(snapshotRoot[:]).String()),
		attribute.String("txHash", log.TxHash.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	for _, s := range snapshot.States() {
		state := s
		statePayload, err := state.Encode()
		if err != nil {
			return fmt.Errorf("could not encode state: %w", err)
		}
		// Verify that the state is valid w.r.t. Origin.
		var valid bool
		contractCall := func(ctx context.Context) error {
			valid, err = e.chainExecutors[state.Origin()].boundOrigin.IsValidState(
				ctx,
				statePayload,
			)
			if err != nil {
				return fmt.Errorf("could not check validity of state: %w", err)
			}

			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
		if err != nil {
			return fmt.Errorf("could not check validity of state: %w", err)
		}

		if !valid {
			stateRoot := state.Root()
			logger.Infof("snapshot has invalid state. Origin: %d. SnapshotRoot: %s", state.Origin(), common.BytesToHash(stateRoot[:]).String())
			return nil
		}
	}

	span.AddEvent("storing states", trace.WithAttributes(attribute.Int("num_states", len(snapshot.States()))))
	err = e.executorDB.StoreStates(ctx, snapshot.States(), snapshotRoot, proofs, log.BlockNumber)
	if err != nil {
		return fmt.Errorf("could not store states: %w", err)
	}

	return nil
}

// processAttestation processes and stores an attestation.
func (e Executor) processAttestation(ctx context.Context, attestation types.Attestation, chainID uint32, log ethTypes.Log) (err error) {
	fmt.Printf("processAttestation on chain %d: %v\n", chainID, attestation)
	snapshotRoot := attestation.SnapshotRoot()
	ctx, span := e.handler.Tracer().Start(ctx, "processAttestation", trace.WithAttributes(
		attribute.Int("chainID", int(chainID)),
		attribute.Int("logBlockNumber", int(log.BlockNumber)),
		attribute.String("snapRoot", common.BytesToHash(snapshotRoot[:]).String()),
		attribute.String("txHash", log.TxHash.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// If the attestation is on the SynChain, we can directly use its block number and timestamp.
	if chainID == e.config.SummitChainID {
		span.AddEvent("storing summit attestation", trace.WithAttributes(attribute.Int("time", int(attestation.Timestamp().Uint64()))))
		err := e.executorDB.StoreAttestation(ctx, attestation, chainID, attestation.BlockNumber().Uint64(), attestation.Timestamp().Uint64())
		if err != nil {
			return fmt.Errorf("could not store attestation: %w", err)
		}

		return nil
	}

	// If the attestation is on a remote chain, we need to fetch the timestamp via an RPC call.
	var logHeader *ethTypes.Header
	contractCall := func(ctx context.Context) error {
		logHeader, err = e.chainExecutors[chainID].rpcClient.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
		if err != nil {
			return fmt.Errorf("could not get log header: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get log header: %w", err)
	}

	if logHeader == nil {
		return fmt.Errorf("could not get log header")
	}

	fmt.Printf("storing attestation: %v\n", attestation)
	span.AddEvent("storing remote attestation", trace.WithAttributes(attribute.Int("time", int(logHeader.Time))))
	err = e.executorDB.StoreAttestation(ctx, attestation, chainID, log.BlockNumber, logHeader.Time)
	if err != nil {
		return fmt.Errorf("could not store attestation: %w", err)
	}

	return nil
}
