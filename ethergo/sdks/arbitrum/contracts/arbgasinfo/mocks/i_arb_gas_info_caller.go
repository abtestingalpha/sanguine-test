// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// IArbGasInfoCaller is an autogenerated mock type for the IArbGasInfoCaller type
type IArbGasInfoCaller struct {
	mock.Mock
}

// GetAmortizedCostCapBips provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetAmortizedCostCapBips(opts *bind.CallOpts) (uint64, error) {
	ret := _m.Called(opts)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentTxL1GasFees provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetCurrentTxL1GasFees(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasAccountingParams provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetGasAccountingParams(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) *big.Int); ok {
		r1 = rf(opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*bind.CallOpts) *big.Int); ok {
		r2 = rf(opts)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*bind.CallOpts) error); ok {
		r3 = rf(opts)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetGasBacklog provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetGasBacklog(opts *bind.CallOpts) (uint64, error) {
	ret := _m.Called(opts)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasBacklogTolerance provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetGasBacklogTolerance(opts *bind.CallOpts) (uint64, error) {
	ret := _m.Called(opts)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1BaseFeeEstimate provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetL1BaseFeeEstimate(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1BaseFeeEstimateInertia provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetL1BaseFeeEstimateInertia(opts *bind.CallOpts) (uint64, error) {
	ret := _m.Called(opts)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1FeesAvailable provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetL1FeesAvailable(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1GasPriceEstimate provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetL1GasPriceEstimate(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1PricingSurplus provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetL1PricingSurplus(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMinimumGasPrice provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetMinimumGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerBatchGasCharge provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetPerBatchGasCharge(opts *bind.CallOpts) (int64, error) {
	ret := _m.Called(opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) int64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPricesInArbGas provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetPricesInArbGas(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) *big.Int); ok {
		r1 = rf(opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*bind.CallOpts) *big.Int); ok {
		r2 = rf(opts)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*bind.CallOpts) error); ok {
		r3 = rf(opts)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetPricesInArbGasWithAggregator provides a mock function with given fields: opts, aggregator
func (_m *IArbGasInfoCaller) GetPricesInArbGasWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(opts, aggregator)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r0 = rf(opts, aggregator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r1 = rf(opts, aggregator)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r2 = rf(opts, aggregator)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*bind.CallOpts, common.Address) error); ok {
		r3 = rf(opts, aggregator)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetPricesInWei provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetPricesInWei(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) *big.Int); ok {
		r1 = rf(opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*bind.CallOpts) *big.Int); ok {
		r2 = rf(opts)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 *big.Int
	if rf, ok := ret.Get(3).(func(*bind.CallOpts) *big.Int); ok {
		r3 = rf(opts)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*big.Int)
		}
	}

	var r4 *big.Int
	if rf, ok := ret.Get(4).(func(*bind.CallOpts) *big.Int); ok {
		r4 = rf(opts)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).(*big.Int)
		}
	}

	var r5 *big.Int
	if rf, ok := ret.Get(5).(func(*bind.CallOpts) *big.Int); ok {
		r5 = rf(opts)
	} else {
		if ret.Get(5) != nil {
			r5 = ret.Get(5).(*big.Int)
		}
	}

	var r6 error
	if rf, ok := ret.Get(6).(func(*bind.CallOpts) error); ok {
		r6 = rf(opts)
	} else {
		r6 = ret.Error(6)
	}

	return r0, r1, r2, r3, r4, r5, r6
}

// GetPricesInWeiWithAggregator provides a mock function with given fields: opts, aggregator
func (_m *IArbGasInfoCaller) GetPricesInWeiWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(opts, aggregator)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r0 = rf(opts, aggregator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r1 = rf(opts, aggregator)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r2 = rf(opts, aggregator)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 *big.Int
	if rf, ok := ret.Get(3).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r3 = rf(opts, aggregator)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*big.Int)
		}
	}

	var r4 *big.Int
	if rf, ok := ret.Get(4).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r4 = rf(opts, aggregator)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).(*big.Int)
		}
	}

	var r5 *big.Int
	if rf, ok := ret.Get(5).(func(*bind.CallOpts, common.Address) *big.Int); ok {
		r5 = rf(opts, aggregator)
	} else {
		if ret.Get(5) != nil {
			r5 = ret.Get(5).(*big.Int)
		}
	}

	var r6 error
	if rf, ok := ret.Get(6).(func(*bind.CallOpts, common.Address) error); ok {
		r6 = rf(opts, aggregator)
	} else {
		r6 = ret.Error(6)
	}

	return r0, r1, r2, r3, r4, r5, r6
}

// GetPricingInertia provides a mock function with given fields: opts
func (_m *IArbGasInfoCaller) GetPricingInertia(opts *bind.CallOpts) (uint64, error) {
	ret := _m.Called(opts)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint64); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIArbGasInfoCaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewIArbGasInfoCaller creates a new instance of IArbGasInfoCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIArbGasInfoCaller(t mockConstructorTestingTNewIArbGasInfoCaller) *IArbGasInfoCaller {
	mock := &IArbGasInfoCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
