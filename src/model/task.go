package model

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

//任务
type Task struct {
	Id          	int           //任务的主键
	TaskType    	int           //任务类型 0命令型，1运行本地文件(上传文件),2调用外部接口
	Name        	string        //任务名称
	CronSpec   		string     	  //cron表达式
	RunFilefolder   string    	  //任务的文件夹（代码放的文件夹名）
	OldZipFile    	string        //原来的zip文件
	Command     	string        //任务的命令如Init.exe xxx
	TimeOut     	int32         //任务执行的超时时间
	Concurrent  	int8   		  //是否允许在再一次没有运行完成的情况运行下一次
	Notify      	int8          //是否需要通知
	NotifyEmail 	string 		  //通知的邮件地址
	Version     	int  		  //程序的版本号
	ZipFilePath     string 		  //zip的存储位置
}

//任务执行日志
type TaskLog struct {
	Id 			int				  //日志主键
	TaskId  	int				  //任务主键
	Output      string            //正常输出值
	Error		string			  //错误输出值
	ProcessTime int				  //执行时间
	CreateTime  int64			  //创建时间
	Status      int				  //日志状态
}
