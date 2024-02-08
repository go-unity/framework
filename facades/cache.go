package facades

import (
	"github.com/go-unity/framework/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
