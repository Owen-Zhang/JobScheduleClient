package jobs

import (
	"fmt"
	"jobworker/storage"
	"model"
	"sync"

	"github.com/robfig/cron"
)

type CronArg struct {
	PoolSize int32
}

var (
	workPool chan bool
	mainCron *cron.Cron
	data     *storage.DataStorage
)

func NewCron(arg *CronArg, storage *storage.DataStorage) {
	if arg.PoolSize > 0 {
		workPool = make(chan bool, arg.PoolSize)
	}

	data = storage
	mainCron = cron.New()
	mainCron.Start()
	fmt.Println("cron started")
}

//增加任务
func AddJob(task *model.Task) bool {
	job, err := newJobFromTask(task)
	if err != nil {
		return false
	}

	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	if getEntryById(job.id) != nil {
		return false
	}

	if err := mainCron.AddJob(task.CronSpec, job); err == nil {
		return true
	}
	return false
}

func getEntryById(id string) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, flag := e.Job.(*Job); flag {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}
