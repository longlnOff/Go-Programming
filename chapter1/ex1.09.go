package main

import "fmt"

func main() {
	var total float64 = 2 * 13
	fmt.Println("Sub: ", total)

	// Drinks
	total = total + (4 * 2.25)
	fmt.Println("Sub: ", total)

	// Discount 5
	total = total - 5
	fmt.Println("Sub: ", total)

	// tip 10%
	tip := total * 0.1
	fmt.Println("tip: ", tip)

	// add tip for total
	total = total + tip
	fmt.Println("Total: ", total)

	// Split bill
	split := total / 2
	fmt.Println("Split: ", split)

	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount  + 1
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("with this visit, you've earned a reward!")
	}
}