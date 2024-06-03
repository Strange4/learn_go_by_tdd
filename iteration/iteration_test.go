package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("yes", 3)
	expected := "yesyesyes"
	assertString(t, actual, expected)
}

func ExampleRepeat() {
	repeated := Repeat("Hello", 4)
	fmt.Println(repeated)
	// output: HelloHelloHelloHello
}

func BenchmarkRepeat(b *testing.B) {
	const repeatCount = 4
	for i := 0; i < b.N; i++ {
		Repeat("f", repeatCount)
	}
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
