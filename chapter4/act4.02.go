package main

import (
	"fmt"
	"os"
)

func main() {
	userIDsMap := map[string]string {
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program ID")
		os.Exit(1)
	}

	name, exist := userIDsMap[os.Args[1]]
	if exist == false {
		fmt.Printf("Key %v doesn't exist!\n", os.Args[1])
	} else {
		fmt.Println("Hi,", name)
	}
}