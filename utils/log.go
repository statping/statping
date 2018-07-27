package utils

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	logFile  *os.File
	logLevel int
	fmtLogs  *log.Logger
	ljLogger *lumberjack.Logger
	LastLine interface{}
)

func InitLogs() error {
	var err error

	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0777)
	}

	logFile, err = os.OpenFile("./logs/statup.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Printf("ERROR opening file: %v", err)
		return err
	}
	ljLogger = &lumberjack.Logger{
		Filename:   "./logs/statup.log",
		MaxSize:    16,
		MaxBackups: 3,
		MaxAge:     28,
	}
	fmtLogs = log.New(logFile, "", log.Ldate|log.Ltime)

	log.SetOutput(ljLogger)

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

	return err
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

func Log(level int, err interface{}) {
	LastLine = err
	switch level {
	case 5:
		fmt.Printf("PANIC: %v\n", err)
		fmtLogs.Fatalf("PANIC: %v\n", err)
	case 4:
		fmt.Printf("FATAL: %v\n", err)
		fmtLogs.Printf("FATAL: %v\n", err)
		//color.Red("ERROR: %v\n", err)
		//os.Exit(2)
	case 3:
		fmt.Printf("ERROR: %v\n", err)
		fmtLogs.Printf("ERROR: %v\n", err)
		//color.Red("ERROR: %v\n", err)
	case 2:
		fmt.Printf("WARNING: %v\n", err)
		fmtLogs.Printf("WARNING: %v\n", err)
		//color.Yellow("WARNING: %v\n", err)
	case 1:
		fmt.Printf("INFO: %v\n", err)
		fmtLogs.Printf("INFO: %v\n", err)
		//color.Blue("INFO: %v\n", err)
	case 0:
		fmt.Printf("%v\n", err)
		fmtLogs.Printf("%v\n", err)
		color.White("%v\n", err)
	}
}

func Http(r *http.Request) {
	msg := fmt.Sprintf("%v (%v) | IP: %v", r.RequestURI, r.Method, r.Host)
	fmtLogs.Printf("WEB: %v\n", msg)
	fmt.Printf("WEB: %v\n", msg)
	LastLine = msg
}
