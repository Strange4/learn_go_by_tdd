package integers

import (
	"fmt"
	"hello/assertions"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(123, 321)
	expected := 444
	assertions.AssertInteger(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// output: 4
}
