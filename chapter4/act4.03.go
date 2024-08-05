package main

import (
	"fmt"
)

func main() {
	week1 := []string {
		 "Monday",
		 "Tuesday",
		 "Wednesday",
		 "Thursday",
		 "Friday", 
		 "Saturday",
		 "Sunday",
	}

	week2 := []string{}
	week2 = append(week2, week1[6])
	week2 = append(week2, week1[:6]...)


	fmt.Println(week1)
	fmt.Println(week2)

}