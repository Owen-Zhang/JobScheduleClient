package main

import (
	"jobworker/server"
	"os"
	"utils/system"
	"fmt"
)

// http://127.0.0.1:8084/task/Console.zip

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
