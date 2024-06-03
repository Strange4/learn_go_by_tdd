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
	// secondHandLength := c.Radius * 0.6
	return angleToUnitPoint(secondsInRadians(c.Time.Second()))
}

func (c *Clock) MinuteHand() Point {
	// minuteHandLength := c.Radius * 0.5
	return angleToUnitPoint(minutesInRadians(c.Time.Minute()))

}

func secondsInRadians(seconds int) (rads float64) {
	const oneSecondInRads = math.Pi / 30
	return float64(seconds) * oneSecondInRads
}

func minutesInRadians(minutes int) (rads float64) {
	return secondsInRadians(minutes)
}

func angleToUnitPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
