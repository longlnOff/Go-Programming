package main

import (
	"fmt"
)

var foo string = "bar"	// var declaration can be used outside any function

func main() {
	var baz string = "qux"
	fmt.Println(foo, baz)
}