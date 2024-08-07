package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type person struct {
	name string
	age int
	isMarried bool
}

func main() {
	p := person{
		name: "longln",
		age: 24,
		isMarried: false,
	}
	fmt.Println(p.Speak())
	fmt.Println(p)
}


// Stringer{} interface is an interface that is in the Go language.
// It can be used for formatting when printing values.
// Create a String() method for person and return a string value.
// This will satisfy the Stringer{} interface, which will now 
// allow it to be called by the fmt.Println() method.
func (p person) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarried status: %v ", p.name, p.age, p.isMarried)
}

func (p person) Speak() string {
	return "Hi my name is: " + p.name
}