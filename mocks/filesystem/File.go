// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	filesystem "github.com/go-unity/framework/contracts/filesystem"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

// Disk provides a mock function with given fields: disk
func (_m *File) Disk(disk string) filesystem.File {
	ret := _m.Called(disk)

	var r0 filesystem.File
	if rf, ok := ret.Get(0).(func(string) filesystem.File); ok {
		r0 = rf(disk)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.File)
		}
	}

	return r0
}

// Extension provides a mock function with given fields:
func (_m *File) Extension() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File provides a mock function with given fields:
func (_m *File) File() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetClientOriginalExtension provides a mock function with given fields:
func (_m *File) GetClientOriginalExtension() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetClientOriginalName provides a mock function with given fields:
func (_m *File) GetClientOriginalName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// HashName provides a mock function with given fields: path
func (_m *File) HashName(path ...string) string {
	_va := make([]interface{}, len(path))
	for _i := range path {
		_va[_i] = path[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(path...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LastModified provides a mock function with given fields:
func (_m *File) LastModified() (time.Time, error) {
	ret := _m.Called()

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func() (time.Time, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MimeType provides a mock function with given fields:
func (_m *File) MimeType() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Size provides a mock function with given fields:
func (_m *File) Size() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: path
func (_m *File) Store(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreAs provides a mock function with given fields: path, name
func (_m *File) StoreAs(path string, name string) (string, error) {
	ret := _m.Called(path, name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(path, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(path, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFile creates a new instance of File. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *File {
	mock := &File{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
