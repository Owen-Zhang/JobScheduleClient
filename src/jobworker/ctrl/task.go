package ctrl

import "model"

//新增任务
func (this *Controller) New(request *model.TaskNew) bool {
	this.actionlist <- action{
		actionType: new,
		id:         request.Id,
		zipFileUrl: request.ZipFileUrl,
	}
	return true
}

//运行任务
func (this *Controller) Start(request *model.TaskNew) bool {
	this.actionlist <- action{
		actionType: start,
		id:         request.Id,
		zipFileUrl: request.ZipFileUrl,
	}
	return true
}

//停止任务
func (this *Controller) Stop(id string) bool {
	this.actionlist <- action{
		actionType: stop,
		id:         id,
	}
	return true
}

//删除任务
func (this *Controller) Delete(id string) bool {
	this.actionlist <- action{
		actionType: delete,
		id:         id,
	}
	return true
}
