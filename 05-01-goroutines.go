package main

import (
	"fmt"
	"time"
)

// Starting multiple goroutines.
func main() {
	defer fmt.Println("Main function ends.")
	go printNumbers("Routine1")
	go printNumbers("Routine2")

	// Give goroutines time to finish (quick hack for demo).
	// Corrrect approach is to use sync.WaitGroup
	time.Sleep(1 * time.Second)
}

func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}
