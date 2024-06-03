package clock_face

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandUnitPoint(t *testing.T) {
	var cases = []struct {
		time  time.Time
		point Point
	}{
		{
			simpleTime(0, 0, 0),
			Point{X: 0, Y: 1},
		},
		{
			simpleTime(0, 0, 15),
			Point{X: 1, Y: 0},
		},
		{
			simpleTime(0, 0, 30),
			Point{X: 0, Y: -1},
		},
		{
			simpleTime(0, 0, 45),
			Point{X: -1, Y: 0},
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			c := Clock{testCase.time}
			want := testCase.point
			got := c.SecondHand()
			if !pointsRoughlyEqual(want, got) {
				t.Errorf("Wanted point %v but got %v", want, got)
			}
		})
	}
}

func TestMinuteHandUnitPoint(t *testing.T) {
	var cases = []struct {
		time  time.Time
		point Point
	}{
		{
			simpleTime(0, 0, 0),
			Point{X: 0, Y: 1},
		},
		{
			simpleTime(0, 15, 0),
			Point{X: 1, Y: 0},
		},
		{
			simpleTime(0, 30, 0),
			Point{X: 0, Y: -1},
		},
		{
			simpleTime(0, 45, 0),
			Point{X: -1, Y: 0},
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			c := Clock{testCase.time}
			want := testCase.point
			got := c.MinuteHand()
			if !pointsRoughlyEqual(want, got) {
				t.Errorf("Wanted point %v but got %v", want, got)
			}
		})
	}
}

func TestHourHandUnitPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{
			simpleTime(12, 0, 0),
			Point{0, 1},
		},
		{
			simpleTime(0, 0, 0),
			Point{0, 1},
		},
		{
			simpleTime(3, 0, 0),
			Point{1, 0},
		},
		{
			simpleTime(6, 0, 0),
			Point{0, -1},
		},
		{
			simpleTime(9, 0, 0),
			Point{-1, 0},
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			c := Clock{testCase.time}
			want := testCase.point
			got := c.HourHand()
			if !pointsRoughlyEqual(want, got) {
				t.Errorf("Wanted point %v but got %v", want, got)
			}
		})
	}
}

func TestSecondHandInRadians(t *testing.T) {
	const oneSecondInRads = 2 * math.Pi / 60
	cases := []struct {
		time time.Time
		rads float64
	}{
		{
			simpleTime(0, 0, 5),
			oneSecondInRads * 5,
		},
		{
			simpleTime(0, 0, 58),
			oneSecondInRads * 58,
		},
		{
			simpleTime(0, 0, 23),
			oneSecondInRads * 23,
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			want := testCase.rads
			got := secondsInRadians(testCase.time)
			if !floatRoughlyEqual(want, got) {
				t.Errorf("Wanted angle %v but got %v", want, got)
			}
		})
	}
}

func TestMinuteHandInRadians(t *testing.T) {
	const secondInFluenceInRads = 2 * math.Pi / (60 * 60)
	const minuteInRads = 2 * math.Pi / 60
	cases := []struct {
		time time.Time
		rads float64
	}{
		{
			simpleTime(0, 30, 0),
			minuteInRads * 30,
		},
		{
			simpleTime(0, 0, 20),
			secondInFluenceInRads * 20,
		},
		{
			simpleTime(0, 1, 20),
			(secondInFluenceInRads * 20) + minuteInRads,
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			want := testCase.rads
			got := minutesInRadians(testCase.time)
			if !floatRoughlyEqual(want, got) {
				t.Errorf("Wanted angle %v but got %v", want, got)
			}
		})
	}
}

func TestHourHandInRadians(t *testing.T) {
	const secondInfluenceInRads = 2 * math.Pi / (60 * 60 * 12)
	const minuteInfluenceRads = 2 * math.Pi / (60 * 12)
	const hourInRads = 2 * math.Pi / 12
	cases := []struct {
		time time.Time
		rads float64
	}{
		{
			simpleTime(0, 0, 0),
			0,
		},
		{
			simpleTime(6, 0, 0),
			hourInRads * 6,
		},
		{
			simpleTime(9, 0, 0),
			hourInRads * 9,
		},
		{
			simpleTime(1, 17, 0),
			(minuteInfluenceRads * 17) + hourInRads,
		},
		{
			simpleTime(1, 23, 15),
			hourInRads + (23 * minuteInfluenceRads) + (15 * secondInfluenceInRads),
		},
		{
			simpleTime(0, 1, 30),
			math.Pi / ((6 * 60 * 60) / 90),
		},
	}
	for _, testCase := range cases {
		t.Run(timeToTestName(testCase.time), func(t *testing.T) {
			want := testCase.rads
			got := hoursInRadians(testCase.time)
			if !floatRoughlyEqual(want, got) {
				t.Errorf("Wanted angle %v but got %v", want, got)
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

func timeToTestName(t time.Time) string {
	return t.Format(time.TimeOnly)
}
