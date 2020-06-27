package utils

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/getsentry/sentry-go"
	Logger "github.com/sirupsen/logrus"
	"github.com/statping/statping/types/null"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	Log          = Logger.StandardLogger()
	ljLogger     *lumberjack.Logger
	LastLines    []*logRow
	LockLines    sync.Mutex
	VerboseMode  int
	Version      string
	allowReports bool
)

const (
	logFilePath   = "/logs/statping.log"
	errorReporter = "https://ddf2784201134d51a20c3440e222cebe@sentry.statping.com/4"
)

func SentryInit(v *string, allow bool) {
	allowReports = allow
	if v != nil {
		if *v == "" {
			*v = "development"
		}
		Version = *v
	}
	goEnv := Params.GetString("GO_ENV")
	allowReports := Params.GetBool("ALLOW_REPORTS")
	if allowReports || allow || goEnv == "test" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              errorReporter,
			Environment:      goEnv,
			Release:          Version,
			AttachStacktrace: true,
		}); err != nil {
			Log.Errorln(err)
		}
		Log.Infoln("Error Reporting initiated, thank you!")
	}
}

func SentryErr(err error) {
	if !allowReports {
		return
	}
	sentry.CaptureException(err)
}

func SentryLogEntry(entry *Logger.Entry) {
	e := sentry.NewEvent()
	e.Message = entry.Message
	e.Release = Version
	e.Contexts = entry.Data
	sentry.CaptureEvent(e)
}

type hook struct {
	Entries []Logger.Entry
	mu      sync.RWMutex
}

func (t *hook) Fire(e *Logger.Entry) error {
	pushLastLine(e.Message)
	if e.Level == Logger.ErrorLevel && allowReports {
		SentryLogEntry(e)
	}
	return nil
}

func (t *hook) Levels() []Logger.Level {
	return Logger.AllLevels
}

// ToFields accepts any amount of interfaces to create a new mapping for log.Fields. You will need to
// turn on verbose mode by starting Statping with "-v". This function will convert a struct of to the
// base struct name, and each field into it's own mapping, for example:
// type "*services.Service", on string field "Name" converts to "service_name=value". There is also an
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
	case null.NullBool:
		return v.Bool
	case null.NullString:
		return v.String
	case null.NullFloat64:
		return v.Float64
	case null.NullInt64:
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
		return CreateDirectory(dir + "/logs")
	}
	return nil
}

// InitLogs will create the '/logs' directory and creates a file '/logs/statup.log' for application logging
func InitLogs() error {
	InitEnvs()
	if Params.GetBool("DISABLE_LOGS") {
		return nil
	}
	if err := createLog(Directory); err != nil {
		return err
	}
	ljLogger = &lumberjack.Logger{
		Filename:   Directory + logFilePath,
		MaxSize:    Params.GetInt("LOGS_MAX_SIZE"),
		MaxBackups: Params.GetInt("LOGS_MAX_COUNT"),
		MaxAge:     Params.GetInt("LOGS_MAX_AGE"),
	}

	mw := io.MultiWriter(os.Stdout, ljLogger)
	Log.SetOutput(mw)

	Log.SetFormatter(&Logger.TextFormatter{
		ForceColors:   !Params.GetBool("DISABLE_COLORS"),
		DisableColors: Params.GetBool("DISABLE_COLORS"),
	})
	checkVerboseMode()
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
	if ljLogger != nil {
		ljLogger.Rotate()
		Log.Writer().Close()
		ljLogger.Close()
	}
	sentry.Flush(5 * time.Second)
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
