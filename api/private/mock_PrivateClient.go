// Code generated by mockery v1.0.0
package private

import mock "github.com/stretchr/testify/mock"
import models "github.com/fxpgr/go-exchange-client/models"

// MockPrivateClient is an autogenerated mock type for the PrivateClient type
type MockPrivateClient struct {
	mock.Mock
}

// ActiveOrders provides a mock function with given fields:
func (_m *MockPrivateClient) ActiveOrders() ([]*models.Order, error) {
	ret := _m.Called()

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func() []*models.Order); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields: c
func (_m *MockPrivateClient) Address(c string) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balances provides a mock function with given fields:
func (_m *MockPrivateClient) Balances() (map[string]float64, error) {
	ret := _m.Called()

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func() map[string]float64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelOrder provides a mock function with given fields: trading, settlement, ordertype, orderNumber
func (_m *MockPrivateClient) CancelOrder(trading string, settlement string, ordertype models.OrderType, orderNumber string) error {
	ret := _m.Called(trading, settlement, ordertype, orderNumber)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, models.OrderType, string) error); ok {
		r0 = rf(trading, settlement, ordertype, orderNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteBalance provides a mock function with given fields: coin
func (_m *MockPrivateClient) CompleteBalance(coin string) (*models.Balance, error) {
	ret := _m.Called(coin)

	var r0 *models.Balance
	if rf, ok := ret.Get(0).(func(string) *models.Balance); ok {
		r0 = rf(coin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Balance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(coin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteBalances provides a mock function with given fields:
func (_m *MockPrivateClient) CompleteBalances() (map[string]*models.Balance, error) {
	ret := _m.Called()

	var r0 map[string]*models.Balance
	if rf, ok := ret.Get(0).(func() map[string]*models.Balance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*models.Balance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsOrderFilled provides a mock function with given fields: trading, settlement, orderNumber
func (_m *MockPrivateClient) IsOrderFilled(trading string, settlement string, orderNumber string) (bool, error) {
	ret := _m.Called(trading, settlement, orderNumber)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(trading, settlement, orderNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(trading, settlement, orderNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Order provides a mock function with given fields: trading, settlement, ordertype, price, amount
func (_m *MockPrivateClient) Order(trading string, settlement string, ordertype models.OrderType, price float64, amount float64) (string, error) {
	ret := _m.Called(trading, settlement, ordertype, price, amount)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, models.OrderType, float64, float64) string); ok {
		r0 = rf(trading, settlement, ordertype, price, amount)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, models.OrderType, float64, float64) error); ok {
		r1 = rf(trading, settlement, ordertype, price, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TradeFeeRate provides a mock function with given fields: _a0, _a1
func (_m *MockPrivateClient) TradeFeeRate(_a0 string, _a1 string) (TradeFee, error) {
	ret := _m.Called(_a0, _a1)

	var r0 TradeFee
	if rf, ok := ret.Get(0).(func(string, string) TradeFee); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(TradeFee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TradeFeeRates provides a mock function with given fields:
func (_m *MockPrivateClient) TradeFeeRates() (map[string]map[string]TradeFee, error) {
	ret := _m.Called()

	var r0 map[string]map[string]TradeFee
	if rf, ok := ret.Get(0).(func() map[string]map[string]TradeFee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]TradeFee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: typ, addr, amount, additionalFee
func (_m *MockPrivateClient) Transfer(typ string, addr string, amount float64, additionalFee float64) error {
	ret := _m.Called(typ, addr, amount, additionalFee)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, float64, float64) error); ok {
		r0 = rf(typ, addr, amount, additionalFee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferFee provides a mock function with given fields:
func (_m *MockPrivateClient) TransferFee() (map[string]float64, error) {
	ret := _m.Called()

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func() map[string]float64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
