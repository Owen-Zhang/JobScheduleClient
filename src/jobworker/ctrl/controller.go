package ctrl

import (
	"jobworker/storage"
	"time"
)

type Controller struct {
	Ticker     *time.Ticker
	Actionlist chan Action
	Storage    *storage.DataStorage
}

type Action struct {
	ActionType int    //操作类型
	Id         string //任务的主键
	ZipFileUrl string //zip文件的下载地址
}

func NewController(storage *storage.DataStorage) *Controller {
	list := make(chan Action, 10)
	return &Controller{
		Storage:    storage,
		Actionlist: list,
	}
}

func (this *Controller) ListenTask() {
NEW_TICK_DURATION:
	this.Ticker = time.NewTicker(time.Second * 1)
	for {
		select {
		case newtask := <-this.Actionlist:
			this.Ticker.Stop()

			actiontype := newtask.ActionType
			switch actiontype {
			case 1,2:
				this.start(&newtask)
			case 3:
				this.stop(newtask.Id)
			case 4:
				this.delete(newtask.Id)
			}

			goto NEW_TICK_DURATION
		}
	}
}

func (this *Controller) Close() {
	close(this.Actionlist)
}
