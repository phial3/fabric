// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SignerSerializer is an autogenerated mock type for the SignerSerializer type
type SignerSerializer struct {
	mock.Mock
}

// Serialize provides a mock function with given fields:
func (_m *SignerSerializer) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
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

// Sign provides a mock function with given fields: message
func (_m *SignerSerializer) Sign(message []byte) ([]byte, error) {
	ret := _m.Called(message)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSignerSerializer interface {
	mock.TestingT
	Cleanup(func())
}

// NewSignerSerializer creates a new instance of SignerSerializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSignerSerializer(t mockConstructorTestingTNewSignerSerializer) *SignerSerializer {
	mock := &SignerSerializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
