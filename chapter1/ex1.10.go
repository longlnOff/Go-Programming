package main

import (
	"fmt"
)

func main() {
	count := 5
	count += 5
	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)

	name := "longln"

	fmt.Println(name)

	name += "long"
	fmt.Println(name)
}