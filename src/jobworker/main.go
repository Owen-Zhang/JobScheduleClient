package main

import (
	"jobworker/server"
	"os"
)

func main() {
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

	//可以处理一些退出的回调
}
