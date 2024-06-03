package clock_face

import (
	"io"
	"text/template"
)

type svgClockData struct {
	Center     Point
	SecondHand Point
	MinuteHand Point
	HourHand   Point
	Radius     float64
	Width      float64
	Height     float64
}

const svgTemplate = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 {{printf "%.3f" .Width}} {{printf "%.3f" .Height}}"
     version="2.0">

  <!-- bezel -->
  <circle id="bezel" cx="{{printf "%.3f" .Center.X}}" cy="{{printf "%.3f" .Center.Y}}" r="{{printf "%.3f" .Radius}}" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

  <!-- hour hand -->
  <line id="hour_hand" x1="{{printf "%.3f" .Center.X}}" y1="{{printf "%.3f" .Center.Y}}" x2="{{printf "%.3f" .HourHand.X}}" y2="{{printf "%.3f" .HourHand.Y}}"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- minute hand -->
  <line id="minute_hand" x1="{{printf "%.3f" .Center.X}}" y1="{{printf "%.3f" .Center.Y}}" x2="{{printf "%.3f" .MinuteHand.X}}" y2="{{printf "%.3f" .MinuteHand.Y}}"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- second hand -->
  <line id="second_hand" x1="{{printf "%.3f" .Center.X}}" y1="{{printf "%.3f" .Center.Y}}" x2="{{printf "%.3f" .SecondHand.X}}" y2="{{printf "%.3f" .SecondHand.Y}}"
        style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>`

func SVGWriter(writer io.Writer, clock *Clock, center Point, radius float64) {
	tmpl, err := template.New("svg_clock.svg").Parse(svgTemplate)
	if err != nil {
		panic(err)
	}
	secondHandLength := radius * 0.6
	minuteHandLength := radius * 0.5
	hourHandLength := radius * 0.3

	data := svgClockData{
		Center:     center,
		SecondHand: unitPointToHandPoint(secondHandLength, clock.SecondHand(), center),
		MinuteHand: unitPointToHandPoint(minuteHandLength, clock.MinuteHand(), center),
		HourHand:   unitPointToHandPoint(hourHandLength, clock.HourHand(), center),
	}
	err = tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
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
