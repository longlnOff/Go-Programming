package main

import (
	"fmt"
)

func main() {
	itemSold()
}

func itemSold() {
	items := make(map[string]int)
	items["Joihn"] = 41
	items["Doe"] = 42
	items["Smith"] = 43
	items["Wick"] = 44
	for k, v := range items {
		fmt.Printf("%s sold %d item and ", k, v)

		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations.")
		} else if v > 100 {
			fmt.Println("exceeds expectations.")
		}
	}
}