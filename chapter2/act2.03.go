package main

import (
	"fmt"
)

func main() {
	// define a slice with unsorted numbers in it
	slice := []int{5,8,2,4,0,1,3,7,9,6}
	// print the unsorted slice
	fmt.Println(slice)
	// sort the slice
	sortSlice(slice)
	// print the sorted slice
	fmt.Println(slice)
}

func sortSlice(slice []int) {
	for i1 := 0; i1 < len(slice); i1++ {
		for i2 := i1; i2 < len(slice); i2++ {
			if slice[i2] < slice[i1] {
				slice[i2], slice[i1] = slice[i1], slice[i2]
			}
		}
	}
}