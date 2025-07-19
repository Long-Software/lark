package log

import (
	"fmt"
	"strings"
)

type consoleLogger struct {
	isProduction bool
	hasTimestamp bool // show the time stamp of the log
	hasFilepath  bool // show the file that is being called
	hasMethod    bool // show the methods that the current called log is in
}

func (c *consoleLogger) NewLog(logType LogType, message string) error {
	var bd strings.Builder
	entry := NewLog(logType, message)
	// write the timestamp of the log
	if c.hasTimestamp {
		_, err := bd.WriteString(fmt.Sprintf("[%s] ", entry.Timestamp()))
		if err != nil {
			return err
		}
	}

	// write the log type
	_, err := bd.WriteString(fmt.Sprintf("%s\t: ", c.logTypeWithColor(logType)))
	if err != nil {
		return err
	}
	// write the file path of the error
	if c.hasFilepath {
		_, err := bd.WriteString(fmt.Sprintf("[%s] ", entry.file))
		if err != nil {
			return err
		}
	}
	// write the methods and the line that was called before the error
	if c.hasMethod {
			_, err := bd.WriteString(fmt.Sprintf("[%s:%d:%d] ", entry.funcName, entry.line, entry.column))
		if err != nil {
			return err
		}
	}

	bd.WriteString(fmt.Sprintf(": %s\n",entry.message))
	fmt.Print(bd.String())
	return nil
}

func (c *consoleLogger) logTypeWithColor(log LogType) string {
	switch log {
	case FATAL:
		return fmt.Sprintf(clrFatal, log)
	case ERROR:
		return fmt.Sprintf(clrError, log)
	case WARNING:
		return fmt.Sprintf(clrWarning, log)
	case INFO:
		return fmt.Sprintf(clrInfo, log)
	case DEBUG:
		return fmt.Sprintf(clrDebug, log)
	default:
		return fmt.Sprintf(clrNone, log)
	}
}

const (
	clrNone    = "%s"                  // NONE
	clrFatal   = "\033[1;41m%s\033[0m" // Bright Red
	clrError   = "\033[31m%s\033[0m"   // Red
	clrWarning = "\033[33m%s\033[0m"   // Yellow
	clrInfo    = "\033[34m%s\033[0m"   // Blue
	clrDebug = "\033[35m%s\033[0m"   // Magenta
)
