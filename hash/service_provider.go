package hash

import (
	"github.com/go-unity/framework/contracts/foundation"
)

const Binding = "gounity.hash"

type ServiceProvider struct {
}

func (hash *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (hash *ServiceProvider) Boot(app foundation.Application) {

}
