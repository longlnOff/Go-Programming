package main
import (
  "fmt"
  "os"
)

func getUsers1() map[string]string {
	users := map[string]string {
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	users["073"] = "Tracy"

	return users
}

func getUser(id string) (string, bool) {
	users := getUsers1()

	user, exists := users[id]

	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	
	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers1() {
				fmt.Println("  ID:", key, "Name:", value)
			}
			os.Exit(1)
	}
	fmt.Println("Name:", name)
}