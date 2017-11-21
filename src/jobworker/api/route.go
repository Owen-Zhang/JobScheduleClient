package api

func (this *ApiServer) InitRoute() {
	//将其包含在一个组中
	worker := this.s.Group("/worker")

	//服务器心跳检查worker的运行状态
	worker.POST("/ping", this.ping)

	//新增任务
	worker.POST("/newtask", this.newtask)

	//运行任务
	worker.POST("/starttask", this.starttask)

	//停止任务
	worker.POST("stoptask", this.stoptask)

	//删除任务
	worker.POST("deletetask", this.deletetask)
}
