package log

import (
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type LogType string

type Log struct {
	log       LogType
	timestamp time.Time
	message   string
	file      string
	line      int
	column    int
	funcName  string
}

const (
	NONE    LogType = "NONE"    // Default log level: do not log
	FATAL   LogType = "FATAL"   // Something failed and we CANNOT continue the execution
	ERROR   LogType = "ERROR"   // Something failed, the UI should show this, and we can continue
	WARNING LogType = "WARNING" // Something failed, the UI don't need to show this, and we can continue
	INFO    LogType = "INFO"    // Information about some important event
	DEBUG   LogType = "DEBUG"   // Information step by step
)

func (l *Log) Timestamp() string {
	return l.timestamp.Format("2006-01-02 15:04:05")
}

func NewLog(log LogType, message string) *Log {
	timestamp := time.Now()
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return &Log{
			log:       log,
			message:   message,
			timestamp: timestamp,
			file:      "unknown",
			line:      0,
			column:    0,
			funcName:  "unknown",
		}
	}
	file = filepath.Base(file)
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	funcPC := uintptr(0)
	if fn != nil {
		funcName = fn.Name()
		funcPC = fn.Entry()
		if last := strings.LastIndex(funcName, "/"); last >= 0 {
			funcName = funcName[last+1:]
		}
	}
	return &Log{
		log:       log,
		message:   message,
		timestamp: timestamp,
		file:      file,
		line:      line,
		column:    int(pc - funcPC),
		funcName:  funcName,
	}
}
