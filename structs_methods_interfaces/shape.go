package shape

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
