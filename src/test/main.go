package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"model"
	"encoding/json"
	"bytes"
	//"github.com/shirou/gopsutil/mem"
	"github.com/hpcloud/tail"
	"log"
	"time"
)

func main()  {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	
	//loggerManager()
	
	watchFile()

	for ; ;  {
		
	}
}

func watchFile() {
	t, err := tail.TailFile("1.log", tail.Config{ReOpen: true,MustExist: false, Follow: true,Poll: true})
	if err != nil {
		fmt.Println(err)
	}
	for line := range t.Lines {
	    fmt.Println(line.Text)
	}
}

//记录日志(此方法不可取)
func loggerManager() {
	fileName := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0666)
	//file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		//return
	}	
	defer file.Close()
	logIn := log.New(file, "[dddd]", log.Llongfile)
	logIn.Println("test \n")
	logIn.Fatalln("fatal\n")
}

//上传文件测试
func uploadFile() {
	fileopen, err1 := os.Open("./jobworker.exe")
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	fd,err2 := ioutil.ReadAll(fileopen)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fileopen.Close()

	encodeString := base64.StdEncoding.EncodeToString(fd)

	info := model.Fileinfo{
		FilePath 		: "exefile",
		FileContent		:encodeString,
		FileSuffixName	: "exe",
	}

	content, errjson := json.Marshal(info)
	if errjson != nil {
		fmt.Println(errjson.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://127.0.0.1:8988/upload", bytes.NewReader(content))
	if err != nil {
		fmt.Println(errjson.Error())
		return
	}
	request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("Content-Type", "application/json")
	resp, errres  := http.DefaultClient.Do(request)
	if errres != nil {
		fmt.Println(errjson.Error())
		return
	}

	byteres, errrespose := ioutil.ReadAll(resp.Body)
	if errrespose != nil {
		fmt.Println(errrespose)
		return
	}
	fmt.Println(string(byteres))
	resp.Body.Close()
}

/*
//获取电脑相关的cpu 等
func getComputerInfo() {
	v, _ := mem.VirtualMemory()
	fmt.Println(v)
}
*/
