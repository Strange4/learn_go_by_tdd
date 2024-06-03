package clock_face

import (
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	center := Point{150, 150}
	radius := 150.0
	secondHandLength := radius * 0.6

	cases := []struct {
		name  string
		time  time.Time
		point Point
	}{
		{
			"0 seconds",
			simpleTime(0, 0, 0),
			Point{X: center.X, Y: center.Y - secondHandLength},
		},
		{
			"30 seconds",
			simpleTime(0, 0, 30),
			Point{X: center.X, Y: center.Y + secondHandLength},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Clock{testCase.time, center, radius}
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

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2000, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}
