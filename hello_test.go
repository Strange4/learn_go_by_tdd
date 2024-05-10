package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Me."
	got := Hello("Me")
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
