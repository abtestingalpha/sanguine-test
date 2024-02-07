package arbitrum

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/contracts/arbgasinfo"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/contracts/nodeinterface"
)

// SDK is an interface for interacting with the Arbitrum SDK.
type SDK interface {
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
}

type arbitrumSDKImpl struct {
	client        bind.ContractBackend
	nodeInterface nodeinterface.INodeInterface
	gasInfo       arbgasinfo.IArbGasInfo
}

// NewArbitrumSDK creates a new SDK.
func NewArbitrumSDK(client bind.ContractBackend, options ...Option) (SDK, error) {
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	nodeInterface, err := nodeinterface.NewNodeInterfaceRef(opts.nodeInterfaceAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create node interface: %w", err)
	}

	gasInfo, err := arbgasinfo.NewArbGasInfo(opts.gasInfoAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gas info: %w", err)
	}

	return &arbitrumSDKImpl{
		client:        client,
		nodeInterface: nodeInterface,
		gasInfo:       gasInfo,
	}, nil
}

func (a *arbitrumSDKImpl) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	if call.To == nil {
		return 0, errors.New("call.To cannot be nil")
	}
	// TODO: maybe need to copy the logic that sets the gasprice if it's empty?
	gasEstimate, gasEstimateForL1, _, _, err := a.nodeInterface.GetGasEstimateComponents(&bind.TransactOpts{
		Context: ctx,
		From:    call.From,
		// note: this is ignored
		GasLimit:  call.Gas,
		GasPrice:  core.CopyBigInt(call.GasPrice),
		GasFeeCap: core.CopyBigInt(call.GasFeeCap),
		GasTipCap: core.CopyBigInt(call.GasTipCap),
		Value:     core.CopyBigInt(call.Value),
	}, *call.To, false, call.Data)
	if err != nil {
		return 0, fmt.Errorf("failed to get gas estimate components: %w", err)
	}
	return gasEstimate + gasEstimateForL1, nil
}

// This is a type assertion used to make sure the arbitrum sdk matches the standard contracttransactor interface
// methods for gas without doing the whole things.
// It will fail to compile if methods are different.
//
//nolint:deadcode
type unfiormMethodAssertion interface {
	bind.ContractTransactor
	SDK
}
