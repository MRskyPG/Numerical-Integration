package main

type Point struct {
	X, Y, Z float64
}

func NewPoint(x, y, z float64) *Point {
	return &Point{X: x, Y: y, Z: z}
}
