package utils

import "math"

type Vector struct {
	Start Point
	End   Point
}

func (vector *Vector) Magnitude() float64 {
	xSquared := math.Pow(vector.End.X-vector.Start.X, 2)
	ySquared := math.Pow(vector.End.Y-vector.Start.Y, 2)

	return math.Sqrt(xSquared + ySquared)
}

func (vector *Vector) AngleTo(otherVector *Vector) float64 {
	return CrossProduct(vector, otherVector)
}

type DirectionVector struct {
	Start   Point
	End     Point
	LengthX float64
	LengthY float64
	Angle   Radian
}

func CrossProduct(v1 *Vector, v2 *Vector) float64 {
	x1 := v1.End.X - v1.Start.X
	y1 := v1.End.Y - v1.Start.Y

	x2 := v2.End.X - v2.Start.X
	y2 := v2.End.Y - v2.Start.Y

	return (x1 * x2) + (y1 * y2)
}
