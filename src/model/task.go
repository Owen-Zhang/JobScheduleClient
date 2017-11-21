package model

type Task struct {
	Id 			string  `json:"id"`       	//任务的主键
	IsNew 		bool	`json:"isnew"`      //是否为新增任务
	ZipFileUrl 	string  `json:"zipfileurl"` //zip文件的下载地址
	ExeFolder   string  `json:"exefolder"`  //程序的安装目录名(要执行程序的目录名)
}
