package clock

import (
	"fmt"
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

var Center = Point{150, 150}
var SecondHandLength = float64(90)

func (c *Clock) SecondHand() Point {
	const oneSecond = math.Pi / 30
	seconds := float64(c.Time.Second())
	rads := seconds * oneSecond
	x := SecondHandLength*math.Sin(rads) + Center.X
	y := -SecondHandLength*math.Cos(rads) + Center.Y
	return Point{X: x, Y: y}
}

func DrawClock(clock *Clock) string {
	const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

  <!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

  <!-- hour hand -->
  <line x1="150" y1="150" x2="114.150000" y2="132.260000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- minute hand -->
  <line x1="150" y1="150" x2="101.290000" y2="99.730000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>`
	secondHand := clock.SecondHand()
	seconds := fmt.Sprintf(`<!-- second hand -->
	<line x1="150" y1="150" x2="%f" y2="%f"
		  style="fill:none;stroke:#f00;stroke-width:3px;"/>`, secondHand.X, secondHand.Y)
	const svgEnd = `</svg>`

	return svgStart + seconds + svgEnd
}
