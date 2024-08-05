package main

import (
	"fmt"
)

func main() {
	var a int8 = 125
	var b uint8 = 125

	for i := 0; i < 5; i++ {
		a++
		b++
	fmt.Println(i, ")", "int8 ", a, "uint8 ", b)
	}
}