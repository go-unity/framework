package facades

import (
	"github.com/go-unity/framework/contracts/schedule"
)

func Schedule() schedule.Schedule {
	return App().MakeSchedule()
}
