package jobs

import (
	"github.com/robfig/cron"
)

type CronArg struct {
	PoolSize int32
}

type CronService struct {
	mainCron *cron.Cron
	workPool chan bool
}

func NewCron(arg *CronArg) * CronService {
	var workPoll chan bool
	if arg.PoolSize > 0 {
		workPoll = make(chan bool, arg.PoolSize)
	}

	mainCron := cron.New()
	mainCron.Start()

	result := &CronService{
		mainCron :mainCron,
		workPool :workPoll,
	}
	return result
}

