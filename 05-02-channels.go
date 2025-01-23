package main

import (
	"fmt"
)

// Communication via channels.
func main() {
	ch := make(chan int)

	// Or  with Buffer:
	// ch := make(chan int, 4)

	// Start a goroutine that sends data on the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // close channel to signal completion
	}()

	// Receive data from the channel
	for num := range ch {
		fmt.Println("Received:", num)
	}

	fmt.Println("Channel closed, program ends.")
}
