package engine3D

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p *Point) Translate(vector Point) {
	p.X += vector.X
	p.Y += vector.Y
	p.Z += vector.Z
}

func (p Point) Rotate(angle float64) Point {
	return Point{
		X: p.X*math.Cos(angle) + p.Z*math.Sin(angle),
		Y: p.Y,
		Z: -p.X*math.Sin(angle) + p.Z*math.Cos(angle),
	}
}
