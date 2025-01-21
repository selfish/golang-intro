package main

import "fmt"

// Basic pointer usage in Go.
func main() {
	x := 10
	fmt.Println("Before:", x)
	modifyValue(&x)
	fmt.Println("After:", x)
}

func modifyValue(num *int) {
	*num = *num + 5
}
