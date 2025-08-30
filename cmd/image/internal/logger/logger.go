package logger

import "github.com/Long-Software/Bex/packages/log"

var logger = log.ConsoleLogger{}

func NewLog(level log.LogLevel, message string) {
	logger.NewLog(level, message)
}
