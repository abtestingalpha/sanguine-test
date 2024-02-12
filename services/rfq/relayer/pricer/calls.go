package pricer

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
)

const mockRequestData = "0x000000000000000000000000000000000000000000000000000000000000a4b1000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000b73acb429ba868984c0236bdf940d4fe1e643f27000000000000000000000000b73acb429ba868984c0236bdf940d4fe1e643f27000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000002dc6c00000000000000000000000000000000000000000000000000000000000104d35000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000065a17be3000000000000000000000000000000000000000000000000000000000000001c"

func getCallRequestData() []byte {
	data, err := hexutil.Decode(mockRequestData)
	if err != nil {
		panic(err)
	}
	return data
}

func getCallAddress() common.Address {
	return chain.EthAddress
}

func getCallHash() [32]byte {
	return [32]byte{}
}

type callType int

const (
	// claimCallType is the call type for claim.
	claimCallType callType = iota + 1
	// proveCallType is the call type for prove.
	proveCallType
	// relayCallType is the call type for relay.
	relayCallType
)

func (c callType) String() string {
	switch c {
	case claimCallType:
		return "claim"
	case proveCallType:
		return "prove"
	case relayCallType:
		return "relay"
	}
	return ""
}

func getCall(transactor *bind.TransactOpts, bridge *fastbridge.FastBridgeRef, cType callType) (call *ethereum.CallMsg, err error) {
	var tx *types.Transaction
	var callFunc func(context.Context) error
	switch cType {
	case claimCallType:
		callFunc = func(context.Context) error {
			tx, err = bridge.Claim(transactor, getCallRequestData(), getCallAddress())
			if err != nil {
				fmt.Printf("inner call err: %v\n", err)
			}
			return err
		}
	case proveCallType:
		callFunc = func(context.Context) error {
			tx, err = bridge.Prove(transactor, getCallRequestData(), getCallHash())
			if err != nil {
				fmt.Printf("inner call err: %v\n", err)
			}
			return err
		}
	case relayCallType:
		callFunc = func(context.Context) error {
			tx, err = bridge.Relay(transactor, getCallRequestData())
			if err != nil {
				fmt.Printf("inner call err: %v\n", err)
			}
			return err
		}
	default:
		return nil, fmt.Errorf("unknown call type: %d", cType)
	}
	err = retry.WithBackoff(transactor.Context, func(ctx context.Context) error {
		return callFunc(transactor.Context)
	}, retry.WithMaxTotalTime(5*time.Second))
	if err != nil {
		fmt.Printf("call err: %v\n", err)
		return nil, fmt.Errorf("could not get tx with type %s: %w", cType.String(), err)
	}
	call, err = util.TxToCall(tx)
	if err != nil {
		return nil, fmt.Errorf("could not get call: %w", err)
	}
	if call == nil {
		return nil, fmt.Errorf("call is nil")
	}
	return call, nil
}
