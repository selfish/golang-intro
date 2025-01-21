package main

import "fmt"

// Map usage (key-value pairs) in Go.
func main() {
	// Create a map of string -> int
	scores := map[string]int{
		"Alice": 10,
		"Bob":   8,
	}
	fmt.Println("Initial map:", scores)

	// Add or update
	scores["Charlie"] = 9
	fmt.Println("Updated map:", scores)

	// Check existence
	if val, ok := scores["Bob"]; ok {
		fmt.Println("Bob's score is", val)
	}

	// Delete a key
	delete(scores, "Alice")
	fmt.Println("Map after deletion:", scores)

	// Iterate over map
	for name, score := range scores {
		fmt.Printf("%s => %d\n", name, score)
	}
}
