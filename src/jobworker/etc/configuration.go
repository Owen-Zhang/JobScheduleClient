package etc

import "storage"
import "gopkg.in/yaml.v2"

import (
	"io/ioutil"
	"jobworker/api"
	"os"
	"jobworker/jobs"
	"jobworker/ctrl"
)

type Configuration struct {
	Storage struct {
		Hosts  string `yaml:"hosts,omitempty"`
		DBName string `yaml:"dbname,omitempty"`
		User   string   `yaml:"user,omitempty"`
		Password string  `yaml:"password,omitempty"` 
		Port int32  `yaml:"port,omitempty"`
	} `yaml:"storage,omitempty"`

	ApiServer struct {
		Bind string `yaml:"bind,omitempty"`
	} `yaml:"apiserver,omitempty"`

	CronService struct {
		PoolSize int32 `yaml:"poolsize,omitempty"`
	} `yaml:"cron,omitempty"`

	ExeConfig struct {
		ClientPath string `yaml:"clientpath,omitempty"`
		TempZipFolder string `yaml:"tempzipfolder,omitempty"`
		TaskFolder string `yaml:"taskfolder,omitempty"`
	} `yaml:"exeinfo,omitempty"`
}

var configuration *Configuration

func New(file string) error {
	fp, err := os.OpenFile(file, os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}

	c := makeDefault()
	if err := yaml.Unmarshal([]byte(data), c); err != nil {
		return err
	}

	configuration = c
	return nil
}

func makeDefault() *Configuration {

	return &Configuration{
		Storage: struct {
			Hosts  string `yaml:"hosts,omitempty"`
			DBName string `yaml:"dbname,omitempty"`
			User   string   `yaml:"user,omitempty"`
			Password string  `yaml:"password,omitempty"` 
			Port int32  `yaml:"port,omitempty"`
		}{
			Hosts:  "127.0.0.1",
			DBName: "jobschedule",
			User:   "guest",
			Password: "123456", 
			Port: 3306,
		},

		ApiServer: struct {
			Bind string `yaml:"bind,omitempty"`
		}{
			Bind: ":8985",
		},

		CronService: struct {
			PoolSize int32 `yaml:"poolsize,omitempty"`
		}{
			PoolSize: 10,
		},

		ExeConfig : struct {
			ClientPath string `yaml:"clientpath,omitempty"`
			TempZipFolder string `yaml:"tempzipfolder,omitempty"`
			TaskFolder string `yaml:"taskfolder,omitempty"`
		}{
			ClientPath: "D:\\code\\JobScheduleClient\\bin",
			TempZipFolder: "TempFile",
			TaskFolder: "Data",
		},
	}
}

//数据访问的相关配制
func GetStorageArg() *storage.DataStorageArgs {

	if configuration != nil {
		return &storage.DataStorageArgs{
			Hosts:  configuration.Storage.Hosts,
			DBName: configuration.Storage.DBName,
			User:   configuration.Storage.User,
			Password: configuration.Storage.Password,
			Port:  configuration.Storage.Port,
		}
	}
	return nil
}

//对外api的相关配制
func GetApiServerArg() *api.ApiServerArg {
	if configuration != nil {
		return &api.ApiServerArg{
			Bind: configuration.ApiServer.Bind,
		}
	}
	return nil
}

//cron的相关配制
func GetCronArg() *jobs.CronArg {
	if configuration != nil {
		return &jobs.CronArg{
			PoolSize: configuration.CronService.PoolSize,
		}
	}
	return nil
}

//client程序的配制
func GetExeConfig() * ctrl.ExeConfig {
	if configuration != nil {
		return &ctrl.ExeConfig{
			ClientPath : configuration.ExeConfig.ClientPath,
			TempZipFolder:configuration.ExeConfig.TempZipFolder,
			TaskFolder:configuration.ExeConfig.TaskFolder,
		}
	}
	return nil
}
