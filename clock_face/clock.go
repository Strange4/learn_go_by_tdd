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
	const oneSecond = math.Pi / 30
	seconds := float64(c.Time.Second())
	rads := seconds * oneSecond

	x := secondHandLength*math.Sin(rads) + c.Center.X
	y := -secondHandLength*math.Cos(rads) + c.Center.Y
	return Point{X: x, Y: y}
}
