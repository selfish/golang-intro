package main

import "fmt"

// Slices, their creation, and append.
func main() {
	// Create a slice using shorthand
	nums := []int{1, 2, 3}
	fmt.Println("Initial slice:", nums)

	// Append values
	nums = append(nums, 4, 5)
	fmt.Println("After append:", nums)

	// Using make
	letters := make([]string, 3)
	letters[0] = "a"
	letters[1] = "b"
	letters[2] = "c"
	fmt.Println("Letters:", letters)

	// Slicing a slice
	sub := nums[1:3]
	fmt.Println("Sub-slice of nums[1:3]:", sub)
}
