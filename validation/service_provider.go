package validation

import (
	consolecontract "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/validation/console"
)

const Binding = "gounity.validation"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewValidation(), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
