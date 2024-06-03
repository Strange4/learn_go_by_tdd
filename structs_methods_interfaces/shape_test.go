package shape

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %v and wanted %v", got, want)
	}
}

func TestArea(t *testing.T) {

	tests := []struct {
		shape    Shape
		want     float64
		testName string
	}{
		{shape: Rectangle{40, 20}, want: 800.0, testName: "Rectangles"},
		{shape: Circle{math.Sqrt(10)}, want: 31.41592653589793, testName: "Circle"},
		{shape: Triangle{10, 10}, want: 50.0, testName: "Triangles"},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		precisionDelta := 0.000000001
		pass := got > want-precisionDelta && got < want+precisionDelta
		if !pass {
			t.Errorf("for shape %#v got %v and wanted %v", shape, got, want)
		}
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}
