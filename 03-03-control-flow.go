package main

import "fmt"

// Quick examples of if, for, and switch in Go.
func main() {
	x := 10

	// if-else
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less or equal to 5")
	}

	// for loop (Go has no 'while')
	for i := 0; i < 3; i++ {
		fmt.Println("Loop i =", i)
	}

	// switch
	switch x {
	case 10:
		fmt.Println("x is exactly 10")
	case 5:
		fmt.Println("x is 5")
	default:
		fmt.Println("x is something else")
	}
}
