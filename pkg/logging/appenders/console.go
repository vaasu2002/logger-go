package appenders

import (
	"fmt"
	"logger/pkg/logging"
)

// ConsoleAppender writes logs to console/stdout
type ConsoleAppender struct{}

func (c *ConsoleAppender) Setup() error {
	return nil
}

func (c *ConsoleAppender) Append(log *logging.Log) error {
	fmt.Println(log.String())
	return nil
}
