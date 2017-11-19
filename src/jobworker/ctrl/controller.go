package ctrl

import "jobworker/storage"

type Controller struct {
	Storage *storage.DataStorage
}

func NewController(storage *storage.DataStorage) *Controller {
	return &Controller{
		Storage : storage,
	}
}
