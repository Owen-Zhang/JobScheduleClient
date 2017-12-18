package healthy

import (
	"github.com/robfig/cron"
	"sync"
	"model"
)

var (
	mainCron *cron.Cron
	lock     sync.Mutex
)

func init()  {
	mainCron = cron.New()
	mainCron.Start()
}

/*
//加载worker，监控worker
func InitHealthCheck(spec string) {
	list, _ := models.TaskGetList(1, 1000000, "status", 1)
	for _, task := range list {
		AddHealthyCheck(spec, nil)
	}
}*/

//增加心跳任务，检查worker机子是否正常运行
func AddHealthyCheck(spec string, info *model.HealthInfo) bool {
	heal, err := newHealth(info)
	if err != nil {
		return false
	}

	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	if ExistJob(heal.id) {
		return false
	}

	if err := mainCron.AddJob(spec, heal); err == nil {
		return true
	}
	return false
}

//删除运行中的任务
func RemoveJob(id int) {
	if !ExistJob(id) {
		return
	}
	mainCron.RemoveJob(func(e *cron.Entry) bool {
		if v, flag := e.Job.(*Health); flag {
			if v.id == id {
				return true
			}
		}
		return false
	})
}

//判断任务是否在指行队列中
func ExistJob(id int) bool  {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, flag := e.Job.(*Health); flag {
			if v.id == id {
				return true
			}
		}
	}
	return false
}