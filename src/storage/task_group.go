package storage

import (
	"model"
)

func (this *DataStorage) UpdateGroup(fields ...string) error {
	return nil
}

func (this *DataStorage) TaskGroupAdd(obj *model.TaskGroup) (int64, error) {
	return 0, nil
}

func (this *DataStorage) TaskGroupGetById(id int) (*model.TaskGroup, error) {
	return nil, nil
}

func (this *DataStorage) TaskGroupDelById(id int) error {
	return nil
}

func (this *DataStorage) TaskGroupGetList(page, pageSize int) ([]*model.TaskGroup, int64) {
	return nil, 0
}
