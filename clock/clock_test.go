package clock

import (
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		name  string
		time  time.Time
		point Point
	}{
		{
			"0 seconds",
			time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			Point{X: Center.X, Y: Center.Y - SecondHandLength},
		},
		{
			"30 seconds",
			time.Date(2000, time.January, 1, 0, 0, 30, 0, time.UTC),
			Point{X: Center.X, Y: Center.Y + SecondHandLength},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Clock{testCase.time}
			want := testCase.point
			got := c.SecondHand()
			if !pointsRoughlyEqual(want, got) {
				t.Errorf("Wanted point %v but got %v", want, got)
			}
		})
	}
}

func pointsRoughlyEqual(a, b Point) bool {
	return floatRoughlyEqual(a.X, b.X) && floatRoughlyEqual(a.Y, b.Y)
}

func floatRoughlyEqual(a, b float64) bool {
	const deviation = 1e-6

	return math.Abs(a-b) <= deviation
}
