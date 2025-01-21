package main

import (
	"fmt"
	"time"
)

// Starting multiple goroutines.
func main() {
	go printNumbers("Routine1")
	go printNumbers("Routine2")

	// Give goroutines time to finish (quick hack for demo).
	time.Sleep(1 * time.Second)
	fmt.Println("Main function ends.")
}

func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}
