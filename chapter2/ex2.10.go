package main

import "fmt"

// REMEMBER
// map data type is an unordered collection of key-value pairs
// so the order of iteration is not guaranteed
func main() {
	config := map[string]string {
		"debug": "true",
		"loglevel": "info",
		"version": "1.0.0",
	}

	for key, value := range config {
		fmt.Printf("%s: %s\n", key, value)
	}
}