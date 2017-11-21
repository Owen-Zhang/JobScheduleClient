package etc

//import "api"
import "jobworker/storage"
import "gopkg.in/yaml.v2"

import (
	"errors"
	"io/ioutil"
	"jobworker/api"
	"os"
	"utils/system"
)

type Auth map[string]string

type Configuration struct {
	Storage struct {
		Hosts  string `yaml:"hosts,omitempty"`
		DBName string `yaml:"dbname,omitempty"`
		Auth   Auth   `yaml:"auth,omitempty"`
	} `yaml:"storage,omitempty"`

	ApiServer struct {
		Bind string `yaml:"bind,omitempty"`
	} `yaml:"apiserver,omitempty"`

	//api 和 日志
	/*
		Logger struct {
			LogFile  string `yaml:"logfile,omitempty"`
			LogLevel string `yaml:"loglevel,omitempty"`
			LogSize  int64  `yaml:"logsize,omitempty"`
		} `yaml:"logger,omitempty"`
	*/
}

var configuration *Configuration

func New(file string) error {
	if !system.FileExist(file) {
		return errors.New("etc -> worker.yml invalid")
	}

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
			Auth   Auth   `yaml:"auth,omitempty"`
		}{
			Hosts:  "127.0.0.1:4705",
			DBName: "jobschedule",
			Auth:   map[string]string{},
		},

		ApiServer: struct {
			Bind string `yaml:"bind,omitempty"`
		}{
			Bind: ":8910",
		},

		//日志
		/*
			Logger: struct {
				LogFile  string `yaml:"logfile,omitempty"`
				LogLevel string `yaml:"loglevel,omitempty"`
				LogSize  int64  `yaml:"logsize,omitempty"`
			}{
				LogFile:  "logs/jobworker.log",
				LogLevel: "debug",
				LogSize:  2097152,
			},
		*/
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
			Auth:   configuration.Storage.Auth,
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

//日志
/*
func GetLogger() *logger.Args {

	if configuration != nil {
		return &logger.Args{
			FileName: configuration.Logger.LogFile,
			Level:    configuration.Logger.LogLevel,
			MaxSize:  configuration.Logger.LogSize,
		}
	}
	return nil
}
*/
