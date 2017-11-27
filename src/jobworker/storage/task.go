package storage

import "fmt"

//根据id获取相关的任务信息
func (this *DataStorage) GetTaskById(id string) {
	err := this.db.QueryRow("11", 11)
	if err != nil {
		fmt.Println(err)
	}
}
