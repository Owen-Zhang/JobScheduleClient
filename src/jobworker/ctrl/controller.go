package ctrl

import (
	"fmt"
	"jobworker/jobs"
	"jobworker/storage"
	"model"
	"time"
)

const (
	new    = 1
	start  = 2
	stop   = 3
	delete = 4
)

type Controller struct {
	Ticker     *time.Ticker
	actionlist chan action
	Storage    *storage.DataStorage
}

type action struct {
	actionType int    //操作类型
	id         string //任务的主键
	zipFileUrl string //zip文件的下载地址
}

func NewController(storage *storage.DataStorage) *Controller {
	list := make(chan action, 10)
	return &Controller{
		Storage:    storage,
		actionlist: list,
	}
}

func (this *Controller) ListenTask() {
NEW_TICK_DURATION:
	this.Ticker = time.NewTicker(time.Second * 1)
	for {
		select {
		case newtask := <-this.actionlist:
			fmt.Printf("ListenTask task id is : %s \n", newtask.id)
			jobs.AddJob(&model.Task{
				Id:       "123456789-qwert",
				Name:     "testJob",
				CronSpec: "0 */1 * * * ?",
				Command:  "echo first",
			})
			this.Ticker.Stop()
			goto NEW_TICK_DURATION
		}
	}
}

func (this *Controller) Close() {
	close(this.actionlist)
}
