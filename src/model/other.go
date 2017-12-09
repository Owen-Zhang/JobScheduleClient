package model

type Fileinfo struct {
	FilePath 		string `json:"filepath"`       //文件路径，如：order/detail
	FileSuffixName  string `json:"filesuffixname"` //文件的后缀名 如: exe, jpg
	FileContent  	string `json:"filecontent"`    //文件的byte的64编码
}

type FileResponse struct {
	Status      		bool     `json:"status"` 		  //文件路径，如：order/detail
	Message     		string   `json:"message"` 		  //文件路径，如：order/detail
	PathName     		string   `json:"pathname"` 		  //文件路径，如：order/detail
}
