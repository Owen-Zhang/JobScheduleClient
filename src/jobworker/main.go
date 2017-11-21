package main

import (
	"jobworker/server"
	"os"
	"utils/system"
)

func main() {
	/*
	defer func() {
		recover()
	}()*/

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
