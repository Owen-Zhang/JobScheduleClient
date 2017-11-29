package etc

import "jobworker/storage"
import "gopkg.in/yaml.v2"

import (
	"io/ioutil"
	"jobworker/api"
	"os"
	"jobworker/jobs"
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
	}
}

func GetConfiguration() *Configuration {
	return configuration
}

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

func GetApiServerArg() *api.ApiServerArg {
	if configuration != nil {
		return &api.ApiServerArg{
			Bind: configuration.ApiServer.Bind,
		}
	}
	return nil
}

func GetCronArg() *jobs.CronArg {
	if configuration != nil {
		return &jobs.CronArg{
			PoolSize: configuration.CronService.PoolSize,
		}
	}
	return nil
}
