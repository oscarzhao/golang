package main

import (
	"log"
	"os"
	"time"
)

var (
	Logger *log.Logger
)

func setLogger(logFile string) {
	if logFile == "stdout" {
		Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	} else {
		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
			log.Println("Log to stdout instead")
			f = os.Stdout
		}
		Logger = log.New(f, "", log.Ldate|log.Ltime)
	}
}

func startSubProcess() {
	elapsedSec := 0
	for {
		time.Sleep(time.Second)
		elapsedSec += 1
		Logger.Printf("subprocess: %d seconds passed", elapsedSec)
	}
	Logger.Printf("Sub Process exit")
}
func main() {
	setLogger("./goruote.log")
	go startSubProcess()
	elapsedSec := 0
	for {
		time.Sleep(2 * time.Second)
		elapsedSec += 2
		Logger.Printf("main process: %d seconds passed", 2*elapsedSec)
	}
	Logger.Printf("Main process exit")

}
