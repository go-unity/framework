package facades

import (
	"github.com/go-unity/framework/contracts/queue"
)

func Queue() queue.Queue {
	return App().MakeQueue()
}
