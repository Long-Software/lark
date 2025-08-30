package logger

import "github.com/Long-Software/Bex/packages/log"

var l log.ConsoleLogger

func New(level log.LogLevel, message string) {
	l.NewLog(level, message)
}
