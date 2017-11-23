package ctrl

import (
	"fmt"
	"jobworker/jobs"
	"jobworker/storage"
	"time"
)

const (
	new    = 1
	start  = 2
	stop   = 3
	delete = 4
)

type Controller struct {
	Ticker      *time.Ticker
	actionlist  chan action
	cronservice *jobs.CronService
	Storage     *storage.DataStorage
}

type action struct {
	actionType int    //操作类型
	id         string //任务的主键
	zipFileUrl string //zip文件的下载地址
}

func NewController(storage *storage.DataStorage, cronarg *jobs.CronArg) *Controller {
	list := make(chan action, 10)
	cronservice := jobs.NewCron(cronarg)

	return &Controller{
		Storage:     storage,
		actionlist:  list,
		cronservice: cronservice,
	}
}

func (this *Controller) ListenTask() {
NEW_TICK_DURATION:
	this.Ticker = time.NewTicker(time.Second * 1)
	for {
		select {
		case newtask := <-this.actionlist:
			fmt.Println(newtask.id)
			this.Ticker.Stop()
			goto NEW_TICK_DURATION
		}
	}
}

func (this *Controller) Close() {
	close(this.actionlist)
}
