package greetings

import "fmt"

// Hello returns a greeting for the specified name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Acc&App Go demo.", name)
}
