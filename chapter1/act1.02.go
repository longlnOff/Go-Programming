package main

import (
	"fmt"
)

func main() {
	a, b := 5, 10
	fmt.Println("Before swap: ", a, b)
	// call swap here
	swap(&a, &b)
	fmt.Println("After swap: ", a, b)
}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}