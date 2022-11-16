package main

import (
    "testing"
	"math"
)

// Table Test Driven
// membuat struct anonim array untuk mengiterasi pengujian struct dan method yang sama

type Rectangle struct {
	Width float64
	Height float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func TestArea(t *testing.T)  {

	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 32.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}