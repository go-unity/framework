// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConnectionModel is an autogenerated mock type for the ConnectionModel type
type ConnectionModel struct {
	mock.Mock
}

// Connection provides a mock function with given fields:
func (_m *ConnectionModel) Connection() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewConnectionModel creates a new instance of ConnectionModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectionModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectionModel {
	mock := &ConnectionModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}