package server

import (
	"ctrl",
	"storage"
)

type JobWork struct {
	Controller *ctrl.Controller,
	Storage    *storage.DataStorage,
	//ServerTask *task.
}

func NewWorker(*JobWork, error) {
	job := &JobWork{
		
	}
	return job,nil
}

func (s *JobWork) Start() error {
	
	return nil
}

func (s *JobWork) Stop() error {
	
	return nil
}
