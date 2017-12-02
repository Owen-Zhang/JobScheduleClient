package ctrl

import (
	"jobworker/jobs"
	"model"
	"utils/system"
	"regexp"
)

//运行任务(包括新增和重新启动)
func (this *Controller) start(request *Action) {
	//1: 查询数据，得到相关的实体数据
	task := this.Storage.GetTaskById(request.Id)
	if task == nil {
		return
	}

	//2: 看是否为新增，如果为新增看是否要下载文件(如果只是启动，还是要看文件夹是否存在等)
	if request.ZipFileUrl != "" {
		filename := system.UrlFileName(request.ZipFileUrl)
		flag, _ := regexp.MatchString(`^.+\.zip$`, filename)
		if flag && filename != task.OldZipFile {
			//下载文件
			/*
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			f, err := os.Create("qq.exe")
			if err != nil {
				panic(err)
			}
			io.Copy(f, res.Body)
			*/
		}
	}

	//3: 如果有文件要下载，下载后要解压到指定的文件夹中
	//4: 构造Task结构体
	jobs.AddJob(&model.Task{
		Id:       12,
		Name:     "testJob",
		CronSpec: "0 */1 * * * ?",
		Command:  "echo first",
	})
}

//停止任务
func (this *Controller) stop(id int) {
	//1: 查询数据，看是否存在此任务
	//2: 终止任务运行
	//3: 更改数据库的状态信息
}

//删除任务
func (this *Controller) delete(id int) {
	//1: 查询数据，看是否存在此任务
	//2: 终止任务运行
	//3: 删除数据库的任务
}
