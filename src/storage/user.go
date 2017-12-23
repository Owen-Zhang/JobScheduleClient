package storage

import "model"

func (*DataStorage) UpdateUser(fields ...string) error {
	return nil
}

func (*DataStorage) UserAdd(user *model.User) (int64, error) {
	return 0, nil
}

func (*DataStorage) UserGetById(id int) (*model.User, error) {
	return nil, nil
}

func (*DataStorage) UserGetByName(userName string) (*model.User, error) {
	return nil, nil
}

func (*DataStorage) UserUpdate(user *model.User, fields ...string) error {
	return nil
}
