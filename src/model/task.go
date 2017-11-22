package model

type TaskNew struct {
	Id         string `json:"id"`         //任务的主键
	ZipFileUrl string `json:"zipfileurl"` //zip文件的下载地址
}

//任务(现不包括发邮件etc..)
type Task struct {
	Id          string `json:"id"`          //任务的主键
	Name        string `json:"name"`        //任务名称
	Description string `json:"description"` //任务描述
	CronSpec   string `json:"cronspec"`   //cron表达式
	ExeFolder   string `json:"exefolder"`   //任务的文件夹（代码放的文件夹名）
	Command     string `json:"command"`     //任务的命令如Init.exe xxx
	TimeOut     int32  `json:"timeout"`     //任务执行的超时时间
}
