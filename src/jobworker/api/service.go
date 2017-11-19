package api

import "jobworker/ctrl"

type ApiServerArg struct {
	Bind  string
}

type ApiServer struct {
	bind    	string
	controller 	*ctrl.Controller
}

func NewAPiServer(arg *ApiServerArg, contr *ctrl.Controller ) *ApiServer {
	apiserver := &ApiServer{
		bind:arg.Bind,
		controller: contr,
	}
	//InitRouter
	return apiserver
}
