package testing

import (
	"github.com/go-unity/framework/contracts/foundation"
	"github.com/go-unity/framework/contracts/testing"
	"github.com/go-unity/framework/testing/docker"
)

type Application struct {
	app foundation.Application
}

func NewApplication(app foundation.Application) *Application {
	return &Application{
		app: app,
	}
}

func (receiver *Application) Docker() testing.Docker {
	return docker.NewDocker(receiver.app)
}
