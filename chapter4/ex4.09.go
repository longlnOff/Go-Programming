package main

import (
	"fmt"
	"os"
)

func getPassedArgs1() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	fmt.Println(args)

	return args
}

func getLocales(extraLicales []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLicales...)

	return locales
}

func main() {
	locales := getLocales(getPassedArgs1())
	fmt.Println("Locales to use:", locales)
}