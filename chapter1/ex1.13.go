package main

import (
	"fmt"
	"time"
)

func main() {
	// three different ways to of creating a pointer
	var count1 *int

	count2 := new(int)

	countTemp := 5				// remember, you can't take the address of a literal number
	count3 := &countTemp		// so &5 is wrong

	// Some special types allow create pointer without teemporary variable
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)

	fmt.Printf("count2: %#v\n", count2)

	fmt.Printf("count3: %#v\n", count3)

	fmt.Printf("t: %#v\n", t)
}