package main

import (
	"errors"
	"fmt"
)

// Functions, multiple returns, error handling convention in Go.
func main() {
	sum, err := addPositiveNumbers(5, -1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
	}
}

func addPositiveNumbers(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("arguments must be non-negative")
	}
	return a + b, nil
}
