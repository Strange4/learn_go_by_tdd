package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("With a person to name", func(t *testing.T) {
		want := "Hello, Me"
		got := Hello("Me")
		assertString(t, got, want)
	})
	t.Run("Default name when none is provided", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("")
		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
