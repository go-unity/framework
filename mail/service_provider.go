package mail

import (
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/contracts/queue"
)

const Binding = "gounity.mail"

type ServiceProvider struct {
}

func (route *ServiceProvider) Register(app foundation.Application) {
	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig(), app.MakeQueue()), nil
	})
}

func (route *ServiceProvider) Boot(app foundation.Application) {
	app.MakeQueue().Register([]queue.Job{
		NewSendMailJob(app.MakeConfig()),
	})
}
