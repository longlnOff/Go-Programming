package main

import (
	"fmt"
)

const (
	CakeTax = 0.075
	MilkTax = 0.015
	ButterTax = 0.02
)

func main() {
	var (
		cakeCost = 0.99
		milkCost = 2.75
		butterCost = 0.87
	)

	totalTax := CakeTax * cakeCost + MilkTax * milkCost + ButterTax * butterCost
	fmt.Println("Sales Tax Total: ", totalTax)

}