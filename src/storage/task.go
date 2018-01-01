package storage

import (
	"fmt"
	"model"
)

//根据id获取相关的任务信息
func (this *DataStorage) GetTaskById(idinput int) (*model.TaskExend, error) {
	sqltext := `SELECT id, user_id, group_id, task_name, task_type, description, cron_spec, run_file_folder,
			old_zip_file, concurrent, command, status, notify, notify_email, timeout, execute_times,
			prev_time, create_time, version, zip_file_path from task where deleted = 0 and id=?;`

	row := this.db.QueryRow(sqltext, idinput)

	var task_name,description, cron_spec, run_file_folder, old_zip_file, command, notify_email, zip_file_path string
	var id, user_id, group_id,task_type, concurrent, status, notify, timeout, execute_times, version int
	var create_time, prev_time int64

	if er := row.Scan(&id, &user_id, &group_id, &task_name, &task_type, &description, &cron_spec, &run_file_folder,
		&old_zip_file, &concurrent, &command, &status, &notify, &notify_email, &timeout, &execute_times,
		&prev_time, &create_time, &version, &zip_file_path); er != nil {

		return nil, er
		fmt.Printf("GetTaskById has wrong : %s", er)
	}

	return &model.TaskExend {
		Task:model.Task{
			Id 				: id,
			TaskType		: task_type,
			Name			: task_name,
			CronSpec		: cron_spec,
			RunFilefolder	: run_file_folder,
			OldZipFile		: old_zip_file,
			Command			: command,
			Notify			: notify,
			NotifyEmail 	: notify_email,
			Concurrent  	: concurrent,
			TimeOut			: timeout,
			Version			: version,
			ZipFilePath		: zip_file_path,
		},
		UserId		: user_id,
		GroupId		: group_id,
		Description	: description,
		Status		: status,
		ExecuteTimes: execute_times,
		PrevTime	: prev_time,
		CreateTime	: create_time,
	}, nil
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

func (this *DataStorage) UpdateFrontTask(task *model.TaskExend) error {
	if _, err := this.db.Exec(
		`update task set group_id = ?, task_name = ?, task_type = ?, description = ?, cron_spec = ?,
				old_zip_file = ?, concurrent = ?, command = ?, notify = ?, notify_email = ?, timeout = ?,
				version = ?, zip_file_path = ? where id = ?`,
		task.GroupId, task.Name, task.TaskType, task.Description, task.CronSpec,
		task.OldZipFile, task.Concurrent, task.Command, task.Notify, task.NotifyEmail, task.TimeOut,
		task.Version, task.ZipFilePath, task.Id); err != nil {
			return err
	}
	return nil
}

func (this *DataStorage) TaskAdd(task *model.TaskExend) (error) {
	if _, err := this.db.Exec(
		`INSERT into task(user_id, group_id, task_name, task_type, description, cron_spec, run_file_folder,
								old_zip_file, concurrent, command, status, notify, notify_email, timeout, execute_times,
							    prev_time, create_time, version, delete, zip_file_path)
				VALUES(?,?,?,?,?,?,?,  ?,?,?,0,?,?,?,0,  ?,?,?,1,?)`,
		task.UserId, task.GroupId, task.Name, task.TaskType, task.Description, task.CronSpec, task.RunFilefolder,
		task.OldZipFile, task.Concurrent, task.Command, task.Notify, task.NotifyEmail, task.TimeOut,
		task.PrevTime,task.CreateTime,task.Version, task.ZipFilePath); err != nil {
			return err
	}

	return nil
}

//status -1为全部，其它为数据库正常状态; groupid: 0表示全部，其它表示正常分组下的
func (this *DataStorage) TaskGetList(page, pageSize, status, groupid int) ([]*model.TaskExend, int) {

	total := this.taskGetListCount(status, groupid)
	if total <= 0 {
		return nil, 0
	}

	rows, err := this.db.Query(
`SELECT
			id, user_id, group_id, task_name, task_type, description, cron_spec, run_file_folder,
			old_zip_file, concurrent, command, status, notify, notify_email, timeout, execute_times,
			prev_time, create_time, version, zip_file_path
		from task
		where (? = -1 or ? = status) AND
			  (? = 0 or ? = group_id) AND
              deleted = 0
		order by id ASC
		LIMIT ?, ?;`, status, status, groupid, groupid, (page - 1)*pageSize, pageSize)

	if err != nil {
		fmt.Printf("TaskGetList has wrong: %s\n", err)
		return nil, 0
	}
	defer rows.Close()

	var result []*model.TaskExend
	for rows.Next() {

		var task_name,description, cron_spec, run_file_folder, old_zip_file, command, notify_email, zip_file_path string
		var id, user_id, group_id,task_type, concurrent, status, notify, timeout, execute_times, version int
		var create_time, prev_time int64

		if er := rows.Scan(&id, &user_id, &group_id, &task_name, &task_type, &description, &cron_spec, &run_file_folder,
			&old_zip_file, &concurrent, &command, &status, &notify, &notify_email, &timeout, &execute_times,
			&prev_time, &create_time, &version, &zip_file_path); er != nil {

			return nil, 0
			fmt.Printf("Query TaskGetList has wrong : %s", er)
		}
		result = append(result, &model.TaskExend{
			Task:model.Task {
				Id 			 : id,
				TaskType	 : task_type,
				Name		 : task_name,
				CronSpec	 : cron_spec,
				RunFilefolder: run_file_folder,
				OldZipFile	 : old_zip_file,
				Command		 : command,
				Notify		 : notify,
				NotifyEmail  : notify_email,
				Concurrent   : concurrent,
				TimeOut		 : timeout,
				Version		 : version,
				ZipFilePath	 : zip_file_path,
			},
			UserId		: user_id,
			GroupId		: group_id,
			Description	: description,
			Status		: status,
			ExecuteTimes: execute_times,
			PrevTime	: prev_time,
			CreateTime	: create_time,
		})
	}

	return result, total
}

func (this *DataStorage) taskGetListCount(status, groupid int) int {

	var total int
	this.db.QueryRow(
		`SELECT
					count(1) as total
				from task
				where (? = -1 or ? = status) AND
					  (? = 0 or ? = group_id) AND
				      deleted = 0;`,
		status,status,groupid,groupid).Scan(&total)
	return total
}

//更新任务状态
func (this *DataStorage) TaskUpdateStatus(id, status int) error {
	_, err := this.db.Exec("update task set status = ? where id=?;",  status, id)
	return err
}

//删除任务
func (this *DataStorage) TaskDel(id int) error {
	_, err := this.db.Exec("update task set delete = 0 where id=?;", id)
	return err
}
