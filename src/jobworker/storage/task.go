package storage

import (
	"fmt"
	"model"
)

//根据id获取相关的任务信息
func (this *DataStorage) GetTaskById(idinput int) *model.Task {
	sqltext := "SELECT id, task_name, cron_spec, run_file_folder, old_zip_file, concurrent, command, notify, notify_email, timeout, version from task where STATUS = 1 and id=?;"
	row := this.db.QueryRow(sqltext, idinput)

	var id int
	var task_name, cron_spec, run_file_folder, old_zip_file, command, notify_email string
	var timeout, version int32
	var notify, concurrent int8
	err := row.Scan(&id, &task_name, &cron_spec, &run_file_folder, &old_zip_file, &concurrent, &command, &notify, &notify_email, &timeout, &version)

	if err != nil {
		fmt.Println(err)
		return  nil
	}
	return &model.Task {
		Id : id,
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
	}
}

//更新任务的相关信息
func (this *DataStorage) UpdateTask () {
	
}

//增加日志信息
func (this *DataStorage) AddTaskLog() {
	
}

