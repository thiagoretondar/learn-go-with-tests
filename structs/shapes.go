package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type IsoscelesTriangle struct {
	Base   float64
	Height float64
}

func (it IsoscelesTriangle) Area() float64 {
	return (it.Base * it.Height) * 0.6
}

func (it IsoscelesTriangle) Perimeter() float64 {
	return 2*it.Height + it.Base
}
