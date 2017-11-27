package ctrl

import (
	"jobworker/jobs"
	"model"
)

//运行任务(包括新增和重新启动)
func (this *Controller) start(request *Action) {
	//1: 查询数据，得到相关的实体数据
	//2: 看是否为新增，如果为新增看是否要下载文件(如果只是启动，还是要看文件夹是否存在等)
	//3: 如果有文件要下载，下载后要解压到指定的文件夹中
	//4: 构造Task结构体
	jobs.AddJob(&model.Task{
		Id:       "123456789-qwert",
		Name:     "testJob",
		CronSpec: "0 */1 * * * ?",
		Command:  "echo first",
	})
}

//停止任务
func (this *Controller) stop(id string) {
	//1: 查询数据，看是否存在此任务
	//2: 终止任务运行
	//3: 更改数据库的状态信息
}

//删除任务
func (this *Controller) delete(id string) {
	//1: 查询数据，看是否存在此任务
	//2: 终止任务运行
	//3: 删除数据库的任务
}
