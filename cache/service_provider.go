package cache

import (
	"github.com/go-unity/framework/cache/console"
	contractsconsole "github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
)

const Binding = "gounity.cache"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		config := app.MakeConfig()
		log := app.MakeLog()
		store := config.GetString("cache.default")

		return NewCacheStorage(config, log, store)
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]contractsconsole.Command{
		console.NewClearCommand(app.MakeCache()),
	})
}
