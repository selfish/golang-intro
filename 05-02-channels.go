package main

import (
	"fmt"
)

// Communication via channels.
func main() {
	ch := make(chan int)

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
