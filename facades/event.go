package facades

import "github.com/go-unity/framework/contracts/event"

func Event() event.Instance {
	return App().MakeEvent()
}
