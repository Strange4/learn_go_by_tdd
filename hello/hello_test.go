package hello

import (
	"hello/assertions"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("With a person to name", func(t *testing.T) {
		want := "Hello, Me"
		got := Hello("Me", "")
		assertions.AssertString(t, got, want)
	})
	t.Run("Default name when none is provided", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")
		assertions.AssertString(t, got, want)
	})

	t.Run("in Spanish testing", func(t *testing.T) {
		want := "Hola, Hermano"
		got := Hello("Hermano", "Spanish")
		assertions.AssertString(t, got, want)
	})

	t.Run("En français SVP", func(t *testing.T) {
		want := "Salut, Jean"
		got := Hello("Jean", "French")
		assertions.AssertString(t, got, want)
	})
}
