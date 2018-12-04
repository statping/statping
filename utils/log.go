// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	logFile   *os.File
	fmtLogs   *log.Logger
	ljLogger  *lumberjack.Logger
	LastLines []*LogRow
	LockLines sync.Mutex
)

// createLog will create the '/logs' directory based on a directory
func createLog(dir string) error {
	var err error
	_, err = os.Stat(dir + "/logs")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir+"/logs", 0777)
		} else {
			return err
		}
	}
	file, err := os.Create(dir + "/logs/statup.log")
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

// InitLogs will create the '/logs' directory and creates a file '/logs/statup.log' for application logging
func InitLogs() error {
	err := createLog(Directory)
	if err != nil {
		return err
	}
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
	LastLines = make([]*LogRow, 0)
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

// Log creates a new entry in the Logger. Log has 1-5 levels depending on how critical the log/error is
func Log(level int, err interface{}) error {
	pushLastLine(err)
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

// Http returns a log for a HTTP request
func Http(r *http.Request) string {
	msg := fmt.Sprintf("%v (%v) | IP: %v", r.RequestURI, r.Method, r.Host)
	fmt.Printf("WEB: %v\n", msg)
	pushLastLine(msg)
	return msg
}

func pushLastLine(line interface{}) {
	LockLines.Lock()
	defer LockLines.Unlock()
	LastLines = append(LastLines, newLogRow(line))
	// We want to store max 1000 lines in memory (for /logs page).
	for len(LastLines) > 1000 {
		LastLines = LastLines[1:]
	}
}

// GetLastLine returns 1 line for a recent log entry
func GetLastLine() *LogRow {
	LockLines.Lock()
	defer LockLines.Unlock()
	if len(LastLines) > 0 {
		return LastLines[len(LastLines)-1]
	}
	return nil
}

type LogRow struct {
	Date time.Time
	Line interface{}
}

func newLogRow(line interface{}) (logRow *LogRow) {
	logRow = new(LogRow)
	logRow.Date = time.Now()
	logRow.Line = line
	return
}

func (o *LogRow) lineAsString() string {
	switch v := o.Line.(type) {
	case string:
		return v
	case error:
		return v.Error()
	case []byte:
		return string(v)
	}
	return ""
}

func (o *LogRow) FormatForHtml() string {
	return fmt.Sprintf("%s: %s", o.Date.Format("2006-01-02 15:04:05"), o.lineAsString())
}
