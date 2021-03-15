package StructsMethodsAndInterfaces

import "math"

type rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
type Triangle struct {
	base   float64
	height float64
}
type Shape interface {
	Area() float64
}

func Perimeter(rec rectangle) float64 {
	return 2 * (rec.width + rec.height)
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

func Area(rec rectangle) float64 {
	return rec.width * rec.height
}
