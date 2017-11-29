package model

type TaskNew struct {
	Id         int `json:"id"`         //任务的主键
	ZipFileUrl string `json:"zipfileurl"` //zip文件的下载地址
}

type Task struct {
	Id          int           //任务的主键
	Name        string         //任务名称
	CronSpec   string     	//cron表达式
	RunFilefolder   string    //任务的文件夹（代码放的文件夹名）
	OldZipFile    string      //原来的zip文件
	Command     string      //任务的命令如Init.exe xxx
	TimeOut     int32      //任务执行的超时时间
	Concurrent  int8   //是否允许在再一次没有运行完成的情况运行下一次
	Notify      int8    //是否需要通知
	NotifyEmail string //通知的邮件地址
	Version     int32  //程序的版本号 
}
