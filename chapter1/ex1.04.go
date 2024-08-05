package main

import (
	"fmt"
	"time"
)

// use multiple declaration
// skipping the type or value when declaring
var (
	Debug bool = true
	LogLevel = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}