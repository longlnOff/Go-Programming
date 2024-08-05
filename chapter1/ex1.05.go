package main

import (
	"fmt"
	"time"
)

// short variable declaration 
// always infers the type from a required initial value

func main() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)

}