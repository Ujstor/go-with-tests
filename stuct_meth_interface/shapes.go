package stuctmethinterface

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimiter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}