package arbitrum

import "github.com/ethereum/go-ethereum/common"

type arbitrumOptions struct {
	// gasInfoAddress is the address of the ArbGasInfo contract.
	// it is said to be the same on all arbitrum-like l2 chains.
	gasInfoAddress common.Address
	// nodeInterfaceAddress is the address of the NodeInterface contract.
	// it is said to be the same on all arbitrum-like l2 chains.
	nodeInterfaceAddress common.Address
}

func defaultOptions() *arbitrumOptions {
	return &arbitrumOptions{
		// see: https://github.com/OffchainLabs/arbitrum-token-bridge/blob/75915c94e58aaf7bf59fb833a0a1b3be1ae461ec/packages/arb-token-bridge-ui/scripts/generateDenylist.ts#L63 and https://github.com/Tenderly/nitro/blob/master/go-ethereum/core/types/arbitrum_signer.go#L12
		gasInfoAddress:       common.HexToAddress("0x6e"),
		nodeInterfaceAddress: common.HexToAddress("0xc8"),
	}
}

// Option is an option for the Arbitrum SDK.
type Option func(*arbitrumOptions)

// WithGasInfoAddress sets the gas info address.
func WithGasInfoAddress(address common.Address) Option {
	return func(o *arbitrumOptions) {
		o.gasInfoAddress = address
	}
}

// WithNodeInterfaceAddress sets the node interface address.
func WithNodeInterfaceAddress(address common.Address) Option {
	return func(o *arbitrumOptions) {
		o.nodeInterfaceAddress = address
	}
}
