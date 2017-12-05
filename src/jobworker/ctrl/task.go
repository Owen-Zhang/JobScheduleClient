package ctrl

import (
	"jobworker/jobs"
	"model"
	"utils/system"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
	"net/http"
	"regexp"
	"io"
)

//运行任务(包括新增和重新启动)
func (this *Controller) start(request *Action) {
	//1: 查询数据，得到相关的实体数据
	task := this.Storage.GetTaskById(request.Id)
	if task == nil {
		return
	}

	command := task.Command

	if task.TaskType == 1 {
		//运行上传的文件
		taskfolder := strings.TrimSpace(task.RunFilefolder)
		datapath := fmt.Sprintf("%s\\%s\\%s\\", this.ExeConfig.ClientPath, this.ExeConfig.TaskFolder, taskfolder)
		if !system.FileExist(datapath) {
			//数据文件夹没有，需要创建相关的文件夹
			if err := os.MkdirAll(datapath, 0777); err != nil {
				fmt.Printf("create run fileFolder err : %s", err.Error())
				return
			}
		}

		configfile := fmt.Sprintf("%s\\config.txt", datapath)
		if !system.FileExist(configfile) {
			file, err := os.Create(configfile)
			if err != nil {
				fmt.Printf("create config file err: %s", err.Error())
				return
			}
			file.Close()
		}

		bytes, err := ioutil.ReadFile(configfile)
		if err != nil {
			fmt.Printf("read config file err: %s", err.Error())
			return
		}
		needpullfile := false
		if bytes != nil {
			contents := string(bytes)
			if len(contents) > 0 {
				version := strings.Split(contents, ":")[1]
				if err != nil {
					fmt.Printf("config file version err: %s", err.Error())
					return
				}
				if version != strconv.Itoa(task.Version) {
					needpullfile = true
				}
			} else {
				needpullfile = true
			}
		} else {
			needpullfile = true
		}

		//需要更新文件，同时更新配制
		if needpullfile {
			filename := system.UrlFileName(task.ZipFilePath)
			flag, _ := regexp.MatchString(`^.+\.zip$`, filename)
			if !flag {
				fmt.Println("zipfilepath has err, end must *.zip file")
				return
			}

			//下载文件
			res, err := http.Get(task.ZipFilePath)
			if err != nil {
				fmt.Printf("DownLoad File err: %s\n",err.Error())
				return
			}

			zipfile := fmt.Sprintf("%s\\%s\\%s", this.ExeConfig.ClientPath, this.ExeConfig.TempZipFolder, filename)
			file, err := os.Create(zipfile)
			if err != nil {
				fmt.Printf("save temp file err: %s\n", err.Error())
				return
			}

			if _,err := io.Copy(file, res.Body); err != nil {
				fmt.Printf("copy temp file err: %s\n", err.Error())
				return
			}
			defer func() {
				res.Body.Close()
			}()

			//解压到指定的文件夹中
			runfolder := fmt.Sprintf("%s\\Run\\", datapath)
			if err := system.UnzipFile(zipfile, runfolder); err != nil {
				fmt.Printf("unzipfile has wrong err: %s", err.Error())
				return
			}

			defer func() {
				file.Close()
				os.Remove(zipfile)
			}()

			//更新配制
			configcontent := []byte(fmt.Sprintf("version:%d", task.Version))
			err = ioutil.WriteFile(configfile, configcontent , os.ModeAppend)
			if err != nil {
				fmt.Printf("uupdate config file has wrong err: %s", err.Error())
				return
			}
		}
		command = fmt.Sprintf("%s\\Run\\%s", datapath, task.Command)
	}

	if !jobs.ExistJob(task.Id) {
		jobs.RemoveJob(task.Id)
	}

	jobs.AddJob(&model.Task{
		Id:       task.Id,
		Name:     task.Name,
		CronSpec: task.CronSpec, //"0 */1 * * * ?",
		Command:  command,
	})
}

//停止任务
func (this *Controller) stop(id int) {
	jobs.RemoveJob(id)
}

//删除任务
func (this *Controller) delete(id int) {
	if jobs.ExistJob(id) {
		jobs.RemoveJob(id)
	}
	if err := this.Storage.DeleteTask(id); err != nil {
		fmt.Printf("Delete Task has wrong: %s", err.Error())
	}
}
