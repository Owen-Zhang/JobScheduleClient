package storage

import (
	"fmt"
	"model"
)

//根据id获取相关的任务信息
func (this *DataStorage) GetTaskById(idinput int) *model.Task {
	sqltext := "SELECT id, task_type, task_name, cron_spec, run_file_folder, old_zip_file, concurrent, command, notify, notify_email, timeout, version, zip_file_path from task where STATUS = 1 and `delete` = 0 and id=?;"
	row := this.db.QueryRow(sqltext, idinput)

	var id, task_type,version, notify, concurrent, timeout int
	var task_name, cron_spec, run_file_folder, old_zip_file, command, notify_email, zip_file_path string
	err := row.Scan(&id, &task_type, &task_name, &cron_spec, &run_file_folder, &old_zip_file, &concurrent, &command, &notify, &notify_email, &timeout, &version, &zip_file_path)

	if err != nil {
		fmt.Println(err)
		return  nil
	}
	return &model.Task {
		Id : id,
		TaskType:task_type,
		Name: task_name,
		CronSpec:cron_spec,
		RunFilefolder: run_file_folder,
		OldZipFile: old_zip_file,
		Command:command,
		Notify:notify,
		NotifyEmail:notify_email,
		Concurrent : concurrent,
		TimeOut: timeout,
		Version: version,
		ZipFilePath:zip_file_path,
	}
}

//删除任务
func (this *DataStorage) DeleteTask(id int)  error {
	_, err := this.db.Exec("update task set `delete` = 1 where id=?;",id)
	return err
}

//更新任务的相关信息
func (this *DataStorage) UpdateBackTask (prevtime int64, id int) error {
	_, err := this.db.Exec("update task set prev_time = ?, execute_times = execute_times + 1 where id=?;", prevtime, id)
	return err
}

//增加日志信息
func (this *DataStorage) AddTaskLog(log *model.TaskLog) error {
	_, err := this.db.Exec(
			"INSERT into task_log(task_id, output, error, `status`, process_time, create_time) VALUES(?,?,?,?,?,?)",
			log.TaskId, log.Output, log.Error, log.Status, log.ProcessTime, log.CreateTime)
	return err
}

func (this *DataStorage) UpdateFrontTask(fields ...string) error {
	return nil
}

func (this *DataStorage) TaskAdd(task *model.TaskExend) (int64, error) {
	return 0, nil
}

func (this *DataStorage) TaskGetList(page, pageSize int, filters ...interface{}) ([]*model.TaskExend, int64) {
	return nil, 0
}

func (this *DataStorage) TaskResetGroupId(groupId int) (int64, error) {
	return 0, nil
}

func (this *DataStorage) TaskGetById(id int) (*model.TaskExend, error) {

	return nil, nil
}

func (this *DataStorage) TaskDel(id int) error {

	return nil
}
