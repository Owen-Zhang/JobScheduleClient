package models

type Task struct {
	Id           int
	UserId       int
	GroupId      int
	TaskName     string
	TaskType     int
	Description  string
	CronSpec     string
	RunFileName  string
	OldZipFile   string
	Concurrent   int
	Command      string
	Status       int
	Notify       int
	NotifyEmail  string
	Timeout      int
	ExecuteTimes int
	PrevTime     int64
	CreateTime   int64
}

func (t *Task) Update(fields ...string) error {
	return nil
}

func TaskAdd(task *Task) (int64, error) {
	return 0, nil
}

func TaskGetList(page, pageSize int, filters ...interface{}) ([]*Task, int64) {
	return nil, 0
}

func TaskResetGroupId(groupId int) (int64, error) {
	return 0, nil
}

func TaskGetById(id int) (*Task, error) {

	return nil, nil
}

func TaskDel(id int) error {

	return nil
}
