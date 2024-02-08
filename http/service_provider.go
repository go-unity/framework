package http

import (
	"github.com/go-unity/framework/contracts/cache"
	"github.com/go-unity/framework/contracts/config"
	consolecontract "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/contracts/http"
	"github.com/go-unity/framework/http/console"
)

const BindingRateLimiter = "gounity.rate_limiter"
const BindingView = "gounity.view"

type ServiceProvider struct{}

var (
	CacheFacade       cache.Cache
	ConfigFacade      config.Config
	RateLimiterFacade http.RateLimiter
)

func (http *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(BindingRateLimiter, func(app foundation.Application) (any, error) {
		return NewRateLimiter(), nil
	})
	app.Singleton(BindingView, func(app foundation.Application) (any, error) {
		return NewView(), nil
	})
}

func (http *ServiceProvider) Boot(app foundation.Application) {
	CacheFacade = app.MakeCache()
	ConfigFacade = app.MakeConfig()
	RateLimiterFacade = app.MakeRateLimiter()

	http.registerCommands(app)
}

func (http *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
