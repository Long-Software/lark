package utils

import "github.com/Long-Software/lark/pkg/log"

var Log = log.Logger{IsProduction: false, HasTimestamp: true}

func NewLog(l log.LogType, msg string) {
	Log.NewLog(l, msg)
}
