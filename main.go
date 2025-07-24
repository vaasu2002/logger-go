package main

import (
	"fmt"
	"logger/log"
	"logger/loglevel"
)

func main() {
	fmt.Println(loglevel.DEBUG)
	fmt.Println(loglevel.LogLevelToString(loglevel.DEBUG))
	log := log.NewLog("Hello, World!", "main", loglevel.DEBUG)
	fmt.Println(log.String())
}
