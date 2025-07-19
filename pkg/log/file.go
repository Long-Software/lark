package log

import (
	"fmt"
	"os"
	"strings"
)

type fileLogger struct {
	isProduction bool
	HasLogFile   bool
	LogFilepath  string // file path to the log file
	hasTimestamp bool   // show the time stamp of the log
	hasFilepath  bool   // show the file that is being called
	hasMethod    bool   // show the methods that the current called log is in
}

func (f *fileLogger) NewLog(logType LogType, message string) error {
	if !f.HasLogFile {
		return nil
	}

	file, err := os.OpenFile(f.LogFilepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	var bd strings.Builder
	entry := NewLog(logType, message)
	// write the timestamp of the log
	if f.hasTimestamp {
		_, err := bd.WriteString(fmt.Sprintf("[%s] ", entry.Timestamp()))
		if err != nil {
			return err
		}
	}

	// write the log type
	_, err = bd.WriteString(fmt.Sprintf("%s\t: ", logType))
	if err != nil {
		return err
	}
	// write the file path of the error
	if f.hasFilepath {
		_, err := bd.WriteString(fmt.Sprintf("[%s] ", entry.file))
		if err != nil {
			return err
		}
	}
	// write the methods and the line that was called before the error
	if f.hasMethod {
		_, err := bd.WriteString(fmt.Sprintf("[%s:%d:%d] ", entry.funcName, entry.line, entry.column))
		if err != nil {
			return err
		}
	}

	bd.WriteString(fmt.Sprintf(": %s\n", entry.message))
	file.WriteString(bd.String())
	return nil
}
