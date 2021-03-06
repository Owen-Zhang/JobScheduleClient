package api

import "jobworker/ctrl"
import "github.com/gin-gonic/gin"

type ApiServerArg struct {
	Bind string
}

type ApiServer struct {
	bind       string
	controller *ctrl.Controller
	s          *gin.Engine
}

func NewAPiServer(arg *ApiServerArg, contr *ctrl.Controller) *ApiServer {
	server := gin.Default()
	apiserver := &ApiServer{
		bind:       arg.Bind,
		controller: contr,
		s:          server,
	}
	apiserver.InitRoute()

	return apiserver
}

func (this *ApiServer) StartUp() {
	this.s.Run(this.bind)
}
