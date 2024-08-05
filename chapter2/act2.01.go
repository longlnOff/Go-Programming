package main

import "fmt"

func main() {
	words := map[string]int {
		"Gonna": 3,
		"You": 3,
		"Given": 2,
		"Up": 4,
	}
	var mostPopularWord string
	var value int = 0
	for k, v := range words {
		if v > value {
			value = v
			mostPopularWord = k
		}
	}
	fmt.Println("Most popular word: ", mostPopularWord)
	fmt.Println("With a count of: ", value)
}
