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

const templateFile = "./clock.tmpl"

func SVGWriter(writer io.Writer, clock *Clock, center Point, radius float64) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	data := makeSVGData(clock, center, radius)
	err = tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
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
	}
}

func unitPointToHandPoint(handLength float64, unitPoint, center Point) Point {
	x := handLength*unitPoint.X + center.X
	y := -handLength*unitPoint.Y + center.Y

	return Point{X: x, Y: y}
}
