package ctrl

import (
	"jobworker/storage"
	"time"
)

type Controller struct {
	Ticker     *time.Ticker
	Actionlist chan Action
	Storage    *storage.DataStorage
	ExeConfig  *ExeConfig
}

//外部接口传入的任务实体(chan实体)
type Action struct {
	ActionType int    	//操作类型
	Id         int 		//任务的主键
}

//client程序的相关配制信息
type ExeConfig struct {
	ClientPath    string //client程序的运行路径
	TempZipFolder string //zip文件下载后的暂存路径
	TaskFolder    string //每个任务的程序目录
}


func NewController(storage *storage.DataStorage, config *ExeConfig) *Controller {
	list := make(chan Action, 10)
	return &Controller{
		Storage:    storage,
		Actionlist: list,
		ExeConfig:  config,
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
