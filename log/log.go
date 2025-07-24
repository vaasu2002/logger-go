package log

import (
	"fmt"
	"logger/loglevel"
	"time"
)

type Log struct {
	Message   string
	Source    string
	Level     loglevel.LogLevel
	Timestamp time.Time
}

// NewLog creates a new Log instance
func NewLog(message string, source string, level loglevel.LogLevel) *Log {
	return &Log{
		Message:   message,
		Source:    source,
		Level:     level,
		Timestamp: time.Now(),
	}
}

// String returns a string representation of the Log
func (l *Log) String() string {
	return fmt.Sprintf("[%s] %s [%s] %s",
		l.Timestamp.Format(time.RFC3339),
		loglevel.LogLevelToString(l.Level),
		l.Source,
		l.Message,
	)
}
