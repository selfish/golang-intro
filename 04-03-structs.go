package main

import "fmt"

// Defining a struct and attaching methods to it.
func main() {
	p := Person{
		Name: "Gopher",
		Age:  5,
	}
	fmt.Println(p.Greet())
}

// Person is a simple struct with Name and Age fields.
type Person struct {
	Name string
	Age  int
}

// Greet is a method on Person (receiver type is *Person).
func (p *Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old!", p.Name, p.Age)
}
