package main

import (
	"fmt"
)

func main() {
	// REMEMBER
	// convert string to rune type or
	// use range base loop when working with string

	pw := "password"
	pwR := []rune(pw)

	for _, v := range pw {
		fmt.Println(v)
	}
}