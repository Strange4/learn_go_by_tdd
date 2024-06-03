package clock_face

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

type Clock struct {
	Time time.Time
}

func (c *Clock) SecondHand() Point {
	return angleToUnitPoint(secondsInRadians(c.Time))
}

func (c *Clock) MinuteHand() Point {
	return angleToUnitPoint(minutesInRadians(c.Time))
}

func (c *Clock) HourHand() Point {
	return angleToUnitPoint(hoursInRadians(c.Time))
}

func secondsInRadians(t time.Time) (rads float64) {
	const oneSecondInRads = math.Pi / 30
	return float64(t.Second()) * oneSecondInRads
}

func minutesInRadians(t time.Time) (rads float64) {
	const oneMinuteInRads = math.Pi / 30
	var minutesPassed = oneMinuteInRads * float64(t.Minute())
	return (secondsInRadians(t) / 60) + minutesPassed
}

func hoursInRadians(t time.Time) (rads float64) {
	const oneHourInRads = math.Pi / 6
	var hoursPassed = float64(t.Hour()%12) * oneHourInRads
	return hoursPassed + (minutesInRadians(t) / 12)
}

func angleToUnitPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
