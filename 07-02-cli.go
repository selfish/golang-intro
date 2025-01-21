package main

import (
	"flag"
	"fmt"
)

// A simple CLI using the built-in flag package.
func main() {
	name := flag.String("name", "Gopher", "your name")
	age := flag.Int("age", 1, "your age")

	flag.Parse()

	fmt.Printf("Hello %s, you are %d years old.\n", *name, *age)
}
