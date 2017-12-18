package models

type TaskLog struct {
	Id          int
	TaskId      int
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}

func TaskLogGetList(page, pageSize int, filters ...interface{}) ([]*TaskLog, int64) {
	return nil, 0
}

func TaskLogGetById(id int) (*TaskLog, error) {
	return nil, nil
}

func TaskLogDelById(id int) error {
	return nil
}
