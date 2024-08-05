package main

import (
	"fmt"
	"time"
)

// use multiple declaration
var (
	Debug bool = true
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}