package clock_face_test

import (
	"bytes"
	"encoding/xml"
	"hello/clock_face"
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
	Circle  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line []struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}

func TestSVGWriteerAtMidnight(t *testing.T) {
	time := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	c := clock_face.Clock{time, clock_face.Point{150, 150}, 150}
	writer := bytes.Buffer{}
	clock_face.SVGWriter(&writer, &c)
	svg := SVG{}

	xml.Unmarshal(writer.Bytes(), &svg)

	wantX := "150.000"
	wantY := "60.000"

	for _, line := range svg.Line {
		if line.ID == "second_hand" {
			if line.X2 == "150.000" && line.Y2 == "60.000" {
				return
			} else {
				t.Fatalf("Wanted point {%s, %s} but got {%s, %s}", wantX, wantY, line.X2, line.Y2)
			}
		}
	}

}
