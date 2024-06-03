package main

import (
	"hello/clock_face"
	"os"
	"time"
)

func main() {
	clock := clock_face.Clock{Time: time.Now()}
	center := clock_face.Point{X: 150, Y: 150}
	radius := 100.0
	clock_face.SVGWriter(os.Stdout, &clock, center, radius)
}
