package facades

import (
	"github.com/go-unity/framework/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
