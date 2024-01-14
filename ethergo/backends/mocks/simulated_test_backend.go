// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	backends "github.com/synapsecns/sanguine/ethergo/backends"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	chainwatcher "github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	contracts "github.com/synapsecns/sanguine/ethergo/contracts"

	ecdsa "crypto/ecdsa"

	ethereum "github.com/ethereum/go-ethereum"

	gas "github.com/synapsecns/sanguine/ethergo/chain/gas"

	keystore "github.com/ethereum/go-ethereum/accounts/keystore"

	mock "github.com/stretchr/testify/mock"

	nonce "github.com/synapsecns/sanguine/ethergo/signer/nonce"

	params "github.com/ethereum/go-ethereum/params"

	rpc "github.com/ethereum/go-ethereum/rpc"

	testing "testing"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"

	w3types "github.com/lmittmann/w3/w3types"
)

// SimulatedTestBackend is an autogenerated mock type for the SimulatedTestBackend type
type SimulatedTestBackend struct {
	mock.Mock
}

// AttemptReconnect provides a mock function with given fields:
func (_m *SimulatedTestBackend) AttemptReconnect() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BackendName provides a mock function with given fields:
func (_m *SimulatedTestBackend) BackendName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// BalanceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *SimulatedTestBackend) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *SimulatedTestBackend) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchContext provides a mock function with given fields: ctx, calls
func (_m *SimulatedTestBackend) BatchContext(ctx context.Context, calls ...w3types.Caller) error {
	_va := make([]interface{}, len(calls))
	for _i := range calls {
		_va[_i] = calls[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...w3types.Caller) error); ok {
		r0 = rf(ctx, calls...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchWithContext provides a mock function with given fields: ctx, calls
func (_m *SimulatedTestBackend) BatchWithContext(ctx context.Context, calls ...w3types.Caller) error {
	_va := make([]interface{}, len(calls))
	for _i := range calls {
		_va[_i] = calls[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...w3types.Caller) error); ok {
		r0 = rf(ctx, calls...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *SimulatedTestBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *SimulatedTestBackend) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockNumber provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) BlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallContext provides a mock function with given fields: ctx, result, method, args
func (_m *SimulatedTestBackend) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallContract provides a mock function with given fields: ctx, call, blockNumber
func (_m *SimulatedTestBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, call, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, call, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, call, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainConfig provides a mock function with given fields:
func (_m *SimulatedTestBackend) ChainConfig() *params.ChainConfig {
	ret := _m.Called()

	var r0 *params.ChainConfig
	if rf, ok := ret.Get(0).(func() *params.ChainConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*params.ChainConfig)
		}
	}

	return r0
}

// ChainID provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) ChainID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainName provides a mock function with given fields:
func (_m *SimulatedTestBackend) ChainName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ClearNonce provides a mock function with given fields: address
func (_m *SimulatedTestBackend) ClearNonce(address common.Address) {
	_m.Called(address)
}

// ClientID provides a mock function with given fields:
func (_m *SimulatedTestBackend) ClientID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CodeAt provides a mock function with given fields: ctx, contract, blockNumber
func (_m *SimulatedTestBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, contract, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, contract, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, contract, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConcurrencyCount provides a mock function with given fields:
func (_m *SimulatedTestBackend) ConcurrencyCount() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// EstimateGas provides a mock function with given fields: ctx, call
func (_m *SimulatedTestBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(ctx, call)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(ctx, call)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, call)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Estimator provides a mock function with given fields:
func (_m *SimulatedTestBackend) Estimator() gas.PriceEstimator {
	ret := _m.Called()

	var r0 gas.PriceEstimator
	if rf, ok := ret.Get(0).(func() gas.PriceEstimator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.PriceEstimator)
		}
	}

	return r0
}

// FeeHistory provides a mock function with given fields: ctx, blockCount, lastBlock, rewardPercentiles
func (_m *SimulatedTestBackend) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	ret := _m.Called(ctx, blockCount, lastBlock, rewardPercentiles)

	var r0 *ethereum.FeeHistory
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *big.Int, []float64) *ethereum.FeeHistory); ok {
		r0 = rf(ctx, blockCount, lastBlock, rewardPercentiles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethereum.FeeHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, *big.Int, []float64) error); ok {
		r1 = rf(ctx, blockCount, lastBlock, rewardPercentiles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *SimulatedTestBackend) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FundAccount provides a mock function with given fields: ctx, address, amount
func (_m *SimulatedTestBackend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	_m.Called(ctx, address, amount)
}

// GasSetter provides a mock function with given fields:
func (_m *SimulatedTestBackend) GasSetter() gas.Setter {
	ret := _m.Called()

	var r0 gas.Setter
	if rf, ok := ret.Get(0).(func() gas.Setter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.Setter)
		}
	}

	return r0
}

// GetBigChainID provides a mock function with given fields:
func (_m *SimulatedTestBackend) GetBigChainID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetChainID provides a mock function with given fields:
func (_m *SimulatedTestBackend) GetChainID() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GetFundedAccount provides a mock function with given fields: ctx, requestBalance
func (_m *SimulatedTestBackend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	ret := _m.Called(ctx, requestBalance)

	var r0 *keystore.Key
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *keystore.Key); ok {
		r0 = rf(ctx, requestBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keystore.Key)
		}
	}

	return r0
}

// GetHeightWatcher provides a mock function with given fields:
func (_m *SimulatedTestBackend) GetHeightWatcher() chainwatcher.BlockHeightWatcher {
	ret := _m.Called()

	var r0 chainwatcher.BlockHeightWatcher
	if rf, ok := ret.Get(0).(func() chainwatcher.BlockHeightWatcher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chainwatcher.BlockHeightWatcher)
		}
	}

	return r0
}

// GetNextNonce provides a mock function with given fields: address
func (_m *SimulatedTestBackend) GetNextNonce(address common.Address) (*big.Int, error) {
	ret := _m.Called(address)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxContext provides a mock function with given fields: ctx, address
func (_m *SimulatedTestBackend) GetTxContext(ctx context.Context, address *common.Address) backends.AuthType {
	ret := _m.Called(ctx, address)

	var r0 backends.AuthType
	if rf, ok := ret.Get(0).(func(context.Context, *common.Address) backends.AuthType); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(backends.AuthType)
	}

	return r0
}

// HeaderByHash provides a mock function with given fields: ctx, hash
func (_m *SimulatedTestBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Header); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *SimulatedTestBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByTime provides a mock function with given fields: ctx, startBlock, searchTime
func (_m *SimulatedTestBackend) HeaderByTime(ctx context.Context, startBlock *big.Int, searchTime time.Time) (*types.Header, error) {
	ret := _m.Called(ctx, startBlock, searchTime)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, time.Time) *types.Header); ok {
		r0 = rf(ctx, startBlock, searchTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, time.Time) error); ok {
		r1 = rf(ctx, startBlock, searchTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImpersonateAccount provides a mock function with given fields: ctx, address, transact
func (_m *SimulatedTestBackend) ImpersonateAccount(ctx context.Context, address common.Address, transact func(*bind.TransactOpts) *types.Transaction) error {
	ret := _m.Called(ctx, address, transact)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, func(*bind.TransactOpts) *types.Transaction) error); ok {
		r0 = rf(ctx, address, transact)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NetworkID provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) NetworkID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyedTransactor provides a mock function with given fields: realSigner
func (_m *SimulatedTestBackend) NewKeyedTransactor(realSigner *bind.TransactOpts) (*bind.TransactOpts, error) {
	ret := _m.Called(realSigner)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *bind.TransactOpts); ok {
		r0 = rf(realSigner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(realSigner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyedTransactorFromKey provides a mock function with given fields: key
func (_m *SimulatedTestBackend) NewKeyedTransactorFromKey(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	ret := _m.Called(key)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey) *bind.TransactOpts); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecdsa.PrivateKey) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *SimulatedTestBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) uint64); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingBalanceAt provides a mock function with given fields: ctx, account
func (_m *SimulatedTestBackend) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	ret := _m.Called(ctx, account)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) *big.Int); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingCallContract provides a mock function with given fields: ctx, call
func (_m *SimulatedTestBackend) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	ret := _m.Called(ctx, call)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) []byte); ok {
		r0 = rf(ctx, call)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, call)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingCodeAt provides a mock function with given fields: ctx, account
func (_m *SimulatedTestBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ret := _m.Called(ctx, account)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) []byte); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingNonceAt provides a mock function with given fields: ctx, account
func (_m *SimulatedTestBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingStorageAt provides a mock function with given fields: ctx, account, key
func (_m *SimulatedTestBackend) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	ret := _m.Called(ctx, account, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) []byte); ok {
		r0 = rf(ctx, account, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, account, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingTransactionCount provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) PendingTransactionCount(ctx context.Context) (uint, error) {
	ret := _m.Called(ctx)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context) uint); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPCAddress provides a mock function with given fields:
func (_m *SimulatedTestBackend) RPCAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestCount provides a mock function with given fields:
func (_m *SimulatedTestBackend) RequestCount() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *SimulatedTestBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChainConfig provides a mock function with given fields: config
func (_m *SimulatedTestBackend) SetChainConfig(config *params.ChainConfig) {
	_m.Called(config)
}

// SetT provides a mock function with given fields: t
func (_m *SimulatedTestBackend) SetT(t *testing.T) {
	_m.Called(t)
}

// SignTx provides a mock function with given fields: ogTx, signer, prv, options
func (_m *SimulatedTestBackend) SignTx(ogTx *types.Transaction, signer types.Signer, prv *ecdsa.PrivateKey, options ...nonce.Option) (*types.Transaction, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ogTx, signer, prv)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*types.Transaction, types.Signer, *ecdsa.PrivateKey, ...nonce.Option) *types.Transaction); ok {
		r0 = rf(ogTx, signer, prv, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Transaction, types.Signer, *ecdsa.PrivateKey, ...nonce.Option) error); ok {
		r1 = rf(ogTx, signer, prv, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Signer provides a mock function with given fields:
func (_m *SimulatedTestBackend) Signer() types.Signer {
	ret := _m.Called()

	var r0 types.Signer
	if rf, ok := ret.Get(0).(func() types.Signer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Signer)
		}
	}

	return r0
}

// StorageAt provides a mock function with given fields: ctx, account, key, blockNumber
func (_m *SimulatedTestBackend) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, account, key, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash, *big.Int) []byte); ok {
		r0 = rf(ctx, account, key, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash, *big.Int) error); ok {
		r1 = rf(ctx, account, key, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: key
func (_m *SimulatedTestBackend) Store(key *keystore.Key) {
	_m.Called(key)
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, q, ch
func (_m *SimulatedTestBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, q, ch)

	var r0 ethereum.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, q, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, q, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *SimulatedTestBackend) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	var r0 ethereum.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasTipCap provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncProgress provides a mock function with given fields: ctx
func (_m *SimulatedTestBackend) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	ret := _m.Called(ctx)

	var r0 *ethereum.SyncProgress
	if rf, ok := ret.Get(0).(func(context.Context) *ethereum.SyncProgress); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethereum.SyncProgress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// T provides a mock function with given fields:
func (_m *SimulatedTestBackend) T() *testing.T {
	ret := _m.Called()

	var r0 *testing.T
	if rf, ok := ret.Get(0).(func() *testing.T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*testing.T)
		}
	}

	return r0
}

// TransactionByHash provides a mock function with given fields: ctx, txHash
func (_m *SimulatedTestBackend) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionCount provides a mock function with given fields: ctx, blockHash
func (_m *SimulatedTestBackend) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) uint); ok {
		r0 = rf(ctx, blockHash)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionInBlock provides a mock function with given fields: ctx, blockHash, index
func (_m *SimulatedTestBackend) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Transaction); ok {
		r0 = rf(ctx, blockHash, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, blockHash, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *SimulatedTestBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyContract provides a mock function with given fields: contractType, contract
func (_m *SimulatedTestBackend) VerifyContract(contractType contracts.ContractType, contract contracts.DeployedContract) error {
	ret := _m.Called(contractType, contract)

	var r0 error
	if rf, ok := ret.Get(0).(func(contracts.ContractType, contracts.DeployedContract) error); ok {
		r0 = rf(contractType, contract)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForConfirmation provides a mock function with given fields: ctx, transaction
func (_m *SimulatedTestBackend) WaitForConfirmation(ctx context.Context, transaction *types.Transaction) {
	_m.Called(ctx, transaction)
}

type mockConstructorTestingTNewSimulatedTestBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewSimulatedTestBackend creates a new instance of SimulatedTestBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSimulatedTestBackend(t mockConstructorTestingTNewSimulatedTestBackend) *SimulatedTestBackend {
	mock := &SimulatedTestBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
