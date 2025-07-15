package log

import "time"

type Logger struct {
	IsProduction bool
	HasLogFile   bool
	LogFilepath  string // file path to the log file
	HasTimestamp bool   // show the time stamp of the log
	HasFilepath  bool   // show the file that is being called
	HasMethod    bool   // show the methods that the current called log is in
}

type LogType string

type Log struct {
	log       LogType
	timestamp time.Time
	message   string
}

const (
	NONE    LogType = "NONE"    // Default log level: do not log
	FATAL   LogType = "FATAL"   // Something failed and we CANNOT continue the execution
	ERROR   LogType = "ERROR"   // Something failed, the UI should show this, and we can continue
	WARNING LogType = "WARNING" // Something failed, the UI don't need to show this, and we can continue
	INFO    LogType = "INFO"    // Information about some important event
	VERBOSE LogType = "VERBOSE" // Information step by step
)

type LoggerInt interface {
	NewLog(LogType, string) error
}
