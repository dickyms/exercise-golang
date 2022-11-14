package main

import (
    "testing"
)

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rec Rectangle) float64  {
	return 2 * (rec.Width + rec.Height)
}

func Area(rec Rectangle) float64  {
	return rec.Width * rec.Height
}

func  TestPerimeter(t *testing.T)  {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T)  {
	rec := Rectangle{12.0, 6.0}
	got := Area(rec)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}