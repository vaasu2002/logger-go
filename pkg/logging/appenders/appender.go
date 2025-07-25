package appenders

import "logger/pkg/logging"

// ILogAppender defines the interface for log appenders
type ILogAppender interface {
	Setup() error
	Append(log *logging.Log) error
}
