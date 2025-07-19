package log

import "fmt"

type Logger struct {
	IsProduction bool
	HasLogFile   bool
	LogFilepath  string // file path to the log file
	HasTimestamp bool   // show the time stamp of the log
	HasFilepath  bool   // show the file that is being called
	HasMethod    bool   // show the methods that the current called log is in
}

type LoggerInt interface {
	NewLog(LogType, string) error
}

func (l *Logger) NewLog(log LogType, message string) {
	var lgs []LoggerInt

	lgs = append(lgs, &consoleLogger{
		isProduction: l.IsProduction,
		hasTimestamp: l.HasTimestamp,
		hasFilepath:  l.HasFilepath,
		hasMethod:    l.HasMethod,
	}, )

	if l.HasLogFile {
		lgs = append(lgs, &consoleLogger{
		isProduction: l.IsProduction,
		hasTimestamp: l.HasTimestamp,
		hasFilepath:  l.HasFilepath,
		hasMethod:    l.HasMethod,
	}, &fileLogger{
		isProduction: l.IsProduction,
		HasLogFile:   l.HasLogFile,
		LogFilepath:  l.LogFilepath,
		hasTimestamp: l.HasTimestamp,
		hasFilepath:  l.HasFilepath,
		hasMethod:    l.HasMethod,
	})
	}
	for _, lg := range lgs {
		err := lg.NewLog(log, message)
		if err != nil {
			fmt.Printf("Error from logger: %v", err)
		}
	}
}
