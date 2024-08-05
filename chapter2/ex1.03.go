package main

import "fmt"

func main() {
	input := 0
	if input > 0 {
		fmt.Println("Positive")
	} else if input < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}