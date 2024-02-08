package facades

import (
	"github.com/go-unity/framework/contracts/route"
)

func Route() route.Route {
	return App().MakeRoute()
}
