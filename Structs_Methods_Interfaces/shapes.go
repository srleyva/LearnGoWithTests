package Structs_Methods_Interfaces

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	base float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height*2) + (r.Width*2)
}

func (r Rectangle) Area() float64  {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64  {
	return (t.base * t.height) * 0.5
}