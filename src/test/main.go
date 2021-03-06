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
	//"path/filepath"
	//"time"
	//"log"
	//"os/exec"
	"utils/log"
	"utils/system"
	"path"
	"os/exec"
	"github.com/axgle/mahonia"
	"strings"
)

func main()  {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	testcmdFile()

	//testunzipFile()

	//testFile()

	//testCreateFolder()

	//testLog()
	//loggerManager()
	
	//watchFile()

	/*
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
	*/

	for ; ;  {
		
	}
}

//相对目录的问题，exec.cmd() 用相对目录不能正常运行
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}

func testcmdFile()  {

	fmt.Println(getCurrentPath())

	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	//cmd := exec.Command("cmd.exe", "/c", `D:\Code\JobScheduleClient\bin\Data\9d1a4f1e-3878-46ec-8308-eb4571209a9d\Run\Test.exe`)
	cmd := exec.Command("cmd.exe", "/c", getCurrentPath() + "Data/9d1a4f1e-3878-46ec-8308-eb4571209a9d/Run/Test.exe")
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	cmd.Start()

	cmd.Wait()

	encoder := mahonia.NewDecoder("gbk")
	fmt.Println(encoder.ConvertString(bufOut.String()))
	fmt.Println(encoder.ConvertString(bufErr.String()))
}

func testunzipFile()  {
	err := system.UnzipFile("TempFile/71d34eff-530f-4e19-8024-09f29ce77f14.zip", "test/t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unzip file")
}

func testFile()  {
	fileName := path.Base(`job/e15f0794-a87b-4836-8227-0676aeea38bd.zip`)
	fmt.Println(fileName)
}

func testCreateFolder()  {

	if !system.FileExist("147258369/789456123") {
		//数据文件夹没有，需要创建相关的文件夹
		if err := os.MkdirAll("147258369/789456123", 0777); err != nil {
			fmt.Printf("create run fileFolder err : %s", err.Error())
			return
		}
	} else {
		fmt.Println("文件夹存在")
	}
}

func testLog() {
	log.Errorf("test log..\nsdfasdfasdfasd")
	log.Errorf("111111")
}

/*
func watchFile() {
	t, err := tail.TailFile("1.log", tail.Config{ReOpen: true,MustExist: false, Follow: true,Poll: true})
	if err != nil {
		fmt.Println(err)
	}
	for line := range t.Lines {
	    fmt.Println(line.Text)
	}
}
*/

/*
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
*/

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
