package clock_face

import (
	_ "embed"
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

//go:embed clock.tmpl
var svgTemplate string
var tmpl = makeTemplate()

func SVGWriter(writer io.Writer, clock *Clock, center Point, radius float64) {
	data := makeSVGData(clock, center, radius)
	err := tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}

func makeTemplate() *template.Template {
	tmpl, err := template.New("svgClock").Parse(svgTemplate)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func makeSVGData(clock *Clock, center Point, radius float64) svgClockData {
	secondHandLength := radius * 0.6
	minuteHandLength := radius * 0.5
	hourHandLength := radius * 0.3

	return svgClockData{
		Center:     center,
		SecondHand: unitPointToHandPoint(secondHandLength, clock.SecondHand(), center),
		MinuteHand: unitPointToHandPoint(minuteHandLength, clock.MinuteHand(), center),
		HourHand:   unitPointToHandPoint(hourHandLength, clock.HourHand(), center),
		Radius:     radius,
		Width:      center.X * 2,
		Height:     center.Y * 2,
	}
}

func unitPointToHandPoint(handLength float64, unitPoint, center Point) Point {
	x := handLength*unitPoint.X + center.X
	y := -handLength*unitPoint.Y + center.Y

	return Point{X: x, Y: y}
}
