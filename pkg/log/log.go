package log

import (
	"fmt"
)

func (l *Logger) NewLog(log LogType, message string) {
	var lgs []LoggerInt

	lgs = append(lgs, &consoleLogger{
		isProduction: l.IsProduction,
		hasTimestamp: l.HasTimestamp,
		hasFilepath:  l.HasFilepath,
		hasMethod:    l.HasMethod,
	})

	for _, lg := range lgs {
		err := lg.NewLog(log, message)
		if err != nil {
			fmt.Printf("Error from logger: %v", err)
		}
	}
}

func (l *Log) Timestamp() string {
	return l.timestamp.Format("2006-01-02 15:04:05")
}
