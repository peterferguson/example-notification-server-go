// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	interfaces "github.com/xmtp/example-notification-server-go/pkg/interfaces"
)

// Delivery is an autogenerated mock type for the Delivery type
type Delivery struct {
	mock.Mock
}

// CanDeliver provides a mock function with given fields: req
func (_m *Delivery) CanDeliver(req interfaces.SendRequest) bool {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for CanDeliver")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(interfaces.SendRequest) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Send provides a mock function with given fields: ctx, req
func (_m *Delivery) Send(ctx context.Context, req interfaces.SendRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.SendRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDelivery creates a new instance of Delivery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDelivery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Delivery {
	mock := &Delivery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
