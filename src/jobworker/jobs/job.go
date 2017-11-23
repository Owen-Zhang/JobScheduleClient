package jobs

import (
	"time"
	"model"
	"fmt"
)

type Job struct {
	id         string                                            // 任务ID
	//logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *model.Task                                       // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	concurrent bool                                              // 同一个任务是否允许并行执行
}

func (this *Job) Run() {
	if !this.concurrent && this.status > 0 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			//此处最好写日志
			fmt.Printf("Run wrong is : %s", err)
		}
	}()

	//此处是为了控制同时运行任务的个数
	if workPool != nil {
		workPool <- true
		defer func() {
			<- workPool
		}()
	}

	this.status++
	defer func() {
		this.status--
		if this.status < 0 {
			this.status = 0
		}
	}()

	t := time.Now()
	timeout := time.Duration(time.Hour * 24)
	if this.task.TimeOut > 0 {
		timeout = time.Second * time.Duration(this.task.TimeOut)
	}
	cmdOut, cmdErr, err, isTimeout := this.runFunc(timeout)
	ut := time.Now().Sub(t) / time.Millisecond

	//写日志
	//发邮件
	//更新任务执行时间等
}
