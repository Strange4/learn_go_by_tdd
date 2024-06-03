package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(123, 321)
	expected := 444
	assertInteger(t, sum, expected)
}

func assertInteger(t *testing.T, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %v, want %v", actual, expected)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// output: 4
}
