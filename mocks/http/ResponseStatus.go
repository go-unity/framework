// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	http "github.com/go-unity/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// ResponseStatus is an autogenerated mock type for the ResponseStatus type
type ResponseStatus struct {
	mock.Mock
}

// Data provides a mock function with given fields: contentType, data
func (_m *ResponseStatus) Data(contentType string, data []byte) http.Response {
	ret := _m.Called(contentType, data)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, []byte) http.Response); ok {
		r0 = rf(contentType, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Json provides a mock function with given fields: obj
func (_m *ResponseStatus) Json(obj interface{}) http.Response {
	ret := _m.Called(obj)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(interface{}) http.Response); ok {
		r0 = rf(obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// String provides a mock function with given fields: format, values
func (_m *ResponseStatus) String(format string, values ...interface{}) http.Response {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, ...interface{}) http.Response); ok {
		r0 = rf(format, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// NewResponseStatus creates a new instance of ResponseStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponseStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResponseStatus {
	mock := &ResponseStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
