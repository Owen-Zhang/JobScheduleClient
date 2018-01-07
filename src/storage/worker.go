package storage

import (
	"model"
	"fmt"
	"errors"
)

//新增worker
func (this *DataStorage) AddWorker(info *model.HealthInfo) error {
	_, err := this.db.Exec(
		"INSERT into worker(name, url, port, systeminfo, status) VALUES(?,?,?,?,?)",
			info.Name, info.Url, info.Port, info.SystemInfo, info.Status)
	return err
}

// 根据名称查询单个worker
func (this *DataStorage) GetWorkerByName(name string) (*model.HealthInfo, error) {
	var nameT, url, systeminfo string
	var id, port, status int

	row := this.db.QueryRow("SELECT id, name, url, port, systeminfo, status from worker where name = ?;", name)
	if er := row.Scan(&id, &nameT, &url, &port, &systeminfo, &status); er != nil {
		return nil, er
	}

	return &model.HealthInfo{
		Id			: id,
		Name		: nameT,
		Url			: url,
		Port		: port,
		SystemInfo	: systeminfo,
		Status		: status,
	}, nil
}

//删除worker
func (this *DataStorage) DeleteWorker(id int) error  {
	stmt, _ := this.db.Prepare("update worker set status = 0 where id = ?")
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows <= 0 {
		return errors.New("delete faild, please try again soon")
	}
	return  nil
}

//查询出所有的worker机器(status = 2表示全部)
func (this *DataStorage) GetWorkerList(status int) ([]*model.HealthInfo, error) {
	rows, err := this.db.Query("SELECT id, name, url, port, systeminfo, status from worker where (? = 2 or ? = status);", status, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.HealthInfo
	for rows.Next() {
		var name, url, systeminfo string
		var id, port, status int
		if er := rows.Scan(&id, &name, &url, &port, &systeminfo, &status); er != nil {
			fmt.Printf("Query GetWorkerList has wrong : %s", er)
		}
		result = append(result, &model.HealthInfo{
			Id : id,
			Name: name,
			Url: url,
			Port: port,
			SystemInfo: systeminfo,
			Status: status,
		})
	}
	return result, nil
}
