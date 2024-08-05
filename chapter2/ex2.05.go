package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday
	switch  dayBorn {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
		fallthrough
	case time.Thursday:
		fmt.Println("Thursday")
		fallthrough
	case time.Friday:
		fmt.Println("Friday")
		fallthrough
	case time.Saturday:
		fmt.Println("Saturday")
		fallthrough
	case time.Sunday:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}