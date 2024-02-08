package facades

import (
	"github.com/go-unity/framework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
