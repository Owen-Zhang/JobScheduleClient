package api

func (this *ApiServer) InitRoute() {
	//将其包含在一个组中
	worker := this.s.Group("/worker")

	//服务器心跳检查worker的运行状态
	worker.POST("/ping", this.ping)

	//加入任务
	worker.POST("/starttask/:id", this.starttask)

	//停止任务
	worker.POST("stoptask/:id", this.stoptask)

	//删除任务
	worker.POST("deletetask/:id", this.deletetask)
	
	//運行任務
	worker.POST("runtask/:id", this.runtask)
}
