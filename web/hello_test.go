package main

import "testing"

func TestHello(t *testing.T) {
	have := Hello("Nick")
	want := "Hello, Nick"

	if have != want {
		t.Errorf("have %q want %q", have, want)
	}
}
