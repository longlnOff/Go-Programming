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

	// REMEMBER: always check nil pointer before dereferencing
	if count1 != nil {
		fmt.Printf("count1: %#v\n", count1)
	}

	if count2 != nil {
		fmt.Printf("count2: %#v\n", count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v\n", count3)
	}

	if t != nil {
		// Here we call a method of a pointer, so we need to dereference the pointer
		fmt.Printf("t: %#v\n", t.String())
	}


}