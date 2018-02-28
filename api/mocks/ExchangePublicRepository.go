// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/fxpgr/go-ccex-api-client/models"

// ExchangePublicRepository is an autogenerated mock type for the ExchangePublicRepository type
type ExchangePublicRepository struct {
	mock.Mock
}

// CurrencyPairs provides a mock function with given fields:
func (_m *ExchangePublicRepository) CurrencyPairs() ([]*models.CurrencyPair, error) {
	ret := _m.Called()

	var r0 []*models.CurrencyPair
	if rf, ok := ret.Get(0).(func() []*models.CurrencyPair); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.CurrencyPair)
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

// Rate provides a mock function with given fields: trading, settlement
func (_m *ExchangePublicRepository) Rate(trading string, settlement string) (float64, error) {
	ret := _m.Called(trading, settlement)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string, string) float64); ok {
		r0 = rf(trading, settlement)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(trading, settlement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Volume provides a mock function with given fields: trading, settlement
func (_m *ExchangePublicRepository) Volume(trading string, settlement string) (float64, error) {
	ret := _m.Called(trading, settlement)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string, string) float64); ok {
		r0 = rf(trading, settlement)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(trading, settlement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}