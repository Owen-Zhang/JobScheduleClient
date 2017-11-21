package api

func (this *ApiServer) InitRoute() {
	worker := this.s.Group("/worker")

	worker.POST("/ping", ping)
}
