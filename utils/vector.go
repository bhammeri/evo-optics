package utils

import (
	"fmt"
	"math"
)

func NewUnitVector(point1 Point, point2 Point) Vector {
	// create a "non-localized" unit-vector from two points
	// "non-localized" means that the start point is (0,0)
	vector := Vector{
		Origin: point1,
		End:    point2,
	}
	vector.ScaleToUnitVector()

	x, y := DeltaBetweenPoints(vector.End, vector.Origin)

	return Vector{Origin: Point{0, 0}, End: Point{x, y}}
}

type Vector struct {
	Origin Point
	End    Point
}

func (vector *Vector) Magnitude() float64 {
	xSquared := math.Pow(vector.End.X-vector.Origin.X, 2)
	ySquared := math.Pow(vector.End.Y-vector.Origin.Y, 2)

	return math.Sqrt(xSquared + ySquared)
}

func (vector *Vector) EqualTo(otherVector *Vector) bool {
	return vector.Origin.EqualTo(&otherVector.Origin) && vector.End.EqualTo(&otherVector.End)
}

func (vector *Vector) ScaleToUnitVector() {
	magnitude := vector.Magnitude()

	deltaX, deltaY := DeltaBetweenPoints(vector.Origin, vector.End)

	fmt.Println(deltaX, deltaY, magnitude)

	vector.End = Point{
		X: vector.Origin.X + deltaX/magnitude,
		Y: vector.Origin.Y + deltaY/magnitude,
	}
}

func (vector *Vector) AngleTo(otherVector *Vector) float64 {
	return CrossProduct(vector, otherVector)
}

func (vector *Vector) TranslateTo(NewOriginPoint Point) {
	vector.Origin = NewOriginPoint

	deltaX, DeltaY := DeltaBetweenPoints(vector.End, vector.Origin)

	newEndPoint := NewOriginPoint
	newEndPoint.Translate(deltaX, DeltaY)
	vector.End = newEndPoint
}

func CrossProduct(v1 *Vector, v2 *Vector) float64 {
	x1, y1 := DeltaBetweenPoints(v1.End, v1.Origin)
	x2, y2 := DeltaBetweenPoints(v2.End, v2.Origin)

	return (x1 * x2) + (y1 * y2)
}
