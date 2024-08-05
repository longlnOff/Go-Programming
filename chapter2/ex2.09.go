package main

import "fmt"

func main() {
	names := []string{"Jim", "Pam", "Dwight", "Michael", "Andy", "Oscar", "Toby"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}