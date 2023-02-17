// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	config "github.com/danteay/go-cassandra/config"
	gocql "github.com/gocql/gocql"

	logging "github.com/danteay/go-cassandra/logging"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *Client) Config() config.Config {
	ret := _m.Called()

	var r0 config.Config
	if rf, ok := ret.Get(0).(func() config.Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Config)
	}

	return r0
}

// Debug provides a mock function with given fields:
func (_m *Client) Debug() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PrintFn provides a mock function with given fields:
func (_m *Client) PrintFn() logging.DebugPrint {
	ret := _m.Called()

	var r0 logging.DebugPrint
	if rf, ok := ret.Get(0).(func() logging.DebugPrint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logging.DebugPrint)
		}
	}

	return r0
}

// Restart provides a mock function with given fields:
func (_m *Client) Restart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Session provides a mock function with given fields:
func (_m *Client) Session() *gocql.Session {
	ret := _m.Called()

	var r0 *gocql.Session
	if rf, ok := ret.Get(0).(func() *gocql.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocql.Session)
		}
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
