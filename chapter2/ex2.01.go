package main

import (
	"fmt"
)

func main() {
	input := 5
	if input % 2 == 0 {
		fmt.Println("Even")
	}
	if input % 2 != 0 {
		fmt.Println("Odd")
	}
}