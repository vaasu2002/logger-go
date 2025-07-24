package main

import (
	"fmt"
	"logger/loglevel"
)

func main() {
	fmt.Println(loglevel.DEBUG)
	fmt.Println(loglevel.LogLevelToString(loglevel.DEBUG))
}
