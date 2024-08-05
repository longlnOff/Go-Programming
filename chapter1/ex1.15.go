package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Pointer(count *int) {
	*count += 5
	fmt.Println("add5Pointer: ", *count)
}


// REMEMBER
// if you have a pointer variable or have passed a pointer of a variable to a function,
// any changes made to the pointer will be reflected in the original variable outside the function
func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value: ", count)
	add5Pointer(&count)
	fmt.Println("add5Pointer: ", count)
}