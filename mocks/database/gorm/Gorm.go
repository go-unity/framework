// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// Gorm is an autogenerated mock type for the Gorm type
type Gorm struct {
	mock.Mock
}

// Make provides a mock function with given fields:
func (_m *Gorm) Make() (*gorm.DB, error) {
	ret := _m.Called()

	var r0 *gorm.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gorm.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGorm creates a new instance of Gorm. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGorm(t interface {
	mock.TestingT
	Cleanup(func())
}) *Gorm {
	mock := &Gorm{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
