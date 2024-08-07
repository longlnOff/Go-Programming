package main

import "fmt"

func fundMinGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func main() {
	nums := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}

	min := fundMinGeneric(nums)
	fmt.Println(min)
}