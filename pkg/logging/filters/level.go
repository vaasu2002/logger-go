package filter

import "logger/pkg/logging"

// LevelFilter filters logs based on minimum log level
type LevelFilter struct {
	thresholdLevel logging.LogLevel
}

func NewLevelFilter(thresholdLevel logging.LogLevel) *LevelFilter {
	return &LevelFilter{
		thresholdLevel: thresholdLevel,
	}
}

func (l *LevelFilter) Accept(log *logging.Log) bool {
	return log.Level >= l.thresholdLevel
}
