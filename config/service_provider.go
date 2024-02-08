package config

import (
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/support"
)

const Binding = "gounity.config"

type ServiceProvider struct {
}

func (config *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewConfig(support.EnvPath), nil
	})
}

func (config *ServiceProvider) Boot(app foundation.Application) {

}
