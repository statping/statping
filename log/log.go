package log

import (
	"github.com/fatih/color"
	lg "log"
	"os"
	//"github.com/mkideal/log/logger"
)

var (
	logFile  *os.File
	logLevel int
)

func init() {
	var err error
	logFile, err = os.OpenFile("statup.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		lg.Fatalf("error opening file: %v", err)
	}
	lg.SetOutput(logFile)

	logEnv := os.Getenv("LOG")
	if logEnv == "fatal" {
		logLevel = 3
	} else if logEnv == "debug" {
		logLevel = 2
	} else if logEnv == "info" {
		logLevel = 1
	} else {
		logLevel = 0
	}

}

func Panic(err interface{}) {
	panic(err)
}

func Send(level int, err interface{}) {
	switch level {
	case 3:
		lg.Printf("ERROR: %v\n", err)
		color.Red("ERROR: %v\n", err)
		os.Exit(2)
	case 2:
		lg.Printf("WARNING: %v\n", err)
		color.Yellow("WARNING: %v\n", err)
	case 1:
		lg.Printf("INFO: %v\n", err)
		color.Blue("INFO: %v\n", err)
	case 0:
		lg.Printf("%v\n", err)
		color.White("%v\n", err)
	}
}
