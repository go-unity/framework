package facades

import (
	"github.com/go-unity/framework/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
