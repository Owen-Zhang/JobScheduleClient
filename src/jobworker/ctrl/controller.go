package ctrl

import "jobworker/storage"

const (
	new    = 1
	start  = 2
	stop   = 3
	delete = 4
)

type Controller struct {
	actionlist chan action
	Storage    *storage.DataStorage
}

type action struct {
	actionType int    //
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

func (this *Controller) Close() {
	close(this.actionlist)
}
