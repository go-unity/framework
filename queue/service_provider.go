package queue

import (
	"github.com/go-unity/framework/contracts/console"
	"github.com/go-unity/framework/contracts/foundation"
	queueConsole "github.com/go-unity/framework/queue/console"
)

const Binding = "gounity.queue"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]console.Command{
		&queueConsole.JobMakeCommand{},
	})
}
