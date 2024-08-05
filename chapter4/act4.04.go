package main

import (
	"fmt"
)

func main() {
	sliceGB := []string {
		"Good","Good","Bad", "Good","Good",
	}
	fmt.Println("Original slice:", sliceGB)
	sliceGB = append(sliceGB[:2], sliceGB[3:]...)
	fmt.Println("New slice:", sliceGB)

}