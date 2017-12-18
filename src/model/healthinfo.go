package model

//心跳检查的客户端信息
type HealthInfo struct {
	Id  		int
	Name 		string
	SystemInfo  string
	WorkerUrl 	string
	WorkerPort  int
}
