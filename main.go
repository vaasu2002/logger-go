package main

import (
	"fmt"
	"logger/pkg/logging"
)

func main() {
	fmt.Println(logging.DEBUG)
	fmt.Println(logging.LogLevelToString(logging.DEBUG))
	logEntry := logging.NewLog("Hello, World!", "main", logging.DEBUG)
	fmt.Println(logEntry.String())
}
