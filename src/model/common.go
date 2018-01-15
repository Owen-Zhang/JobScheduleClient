package model

type WorkerFileConfig struct {
	Version  int    `json:"version"`
	FileName string `json:"filename"`
}

//文件服务器相关配制信息
type FileServerInfo struct {
	Hosts  string 
	Port   int    
}