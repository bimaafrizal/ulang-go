package main

import "math"

type Shape interface {
	area() float64
}

type Triangle struct {
	alas   int
	tinggi int
}

func (t Triangle) area() float64 {
	return 0.5 * float64(t.tinggi) * float64(t.alas)
}

type Circle struct {
	radius int
}

func (c Circle) area() float64 {
	radius := float64(c.radius * c.radius)
	return math.Pi * radius
}
