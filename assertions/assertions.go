package assertions

import (
	"reflect"
	"testing"
)

func AssertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func AssertInteger(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
