// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// PayloadAttributer is an autogenerated mock type for the PayloadAttributer type
type PayloadAttributer struct {
	mock.Mock
}

type PayloadAttributer_Expecter struct {
	mock *mock.Mock
}

func (_m *PayloadAttributer) EXPECT() *PayloadAttributer_Expecter {
	return &PayloadAttributer_Expecter{mock: &_m.Mock}
}

// GetSuggestedFeeRecipient provides a mock function with given fields:
func (_m *PayloadAttributer) GetSuggestedFeeRecipient() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSuggestedFeeRecipient")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// PayloadAttributer_GetSuggestedFeeRecipient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSuggestedFeeRecipient'
type PayloadAttributer_GetSuggestedFeeRecipient_Call struct {
	*mock.Call
}

// GetSuggestedFeeRecipient is a helper method to define mock.On call
func (_e *PayloadAttributer_Expecter) GetSuggestedFeeRecipient() *PayloadAttributer_GetSuggestedFeeRecipient_Call {
	return &PayloadAttributer_GetSuggestedFeeRecipient_Call{Call: _e.mock.On("GetSuggestedFeeRecipient")}
}

func (_c *PayloadAttributer_GetSuggestedFeeRecipient_Call) Run(run func()) *PayloadAttributer_GetSuggestedFeeRecipient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PayloadAttributer_GetSuggestedFeeRecipient_Call) Return(_a0 common.Address) *PayloadAttributer_GetSuggestedFeeRecipient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PayloadAttributer_GetSuggestedFeeRecipient_Call) RunAndReturn(run func() common.Address) *PayloadAttributer_GetSuggestedFeeRecipient_Call {
	_c.Call.Return(run)
	return _c
}

// IsNil provides a mock function with given fields:
func (_m *PayloadAttributer) IsNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PayloadAttributer_IsNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNil'
type PayloadAttributer_IsNil_Call struct {
	*mock.Call
}

// IsNil is a helper method to define mock.On call
func (_e *PayloadAttributer_Expecter) IsNil() *PayloadAttributer_IsNil_Call {
	return &PayloadAttributer_IsNil_Call{Call: _e.mock.On("IsNil")}
}

func (_c *PayloadAttributer_IsNil_Call) Run(run func()) *PayloadAttributer_IsNil_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PayloadAttributer_IsNil_Call) Return(_a0 bool) *PayloadAttributer_IsNil_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PayloadAttributer_IsNil_Call) RunAndReturn(run func() bool) *PayloadAttributer_IsNil_Call {
	_c.Call.Return(run)
	return _c
}

// Version provides a mock function with given fields:
func (_m *PayloadAttributer) Version() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// PayloadAttributer_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type PayloadAttributer_Version_Call struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
func (_e *PayloadAttributer_Expecter) Version() *PayloadAttributer_Version_Call {
	return &PayloadAttributer_Version_Call{Call: _e.mock.On("Version")}
}

func (_c *PayloadAttributer_Version_Call) Run(run func()) *PayloadAttributer_Version_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PayloadAttributer_Version_Call) Return(_a0 uint32) *PayloadAttributer_Version_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PayloadAttributer_Version_Call) RunAndReturn(run func() uint32) *PayloadAttributer_Version_Call {
	_c.Call.Return(run)
	return _c
}

// NewPayloadAttributer creates a new instance of PayloadAttributer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPayloadAttributer(t interface {
	mock.TestingT
	Cleanup(func())
}) *PayloadAttributer {
	mock := &PayloadAttributer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
