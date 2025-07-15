package log

import (
	"fmt"
	"strings"
	"time"
)

type consoleLogger struct {
	isProduction bool
	hasTimestamp bool // show the time stamp of the log
	hasFilepath  bool // show the file that is being called
	hasMethod    bool // show the methods that the current called log is in
}

func (c *consoleLogger) NewLog(logType LogType, message string) error {
	var sbd strings.Builder
	info := Log{log: logType, timestamp: time.Now(), message: message}
	if c.hasTimestamp {
		_, err := sbd.WriteString(fmt.Sprintf("[%s] ", info.Timestamp()))
		if err != nil {
			return err
		}
	}
	sbd.WriteString(fmt.Sprintf("%s\t: %s\n", c.colorStringFromType(logType), info.message))
	fmt.Print(sbd.String())
	return nil
}

func (c *consoleLogger) colorStringFromType(log LogType) string {
	switch log {
	case FATAL:
		return fmt.Sprintf(clrFatal, log)
	case ERROR:
		return fmt.Sprintf(clrError, log)
	case WARNING:
		return fmt.Sprintf(clrWarning, log)
	case INFO:
		return fmt.Sprintf(clrInfo, log)
	case VERBOSE:
		return fmt.Sprintf(clrVerbose, log)
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
	clrVerbose = "\033[35m%s\033[0m"   // Magenta
)
