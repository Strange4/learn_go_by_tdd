package clock_face

import (
	"fmt"
	"math"
	"testing"
	"time"
)

var secondAndMinuteCases = []struct {
	time  int // from 0 to 59 inclusive
	point Point
}{
	{
		0,
		Point{X: 0, Y: 1},
	},
	{
		15,
		Point{X: 1, Y: 0},
	},
	{
		30,
		Point{X: 0, Y: -1},
	},
	{
		45,
		Point{X: -1, Y: 0},
	},
}

func TestSecondHand(t *testing.T) {
	for _, testCase := range secondAndMinuteCases {
		name := fmt.Sprintf("at %d seconds", testCase.time)
		t.Run(name, func(t *testing.T) {
			c := Clock{simpleTime(0, 0, testCase.time)}
			want := testCase.point
			got := c.SecondHand()
			if !pointsRoughlyEqual(want, got) {
				t.Errorf("Wanted point %v but got %v", want, got)
			}
		})
	}
}

func TestMinuteHand(t *testing.T) {
	for _, testCase := range secondAndMinuteCases {
		name := fmt.Sprintf("at %d minutes", testCase.time)
		t.Run(name, func(t *testing.T) {
			c := Clock{simpleTime(0, testCase.time, 0)}
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
