package assertions

import "testing"

func AssertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func AssertInteger(t *testing.T, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %v, want %v", actual, expected)
	}
}
