package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Guy")
	want := "Hello, Guy."
	got := buffer.String()
	if got != want {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}
