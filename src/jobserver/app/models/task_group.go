package models

type TaskGroup struct {
	Id          int
	UserId      int
	GroupName   string
	Description string
	CreateTime  int64
}

func (t *TaskGroup) Update(fields ...string) error {
	return nil
}

func TaskGroupAdd(obj *TaskGroup) (int64, error) {
	return 0, nil
}

func TaskGroupGetById(id int) (*TaskGroup, error) {
	return nil, nil
}

func TaskGroupDelById(id int) error {
	return nil
}

func TaskGroupGetList(page, pageSize int) ([]*TaskGroup, int64) {
	return nil, 0
}
