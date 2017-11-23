package jobs

import (
	"github.com/robfig/cron"
	"sync"
	"jobworker/ctrl"
)

type CronArg struct {
	PoolSize int32
}

var (
	workPool 	chan bool
	mainCron 	*cron.Cron
	controller  *ctrl.Controller
	)

func NewCron(arg *CronArg, contr *ctrl.Controller) {
	if arg.PoolSize > 0 {
		workPool = make(chan bool, arg.PoolSize)
	}

	controller = contr
	mainCron := cron.New()
	mainCron.Start()
}

func AddJob(spec string, job *Job) bool {
	lock := sync.Mutex{}

	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) == nil {
		return false
	}

	if err := mainCron.AddJob(spec, job); err == nil {
		return true
	}
	return false
}

func GetEntryById(id string) *cron.Entry {
	entries :=  mainCron.Entries()
	for _, e := range entries{
		if v, flag := e.Job.(*Job); flag {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}