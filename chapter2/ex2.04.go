package main

import (
	"fmt"
	"errors"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("Negative value")
	} else if input > 100 {
		return errors.New("Value too large")
	} else if input % 7 == 0 {
		return errors.New("Value divisible by 7")
	} else {
		return nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}