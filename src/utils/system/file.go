package system

import (
	"io"
	"os"
	"strings"
	"path"
	"fmt"
	"github.com/satori/go.uuid"
)

func FileExist(filename string) bool {

	fi, err := os.Stat(filename)
	return (err == nil || os.IsExist(err)) && !fi.IsDir()
}

func DirExist(dirname string) bool {

	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

func FileCopy(source string, dest string) (int64, error) {

	sourcefile, err := os.Open(source)
	if err != nil {
		return 0, err
	}

	defer sourcefile.Close()
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return 0, err
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}

	w, err := io.Copy(destfile, sourcefile)
	if err != nil {
		destfile.Close()
		return 0, err
	}

	if err := os.Chmod(dest, sourceinfo.Mode()); err != nil {
		destfile.Close()
		return 0, err
	}
	destfile.Close()
	return w, nil
}

func DirectoryCopy(source string, dest string) error {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dest, sourceinfo.Mode()); err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()
	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			if err := DirectoryCopy(sourcefilepointer, destinationfilepointer); err != nil {
				return err
			}
		} else {
			if _, err := FileCopy(sourcefilepointer, destinationfilepointer); err != nil {
				return err
			}
		}
	}
	return nil
}


// 获取文件名带后缀
func FileNameWithExt(filepath string) string  {
	if filepath == "" {
		return ""
	}
	return path.Base(filepath)
}

// 获取文件后缀
func Ext(filepath string) string {
	fileName := FileNameWithExt(filepath)
	if fileName == "" {
		return ""
	}
	return path.Ext(fileName)
}

// 获取文件后缀
func FileName(path string) string {
	fileName := FileNameWithExt(path)
	if fileName == "" {
		return ""
	}
	return strings.TrimSuffix(fileName, Ext(path))
}

// 生成UUID的文件名
func CreateUuidFile(filepath string) string  {
	ext := Ext(filepath)
	if ext == "" {
		return ""
	}
	return fmt.Sprintf("%s%s",uuid.NewV4().String(), ext)
}

// 判断文件类型是否为想要类型
func  CheckFileExt(exts []string, filepath string) bool  {
	if len(exts) == 0 {
		return true
	}
	ext := strings.TrimLeft(Ext(filepath), ".")

	for _, value := range exts {
		if ext == value {
			return true
		}
	}
	return false
}

//判断文件是否存在
func IsExist(filepath string) bool {
	if (filepath == "") {
		return false
	}
	_, err := os.Stat(filepath)
	return err == nil || os.IsExist(err)
}

//找出url中的文件名，如http://www.baidu.com/aaa/12.zip?name=sdfasd 要取出12.zip文件名
func UrlFileName(url string) string  {
	array := strings.Split(url, "/")
	length := len(array)
	if length <= 0 {
		return ""
	}

	filename := array[length-1]
	arrry2 := strings.Split(filename, "?")

	return arrry2[0]
}