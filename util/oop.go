package util

import "math"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius float64
}

func (rec *Rectangle) Area() int {
	return rec.Width * rec.Height
}

func (rec *Rectangle) Perimeter() int {
	return 2 * (rec.Width + rec.Height)
}

func (cir *Circle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

func (cir *Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.Radius
}
