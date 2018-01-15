package main

import (
	"jobworker/server"
	"os"
	"utils/system"
	"fmt"
)


// updateFileInfo这个要处理，文件下载地址不对

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	work, err := server.NewWorker()
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	defer func() {
		work.Stop()
		os.Exit(0)
	}()

	if err := work.Start(); err != nil {
		panic(err)
		os.Exit(-2)
	}

	system.InitSignal(nil)
}
