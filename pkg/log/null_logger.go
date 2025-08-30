package log

import "time"

func (n *NullLogger) NewLog(level LogLevel, message string) {
	n.Logs = append(n.Logs, Log{level: level, timestamp: time.Now(), message: message})
}
