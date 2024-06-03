package clock_face

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
			simpleTime(0, 0, 0),
			Point{X: 0, Y: 1},
		},
		{
			"15 seconds",
			simpleTime(0, 0, 15),
			Point{X: 1, Y: 0},
		},
		{
			"30 seconds",
			simpleTime(0, 0, 30),
			Point{X: 0, Y: -1},
		},
		{
			"45 seconds",
			simpleTime(0, 0, 45),
			Point{X: -1, Y: 0},
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

func TestMinuteHand(t *testing.T) {

	cases := []struct {
		name  string
		time  time.Time
		point Point
	}{
		{
			"0 minutes",
			simpleTime(0, 0, 0),
			Point{X: 0, Y: 1},
		},
		{
			"15 minutes",
			simpleTime(0, 15, 0),
			Point{X: 1, Y: 0},
		},
		{
			"30 minutes",
			simpleTime(0, 30, 0),
			Point{X: 0, Y: -1},
		},
		{
			"45 minutes",
			simpleTime(0, 45, 0),
			Point{X: -1, Y: 0},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Clock{testCase.time}
			want := testCase.point
			got := c.MinuteHand()
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
