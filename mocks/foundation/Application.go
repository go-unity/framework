// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	auth "github.com/go-unity/framework/contracts/auth"
	access "github.com/go-unity/framework/contracts/auth/access"

	cache "github.com/go-unity/framework/contracts/cache"

	config "github.com/go-unity/framework/contracts/config"

	console "github.com/go-unity/framework/contracts/console"

	context "context"

	crypt "github.com/go-unity/framework/contracts/crypt"

	event "github.com/go-unity/framework/contracts/event"

	filesystem "github.com/go-unity/framework/contracts/filesystem"

	foundation "github.com/go-unity/framework/contracts/foundation"

	grpc "github.com/go-unity/framework/contracts/grpc"

	hash "github.com/go-unity/framework/contracts/hash"

	http "github.com/go-unity/framework/contracts/http"

	log "github.com/go-unity/framework/contracts/log"

	mail "github.com/go-unity/framework/contracts/mail"

	mock "github.com/stretchr/testify/mock"

	orm "github.com/go-unity/framework/contracts/database/orm"

	queue "github.com/go-unity/framework/contracts/queue"

	route "github.com/go-unity/framework/contracts/route"

	schedule "github.com/go-unity/framework/contracts/schedule"

	seeder "github.com/go-unity/framework/contracts/database/seeder"

	testing "github.com/go-unity/framework/contracts/testing"

	translation "github.com/go-unity/framework/contracts/translation"

	validation "github.com/go-unity/framework/contracts/validation"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// BasePath provides a mock function with given fields: path
func (_m *Application) BasePath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Bind provides a mock function with given fields: key, callback
func (_m *Application) Bind(key interface{}, callback func(foundation.Application) (interface{}, error)) {
	_m.Called(key, callback)
}

// BindWith provides a mock function with given fields: key, callback
func (_m *Application) BindWith(key interface{}, callback func(foundation.Application, map[string]interface{}) (interface{}, error)) {
	_m.Called(key, callback)
}

// Boot provides a mock function with given fields:
func (_m *Application) Boot() {
	_m.Called()
}

// Commands provides a mock function with given fields: _a0
func (_m *Application) Commands(_a0 []console.Command) {
	_m.Called(_a0)
}

// ConfigPath provides a mock function with given fields: path
func (_m *Application) ConfigPath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DatabasePath provides a mock function with given fields: path
func (_m *Application) DatabasePath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLocale provides a mock function with given fields: ctx
func (_m *Application) GetLocale(ctx context.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Instance provides a mock function with given fields: key, instance
func (_m *Application) Instance(key interface{}, instance interface{}) {
	_m.Called(key, instance)
}

// IsLocale provides a mock function with given fields: ctx, locale
func (_m *Application) IsLocale(ctx context.Context, locale string) bool {
	ret := _m.Called(ctx, locale)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, locale)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LangPath provides a mock function with given fields: path
func (_m *Application) LangPath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Make provides a mock function with given fields: key
func (_m *Application) Make(key interface{}) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeArtisan provides a mock function with given fields:
func (_m *Application) MakeArtisan() console.Artisan {
	ret := _m.Called()

	var r0 console.Artisan
	if rf, ok := ret.Get(0).(func() console.Artisan); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(console.Artisan)
		}
	}

	return r0
}

// MakeAuth provides a mock function with given fields: ctx
func (_m *Application) MakeAuth(ctx http.Context) auth.Auth {
	ret := _m.Called(ctx)

	var r0 auth.Auth
	if rf, ok := ret.Get(0).(func(http.Context) auth.Auth); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.Auth)
		}
	}

	return r0
}

// MakeCache provides a mock function with given fields:
func (_m *Application) MakeCache() cache.Cache {
	ret := _m.Called()

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func() cache.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}

// MakeConfig provides a mock function with given fields:
func (_m *Application) MakeConfig() config.Config {
	ret := _m.Called()

	var r0 config.Config
	if rf, ok := ret.Get(0).(func() config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Config)
		}
	}

	return r0
}

// MakeCrypt provides a mock function with given fields:
func (_m *Application) MakeCrypt() crypt.Crypt {
	ret := _m.Called()

	var r0 crypt.Crypt
	if rf, ok := ret.Get(0).(func() crypt.Crypt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypt.Crypt)
		}
	}

	return r0
}

// MakeEvent provides a mock function with given fields:
func (_m *Application) MakeEvent() event.Instance {
	ret := _m.Called()

	var r0 event.Instance
	if rf, ok := ret.Get(0).(func() event.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Instance)
		}
	}

	return r0
}

// MakeGate provides a mock function with given fields:
func (_m *Application) MakeGate() access.Gate {
	ret := _m.Called()

	var r0 access.Gate
	if rf, ok := ret.Get(0).(func() access.Gate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Gate)
		}
	}

	return r0
}

// MakeGrpc provides a mock function with given fields:
func (_m *Application) MakeGrpc() grpc.Grpc {
	ret := _m.Called()

	var r0 grpc.Grpc
	if rf, ok := ret.Get(0).(func() grpc.Grpc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.Grpc)
		}
	}

	return r0
}

// MakeHash provides a mock function with given fields:
func (_m *Application) MakeHash() hash.Hash {
	ret := _m.Called()

	var r0 hash.Hash
	if rf, ok := ret.Get(0).(func() hash.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hash.Hash)
		}
	}

	return r0
}

// MakeLang provides a mock function with given fields: ctx
func (_m *Application) MakeLang(ctx context.Context) translation.Translator {
	ret := _m.Called(ctx)

	var r0 translation.Translator
	if rf, ok := ret.Get(0).(func(context.Context) translation.Translator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(translation.Translator)
		}
	}

	return r0
}

// MakeLog provides a mock function with given fields:
func (_m *Application) MakeLog() log.Log {
	ret := _m.Called()

	var r0 log.Log
	if rf, ok := ret.Get(0).(func() log.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Log)
		}
	}

	return r0
}

// MakeMail provides a mock function with given fields:
func (_m *Application) MakeMail() mail.Mail {
	ret := _m.Called()

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func() mail.Mail); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// MakeOrm provides a mock function with given fields:
func (_m *Application) MakeOrm() orm.Orm {
	ret := _m.Called()

	var r0 orm.Orm
	if rf, ok := ret.Get(0).(func() orm.Orm); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Orm)
		}
	}

	return r0
}

// MakeQueue provides a mock function with given fields:
func (_m *Application) MakeQueue() queue.Queue {
	ret := _m.Called()

	var r0 queue.Queue
	if rf, ok := ret.Get(0).(func() queue.Queue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Queue)
		}
	}

	return r0
}

// MakeRateLimiter provides a mock function with given fields:
func (_m *Application) MakeRateLimiter() http.RateLimiter {
	ret := _m.Called()

	var r0 http.RateLimiter
	if rf, ok := ret.Get(0).(func() http.RateLimiter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.RateLimiter)
		}
	}

	return r0
}

// MakeRoute provides a mock function with given fields:
func (_m *Application) MakeRoute() route.Route {
	ret := _m.Called()

	var r0 route.Route
	if rf, ok := ret.Get(0).(func() route.Route); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Route)
		}
	}

	return r0
}

// MakeSchedule provides a mock function with given fields:
func (_m *Application) MakeSchedule() schedule.Schedule {
	ret := _m.Called()

	var r0 schedule.Schedule
	if rf, ok := ret.Get(0).(func() schedule.Schedule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schedule.Schedule)
		}
	}

	return r0
}

// MakeSeeder provides a mock function with given fields:
func (_m *Application) MakeSeeder() seeder.Facade {
	ret := _m.Called()

	var r0 seeder.Facade
	if rf, ok := ret.Get(0).(func() seeder.Facade); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(seeder.Facade)
		}
	}

	return r0
}

// MakeStorage provides a mock function with given fields:
func (_m *Application) MakeStorage() filesystem.Storage {
	ret := _m.Called()

	var r0 filesystem.Storage
	if rf, ok := ret.Get(0).(func() filesystem.Storage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.Storage)
		}
	}

	return r0
}

// MakeTesting provides a mock function with given fields:
func (_m *Application) MakeTesting() testing.Testing {
	ret := _m.Called()

	var r0 testing.Testing
	if rf, ok := ret.Get(0).(func() testing.Testing); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.Testing)
		}
	}

	return r0
}

// MakeValidation provides a mock function with given fields:
func (_m *Application) MakeValidation() validation.Validation {
	ret := _m.Called()

	var r0 validation.Validation
	if rf, ok := ret.Get(0).(func() validation.Validation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validation)
		}
	}

	return r0
}

// MakeView provides a mock function with given fields:
func (_m *Application) MakeView() http.View {
	ret := _m.Called()

	var r0 http.View
	if rf, ok := ret.Get(0).(func() http.View); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.View)
		}
	}

	return r0
}

// MakeWith provides a mock function with given fields: key, parameters
func (_m *Application) MakeWith(key interface{}, parameters map[string]interface{}) (interface{}, error) {
	ret := _m.Called(key, parameters)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}) (interface{}, error)); ok {
		return rf(key, parameters)
	}
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}) interface{}); ok {
		r0 = rf(key, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, map[string]interface{}) error); ok {
		r1 = rf(key, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Path provides a mock function with given fields: path
func (_m *Application) Path(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PublicPath provides a mock function with given fields: path
func (_m *Application) PublicPath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Publishes provides a mock function with given fields: packageName, paths, groups
func (_m *Application) Publishes(packageName string, paths map[string]string, groups ...string) {
	_va := make([]interface{}, len(groups))
	for _i := range groups {
		_va[_i] = groups[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, packageName, paths)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetLocale provides a mock function with given fields: ctx, locale
func (_m *Application) SetLocale(ctx context.Context, locale string) context.Context {
	ret := _m.Called(ctx, locale)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, locale)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Singleton provides a mock function with given fields: key, callback
func (_m *Application) Singleton(key interface{}, callback func(foundation.Application) (interface{}, error)) {
	_m.Called(key, callback)
}

// StoragePath provides a mock function with given fields: path
func (_m *Application) StoragePath(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *Application) Version() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewApplication creates a new instance of Application. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApplication(t interface {
	mock.TestingT
	Cleanup(func())
}) *Application {
	mock := &Application{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
