package clock_face

import (
	"fmt"
	"io"
)

func SVGWriter(writer io.Writer, clock *Clock, center Point, radius float64) {
	writer.Write([]byte(svgStart))
	writeSecondHand(writer, clock, center, radius)
	writeMinuteHand(writer, clock, center, radius)
	writeHourHand(writer, clock, center, radius)
	writer.Write([]byte(svgEnd))
}

func writeSecondHand(writer io.Writer, clock *Clock, center Point, radius float64) {
	secondHandLength := radius * 0.6
	secondHand := unitPointToHandPoint(secondHandLength, clock.SecondHand(), center)

	fmt.Fprintf(writer, `
	<!-- second hand -->
	<line id="second_hand" x1="%.3f" y1="%.3f" x2="%.3f" y2="%.3f"
		  style="fill:none;stroke:#f00;stroke-width:3px;"/>
		  `, center.X, center.Y, secondHand.X, secondHand.Y)
}

func writeMinuteHand(writer io.Writer, clock *Clock, center Point, radius float64) {
	minuteHandLength := radius * 0.5
	minuteHand := unitPointToHandPoint(minuteHandLength, clock.MinuteHand(), center)
	fmt.Fprintf(writer, `
	<!-- minute hand -->
	<line id="minute_hand" x1="%.3f" y1="%.3f" x2="%.3f" y2="%.3f"
		style="fill:none;stroke:#000;stroke-width:7px;"/>
	`, center.X, center.Y, minuteHand.X, minuteHand.Y)
}

func writeHourHand(writer io.Writer, clock *Clock, center Point, radius float64) {
	hourHandLength := radius * 0.3
	hourHand := unitPointToHandPoint(hourHandLength, clock.HourHand(), center)
	fmt.Fprintf(writer, `
	<!-- hour hand -->
	<line id="hour_hand" x1="%.3f" y1="%.3f" x2="%.3f" y2="%.3f"
		style="fill:none;stroke:#000;stroke-width:7px;"/>
	`, center.X, center.Y, hourHand.X, hourHand.Y)
}

func unitPointToHandPoint(handLength float64, unitPoint, center Point) Point {
	x := handLength*unitPoint.X + center.X
	y := -handLength*unitPoint.Y + center.Y

	return Point{X: x, Y: y}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

	<!-- bezel -->
	<circle id="bezel" cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>
`

const svgEnd = `</svg>`
