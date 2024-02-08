package testing

import (
	contractsconsole "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
)

const Binding = "gounity.testing"

var artisanFacades contractsconsole.Artisan

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	artisanFacades = app.MakeArtisan()
}
