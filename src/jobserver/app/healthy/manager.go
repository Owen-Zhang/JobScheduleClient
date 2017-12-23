package healthy

import (
	"time"
	"fmt"
	"model"
	"net/http"
)

type Health struct {
id 		   int                              // 任务ID
name       string                           // 任务名称
health     *model.HealthInfo                // 心跳对象信息
runFunc    func(time.Duration) bool			// 执行函数
count      int                              // check faild次数(等于三次就表示worker不能用了，报告状态，启用其它的机器)
faild      bool								// 标识某个worker不可用
}

//新增心跳任务
func newHealth(heart *model.HealthInfo) (*Health, error) {
	hea := &Health{
		id:     heart.Id,
		name:   heart.Name,
		count:  0,
		faild:  false,
		health: heart,
	}
	hea.runFunc = func(timeout time.Duration) bool {
		return runWithTimeOut(heart.Url, heart.Port, timeout)
	}
	return hea, nil
}

func (this *Health) Run() {
	if this.faild {
		return
	}

	timeout := time.Duration(time.Second * 5)
	flag := this.runFunc(timeout)

	if !flag {
		this.count++
		fmt.Printf("check [%s:%d] faild", this.health.Url, this.health.Port)
	} else {
		this.count--
	}

	//此worker不可用, 提示相关信息，并将任务分配给其它的worker
	if this.count >= 3 {
		fmt.Printf("[%s:%d] 任务将要转移到其它worker", this.health.Url, this.health.Port)

		//1: 处理转移任务

		//2: 标识此worker不可用
		this.faild = true
	}
}

func runWithTimeOut(url string, port int, timeout time.Duration) bool {
	client := http.Client{}
	client.Timeout = timeout
	respose, err := client.Post(fmt.Sprintf("%s:%d/ping", url, port), "", nil)
	if err != nil {
		return false
	}

	return respose.StatusCode == http.StatusOK
}
