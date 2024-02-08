// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	testing "github.com/go-unity/framework/contracts/testing"
	mock "github.com/stretchr/testify/mock"
)

// Docker is an autogenerated mock type for the Docker type
type Docker struct {
	mock.Mock
}

// Database provides a mock function with given fields: connection
func (_m *Docker) Database(connection ...string) (testing.Database, error) {
	_va := make([]interface{}, len(connection))
	for _i := range connection {
		_va[_i] = connection[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 testing.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(...string) (testing.Database, error)); ok {
		return rf(connection...)
	}
	if rf, ok := ret.Get(0).(func(...string) testing.Database); ok {
		r0 = rf(connection...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(connection...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDocker creates a new instance of Docker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocker(t interface {
	mock.TestingT
	Cleanup(func())
}) *Docker {
	mock := &Docker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
