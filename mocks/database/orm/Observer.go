// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	orm "github.com/go-unity/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// Observer is an autogenerated mock type for the Observer type
type Observer struct {
	mock.Mock
}

// Created provides a mock function with given fields: _a0
func (_m *Observer) Created(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Creating provides a mock function with given fields: _a0
func (_m *Observer) Creating(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deleted provides a mock function with given fields: _a0
func (_m *Observer) Deleted(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deleting provides a mock function with given fields: _a0
func (_m *Observer) Deleting(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDeleted provides a mock function with given fields: _a0
func (_m *Observer) ForceDeleted(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDeleting provides a mock function with given fields: _a0
func (_m *Observer) ForceDeleting(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Retrieved provides a mock function with given fields: _a0
func (_m *Observer) Retrieved(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Saved provides a mock function with given fields: _a0
func (_m *Observer) Saved(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Saving provides a mock function with given fields: _a0
func (_m *Observer) Saving(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updated provides a mock function with given fields: _a0
func (_m *Observer) Updated(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updating provides a mock function with given fields: _a0
func (_m *Observer) Updating(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewObserver creates a new instance of Observer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObserver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Observer {
	mock := &Observer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}