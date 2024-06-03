package clock_face

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	ID string  `xml:"id,attr"`
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	ID string  `xml:"id,attr"`
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

var clockCenter = Point{150, 150}
var clockRadius = 150.0

func TestSVGWriterSecondHand(t *testing.T) {

	secondHandLength := clockRadius * 0.6

	cases := []struct {
		name string
		time time.Time
		line Line
	}{
		{
			"0 seconds",
			simpleTime(0, 0, 0),
			Line{"second_hand", clockCenter.X, clockCenter.Y, clockCenter.X, clockCenter.Y - secondHandLength},
		},
		{
			"30 seconds",
			simpleTime(0, 0, 30),
			Line{"second_hand", clockCenter.X, clockCenter.Y, clockCenter.X, clockCenter.Y + secondHandLength},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Clock{testCase.time}
			buffer := bytes.Buffer{}
			want := testCase.line
			SVGWriter(&buffer, &c, clockCenter, clockRadius)
			svg := SVG{}
			xml.Unmarshal(buffer.Bytes(), &svg)
			if !containsLine(want, svg.Line) {
				t.Errorf("Expected the line %v to be present in the svg but was not. Lines: %v", want, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {

	minuteHandLength := clockRadius * 0.5

	cases := []struct {
		name string
		time time.Time
		line Line
	}{
		{
			"0 seconds",
			simpleTime(0, 0, 0),
			Line{"minute_hand", clockCenter.X, clockCenter.Y, clockCenter.X, clockCenter.Y - minuteHandLength},
		},
		{
			"30 seconds",
			simpleTime(0, 30, 0),
			Line{"minute_hand", clockCenter.X, clockCenter.Y, clockCenter.X, clockCenter.Y + minuteHandLength},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Clock{testCase.time}
			buffer := bytes.Buffer{}
			want := testCase.line
			SVGWriter(&buffer, &c, clockCenter, clockRadius)
			svg := SVG{}
			xml.Unmarshal(buffer.Bytes(), &svg)
			if !containsLine(want, svg.Line) {
				t.Errorf("Expected the line %v to be present in the svg but was not. Lines: %v", want, svg.Line)
			}
		})
	}
}

func containsLine(l Line, lines []Line) bool {
	for _, line := range lines {
		if line == l {
			return true
		}
	}
	return false
}
