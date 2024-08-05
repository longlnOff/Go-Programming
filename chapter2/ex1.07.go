package main

import (
	"fmt"
	"time"
)

func main() {
	// switch combine with initial value
	switch  dayBorn := time.Sunday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Born some other day")
	}
}