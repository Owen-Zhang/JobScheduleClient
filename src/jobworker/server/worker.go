package server

import (
	"flag"
	"jobworker/api"
	"jobworker/ctrl"
	"jobworker/etc"
	"jobworker/jobs"
	"jobworker/storage"
	"fmt"
)

type JobWork struct {
	Controller *ctrl.Controller
	Storage    *storage.DataStorage
	Api        *api.ApiServer
}

func NewWorker() (*JobWork, error) {

	var etcfile string
	flag.StringVar(&etcfile, "f", "etc/worker.yml", "worker etc file.")
	flag.Parse()
	if err := etc.New(etcfile); err != nil {
		return nil, err
	}

	storagearg := etc.GetStorageArg()

	fmt.Printf("Hosts:%s; DBName:%s; Password:%s; User:%s; Port:%d",storagearg.Hosts, storagearg.DBName, storagearg.Password, storagearg.User, storagearg.Port)

	dataaccess, err := storage.NewDataStorage(storagearg)
	if err != nil {
		return nil, err
	}
	controller := ctrl.NewController(dataaccess)
	apiserver := api.NewAPiServer(etc.GetApiServerArg(), controller)
	jobs.NewCron(etc.GetCronArg(), dataaccess)

	job := &JobWork{
		Controller: controller,
		Storage:    dataaccess,
		Api:        apiserver,
	}
	return job, nil
}

func (s *JobWork) Start() error {
	go s.Controller.ListenTask()
	s.Api.StartUp()

	return nil
}

func (s *JobWork) Stop() error {
	defer func() {
		s.Storage.Close()
		s.Controller.Close()
	}()
	return nil
}
