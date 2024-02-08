package filesystem

import (
	configcontract "github.com/go-unity/framework/contracts/config"
	filesystemcontract "github.com/go-unity/framework/contracts/filesystem"
	"github.com/go-unity/framework/contracts/foundation"
)

const Binding = "gounity.filesystem"

var ConfigFacade configcontract.Config
var StorageFacade filesystemcontract.Storage

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewStorage(app.MakeConfig()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	StorageFacade = app.MakeStorage()
}
