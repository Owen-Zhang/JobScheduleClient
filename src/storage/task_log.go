package storage

import "model"

func (*DataStorage) TaskLogGetList(page, pageSize int) ([]*model.TaskLog, int64) {
	return nil, 0
}

func (*DataStorage) TaskLogGetById(id int) (*model.TaskLog, error) {
	return nil, nil
}

func (*DataStorage) TaskLogDelById(id int) error {
	return nil
}
