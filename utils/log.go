// Statping
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
	"github.com/fatih/structs"
	"github.com/hunterlong/statping/types"
	Logger "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	Log         = Logger.StandardLogger()
	ljLogger    *lumberjack.Logger
	LastLines   []*logRow
	LockLines   sync.Mutex
	VerboseMode int
)

const logFilePath = "/logs/statping.log"

type hook struct {
	Entries []Logger.Entry
	mu      sync.RWMutex
}

func (t *hook) Fire(e *Logger.Entry) error {
	pushLastLine(e.Message)
	return nil
}

func (t *hook) Levels() []Logger.Level {
	return Logger.AllLevels
}

// ToFields accepts any amount of interfaces to create a new mapping for log.Fields. You will need to
// turn on verbose mode by starting Statping with "-v". This function will convert a struct of to the
// base struct name, and each field into it's own mapping, for example:
// type "*types.Service", on string field "Name" converts to "service_name=value". There is also an
// additional field called "_pointer" that will return the pointer hex value.
func ToFields(d ...interface{}) map[string]interface{} {
	if !Log.IsLevelEnabled(Logger.DebugLevel) {
		return nil
	}
	fieldKey := make(map[string]interface{})
	for _, v := range d {
		spl := strings.Split(fmt.Sprintf("%T", v), ".")
		trueType := spl[len(spl)-1]
		if !structs.IsStruct(v) {
			continue
		}
		for _, f := range structs.Fields(v) {
			if f.IsExported() && !f.IsZero() && f.Kind() != reflect.Ptr && f.Kind() != reflect.Slice && f.Kind() != reflect.Chan {
				field := strings.ToLower(trueType + "_" + f.Name())
				fieldKey[field] = replaceVal(f.Value())
			}
		}
		fieldKey[strings.ToLower(trueType+"_pointer")] = fmt.Sprintf("%p", v)
	}
	return fieldKey
}

// replaceVal accepts an interface to be converted into human readable type
func replaceVal(d interface{}) interface{} {
	switch v := d.(type) {
	case types.NullBool:
		return v.Bool
	case types.NullString:
		return v.String
	case types.NullFloat64:
		return v.Float64
	case types.NullInt64:
		return v.Int64
	case string:
		if len(v) > 500 {
			return v[:500] + "... (truncated in logs)"
		}
		return v
	case time.Time:
		return v.String()
	case time.Duration:
		return v.String()
	default:
		return d
	}
}

// createLog will create the '/logs' directory based on a directory
func createLog(dir string) error {
	if !FolderExists(dir + "/logs") {
		CreateDirectory(dir + "/logs")
	}
	return nil
}

// InitLogs will create the '/logs' directory and creates a file '/logs/statup.log' for application logging
func InitLogs() error {
	if err := createLog(Directory); err != nil {
		return err
	}
	ljLogger = &lumberjack.Logger{
		Filename:   Directory + logFilePath,
		MaxSize:    16,
		MaxBackups: 5,
		MaxAge:     28,
	}
	mw := io.MultiWriter(os.Stdout, ljLogger)
	Log.SetOutput(mw)

	Log.SetFormatter(&Logger.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	checkVerboseMode()

	LastLines = make([]*logRow, 0)
	return nil
}

// checkVerboseMode will reset the Logging verbose setting. You can set
// the verbose level with "-v 3" or by setting VERBOSE=3 environment variable.
// statping -v 1 (only Warnings)
// statping -v 2 (Info and Warnings, default)
// statping -v 3 (Info, Warnings and Debug)
// statping -v 4 (Info, Warnings, Debug and Traces (SQL queries))
func checkVerboseMode() {
	switch VerboseMode {
	case 1:
		Log.SetLevel(Logger.WarnLevel)
	case 2:
		Log.SetLevel(Logger.InfoLevel)
	case 3:
		Log.SetLevel(Logger.DebugLevel)
	case 4:
		Log.SetReportCaller(true)
		Log.SetLevel(Logger.TraceLevel)
	default:
		Log.SetLevel(Logger.InfoLevel)
	}
	Log.Debugf("logging running in %v mode", Log.GetLevel().String())
}

// CloseLogs will close the log file correctly on shutdown
func CloseLogs() {
	ljLogger.Rotate()
	Log.Writer().Close()
	ljLogger.Close()
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
func GetLastLine() *logRow {
	LockLines.Lock()
	defer LockLines.Unlock()
	if len(LastLines) > 0 {
		return LastLines[len(LastLines)-1]
	}
	return nil
}

type logRow struct {
	Date time.Time
	Line interface{}
}

func newLogRow(line interface{}) (lgRow *logRow) {
	lgRow = new(logRow)
	lgRow.Date = time.Now()
	lgRow.Line = line
	return
}

func (o *logRow) lineAsString() string {
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

func (o *logRow) FormatForHtml() string {
	return fmt.Sprintf("%s: %s", o.Date.Format("2006-01-02 15:04:05"), o.lineAsString())
}
