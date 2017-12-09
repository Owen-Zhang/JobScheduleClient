package main


import (
	"net/http"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"os"
	"fmt"
	"utils/system"
	"regexp"
	"model"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("./staticfile"))))
	fmt.Println("file server start")
	
	http.ListenAndServe(":8988", nil)
}

func uploadHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			result := &model.FileResponse{
				Status : false,
				Message: "Save file faild, please try again soon",
			}

			by, err1 := ioutil.ReadAll(r.Body)
			if err1 != nil {
				result.Message = err1.Error()
				send(result, w)
				return
			}

			body := &model.Fileinfo{}
			err := json.Unmarshal(by, body)
			if err != nil {
				result.Message = err.Error()
				send(result, w)
				return
			}

			//判断文件后缀
			flag, errregex := regexp.MatchString(`^[a-z]+$`, body.FileSuffixName)
			if errregex != nil || !flag {
				result.Message = "FileSuffixName is not valid: [a-z]+"
				send(result, w)
				return
			}

			//创建文件夹
			folder := fmt.Sprintf("%s/%s", "./staticfile", body.FilePath)
			if !system.DirExist(folder) {
				os.MkdirAll(folder, 0777)
			}
			filename := fmt.Sprintf("%s.%s", system.GetUuid(), body.FileSuffixName)

			//生成文件名
			file, err := os.Create(fmt.Sprintf("%s/%s", folder, filename))
			if err != nil {
				result.Message = err.Error()
				send(result, w)
				return
			}
			defer file.Close()

			filecontent, err3 := base64.StdEncoding.DecodeString(body.FileContent)
			if err3 != nil {
				result.Message = err3.Error()
				send(result, w)
				return
			}
			file.Write(filecontent)

		default:
        	w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func send(res *model.FileResponse, w http.ResponseWriter)  {
	jsonResult, errjson := json.Marshal(res)
	if errjson != nil {
		fmt.Println(errjson)
		return
	}
	w.Write(jsonResult)
}