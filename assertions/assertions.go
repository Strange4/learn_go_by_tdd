package assertions

import (
	"reflect"
	"testing"
)

func AssertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func AssertInteger(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Wanted no error but got: %v", err)
	}
}
