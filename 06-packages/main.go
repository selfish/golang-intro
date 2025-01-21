package main

import (
	"fmt"
	"go/intro/greetings"
)

func main() {
	msg := greetings.Hello("Go Developer")
	fmt.Println(msg)
}
