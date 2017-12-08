package main


import (
	"net/http"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"os"
	"fmt"
)

type Fileinfo struct {
	FilePath 		string `json:"filepath"`       //文件路径，如：order/detail
	FileSuffixName  string `json:"filesuffixname"` //文件的后缀名 如: exe, jpg
	FileContent  	string `json:"filecontent"`    //文件的byte的64编码
}

func main() {
	/*
	fileopen, err1 := os.Open("./TEST.rar")
	if (err1 != nil) {
		fmt.Println(err1.Error())
		return
	}
	fd,err2 := ioutil.ReadAll(fileopen)
	if (err2 != nil) {
		fmt.Println(err2.Error())
		return
	}
	encodeString := base64.StdEncoding.EncodeToString(fd)
	
	file, err := os.Create("./staticfile/TEST.rar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	filecontent, err3 := base64.StdEncoding.DecodeString(encodeString)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
	
	file.Write(filecontent)
	file.Close()
	
	
	*/
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("./staticfile"))))
	fmt.Println("文件下载")
	
	http.ListenAndServe(":8988", nil)
}

func uploadHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			by, err1 := ioutil.ReadAll(r.Body)
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			//b := new(bytes.Buffer)
			//b.Write(by)

			body := &Fileinfo{}
			err := json.Unmarshal(by, body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//保存文件 此处还要创建文件夹
			file, err := os.Create(fmt.Sprintf("%s/%s/%s.%s", "./staticfile", body.FilePath, "XXX", body.FileSuffixName))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			filecontent, err3 := base64.StdEncoding.DecodeString(body.FileContent)
			if err3 != nil {
				fmt.Println(err3.Error())
				return
			}

			file.Write(filecontent)
			file.Close()

		default:
        	w.WriteHeader(http.StatusMethodNotAllowed)
	}
}