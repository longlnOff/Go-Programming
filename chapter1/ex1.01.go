package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Seed the random number
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	// Generate random number between 1 and 5
	r := seed.Intn(5) + 1

	// Create a string with the number of starts we need
	stars := strings.Repeat("*", r)

	// Print starts string
	fmt.Println(stars)
}