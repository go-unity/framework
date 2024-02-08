package auth

import (
	"context"

	"github.com/go-unity/framework/auth/access"
	"github.com/go-unity/framework/auth/console"
	contractconsole "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/contracts/http"
)

const BindingAuth = "gounity.auth"
const BindingGate = "gounity.gate"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.BindWith(BindingAuth, func(app foundation.Application, parameters map[string]any) (any, error) {
		config := app.MakeConfig()
		return NewAuth(config.GetString("auth.defaults.guard"),
			app.MakeCache(), config, parameters["ctx"].(http.Context), app.MakeOrm()), nil
	})
	app.Singleton(BindingGate, func(app foundation.Application) (any, error) {
		return access.NewGate(context.Background()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]contractconsole.Command{
		console.NewJwtSecretCommand(app.MakeConfig()),
		console.NewPolicyMakeCommand(),
	})
}
