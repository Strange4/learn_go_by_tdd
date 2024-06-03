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
	Time   time.Time
	Center Point
	Radius float64
}

func (c *Clock) SecondHand() Point {
	secondHandLength := c.Radius * 0.6
	point := angleToUnitPoint(secondsInRadians(c.Time.Second()))

	return unitPointToHandPoint(secondHandLength, point, c.Center)
}

func (c *Clock) MinuteHand() Point {
	minuteHandLength := c.Radius * 0.5
	point := angleToUnitPoint(minutesInRadians(c.Time.Minute()))

	return unitPointToHandPoint(minuteHandLength, point, c.Center)
}

func unitPointToHandPoint(handLength float64, unitPoint, center Point) Point {
	x := handLength*unitPoint.X + center.X
	y := -handLength*unitPoint.Y + center.Y

	return Point{X: x, Y: y}
}

func secondsInRadians(seconds int) (rads float64) {
	const oneSecondInRads = math.Pi / 30
	return float64(seconds) * oneSecondInRads
}

func minutesInRadians(minutes int) (rads float64) {
	return secondsInRadians(minutes)
}

func angleToUnitPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
