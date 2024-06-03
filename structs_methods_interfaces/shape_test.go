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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		precisionDelta := 0.000000001
		pass := got > want-precisionDelta && got < want+precisionDelta
		if !pass {
			t.Errorf("got %v and wanted %v", got, want)
		}
	}
	t.Run("for rectangles", func(t *testing.T) {
		rectangle := Rectangle{40, 20}
		want := 800.0
		checkArea(t, rectangle, want)
	})

	t.Run("for circles", func(t *testing.T) {
		circle := Circle{math.Sqrt(10)}
		want := 31.41592653589793
		checkArea(t, circle, want)
	})
}
