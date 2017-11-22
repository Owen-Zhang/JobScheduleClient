package ctrl

import "model"

//新增任务
func (this *Controller) NewTask(request *model.TaskNew) bool {
	this.actionlist <- action{
		actionType: new,
		id:         request.Id,
		zipFileUrl: request.ZipFileUrl,
	}
	return true
}
