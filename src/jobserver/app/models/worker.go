package models

import "model"

//新增worker
func AddWorker(info *model.HealthInfo) bool {

	return true
}

//更新worker
func UpdateWorker(info *model.HealthInfo) bool {

	return true
}

//删除worker
func DeleteWorker(id int) bool  {

	//需要将当前worker中的任务全部结束
	return true
}

//查询出所有的worker机器
func GetWorkerList() ([]*model.HealthInfo, error) {

	return nil, nil
}
