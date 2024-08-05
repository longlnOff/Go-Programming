package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	// convert string to rune type, which is safe for multi-byte (UTF-8) characters
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		// use unicode package to check conditions
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		// accept symbols and punctions
		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("password") {
		fmt.Println("Good password")
	} else {
		fmt.Println("Bad password")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("Good password")
	} else {
		fmt.Println("Bad password")
	}
}