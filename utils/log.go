package utils

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const MAX_LAST_LINES = 200

var (
	logFile  *os.File
	fmtLogs  *log.Logger
	ljLogger *lumberjack.Logger
	LastLine interface{}
	LastLines []interface{}	// Could be some cache queue in future.
)

func InitLogs() error {
	var err error

	if _, err := os.Stat(Directory + "/logs"); os.IsNotExist(err) {
		os.Mkdir(Directory+"/logs", 0777)
	}

	file, err := os.Create(Directory + "/logs/statup.log")
	if err != nil {
		return err
	}
	defer file.Close()

	logFile, err = os.OpenFile(Directory+"/logs/statup.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Printf("ERROR opening file: %v", err)
		return err
	}
	ljLogger = &lumberjack.Logger{
		Filename:   Directory + "/logs/statup.log",
		MaxSize:    16,
		MaxBackups: 3,
		MaxAge:     28,
	}
	fmtLogs = log.New(logFile, "", log.Ldate|log.Ltime)
	log.SetOutput(ljLogger)
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

func Log(level int, err interface{}) error {
	setLasLine(err)

	var outErr error
	switch level {
	case 5:
		_, outErr = fmt.Printf("PANIC: %v\n", err)
		fmtLogs.Printf("PANIC: %v\n", err)
	case 4:
		_, outErr = fmt.Printf("FATAL: %v\n", err)
		fmtLogs.Printf("FATAL: %v\n", err)
		//color.Red("ERROR: %v\n", err)
		//os.Exit(2)
	case 3:
		_, outErr = fmt.Printf("ERROR: %v\n", err)
		fmtLogs.Printf("ERROR: %v\n", err)
		//color.Red("ERROR: %v\n", err)
	case 2:
		_, outErr = fmt.Printf("WARNING: %v\n", err)
		fmtLogs.Printf("WARNING: %v\n", err)
		//color.Yellow("WARNING: %v\n", err)
	case 1:
		_, outErr = fmt.Printf("INFO: %v\n", err)
		fmtLogs.Printf("INFO: %v\n", err)
		//color.Blue("INFO: %v\n", err)
	case 0:
		_, outErr = fmt.Printf("%v\n", err)
		fmtLogs.Printf("%v\n", err)
		//color.White("%v\n", err)
	}
	return outErr
}

func Http(r *http.Request) string {
	msg := fmt.Sprintf("%v (%v) | IP: %v", r.RequestURI, r.Method, r.Host)
	fmtLogs.Printf("WEB: %v\n", msg)
	fmt.Printf("WEB: %v\n", msg)
	setLasLine(msg)
	return msg
}

func setLasLine(line interface{}) {
	LastLine = line
	LastLines = append(LastLines, LastLine)
	if len(LastLines) > MAX_LAST_LINES {
		LastLines = LastLines[1:]
	}
}
