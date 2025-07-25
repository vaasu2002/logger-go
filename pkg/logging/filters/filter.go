package filter

import "logger/pkg/logging"

// ILogFilter defines the interface for log filters
type ILogFilter interface {
	Accept(log *logging.Log) bool
}
