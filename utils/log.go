package utils

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/natefinch/lumberjack.v2"
	lg "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	logFile  *os.File
	logLevel int
	fmtLogs  *lg.Logger
	ljLogger *lumberjack.Logger
)

func init() {
	var err error
	logFile, err = os.OpenFile("./statup.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		lg.Printf("ERROR opening file: %v", err)
	}
	ljLogger = &lumberjack.Logger{
		Filename:   "./statup.log",
		MaxSize:    16,
		MaxBackups: 3,
		MaxAge:     28,
	}
	fmtLogs = lg.New(logFile, "", lg.Ldate|lg.Ltime)
	fmtLogs.SetOutput(ljLogger)

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

	rotate()
}

func rotate() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-c
			ljLogger.Rotate()
		}
	}()
}

func Panic(err interface{}) {
	lg.Printf("PANIC: %v\n", err)
	panic(err)
}

func Log(level int, err interface{}) {
	switch level {
	case 5:
		lg.Fatalf("PANIC: %v\n", err)
	case 4:
		lg.Printf("FATAL: %v\n", err)
		//color.Red("ERROR: %v\n", err)
		//os.Exit(2)
	case 3:
		lg.Printf("ERROR: %v\n", err)
		//color.Red("ERROR: %v\n", err)
	case 2:
		lg.Printf("WARNING: %v\n", err)
		//color.Yellow("WARNING: %v\n", err)
	case 1:
		lg.Printf("INFO: %v\n", err)
		//color.Blue("INFO: %v\n", err)
	case 0:
		lg.Printf("%v\n", err)
		color.White("%v\n", err)
	}
}

func Http(r *http.Request) {
	msg := fmt.Sprintf("%v (%v) | IP: %v", r.RequestURI, r.Method, r.Host)
	lg.Printf("WEB: %v\n", msg)
}

func ReportLog() {

}
