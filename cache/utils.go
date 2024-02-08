package cache

import (
	"github.com/go-unity/framework/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
