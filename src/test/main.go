package main

import (
	/*
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"model"
	"encoding/json"
	"bytes"
	"path/filepath"
	*/
	"os"
	"fmt"
	"os/exec"
	"path/filepath"
)

func main()  {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	os.Mkdir("Temp", 0777)

	file1, err1 := os.Create("Temp/1.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	file1.Close()

	fmt.Println(os.Args[0])

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Dir(file)
	fmt.Println(path)

	/*

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
	*/

	for ; ;  {
		
	}
}
