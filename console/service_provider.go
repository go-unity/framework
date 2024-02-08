package console

import (
	"github.com/go-unity/framework/console/console"
	consolecontract "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
)

const Binding = "gounity.console"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewConsole(), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	artisan := app.MakeArtisan()
	artisan.Register([]consolecontract.Command{
		console.NewListCommand(artisan),
		console.NewMakeCommand(),
	})
}
