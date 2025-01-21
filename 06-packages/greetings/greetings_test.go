package greetings

import "testing"

// Basic test example using the "testing" package.
func TestHello(t *testing.T) {
	got := Hello("Test")
	want := "Hello, Test!"

	if got[:len(want)] != want { // Quick substring check
		t.Errorf("Hello('Test') = %q, want %q", got, want)
	}
}
