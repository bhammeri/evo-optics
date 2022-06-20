package components

import "math"

type Radiant float64

type Point struct {
	X float64
	Y float64
}

type DirectionVector struct {
	Start Point
	End   Point
	Angle Radiant
}

func NewDirectionVector(point Point, angle Radiant) DirectionVector {
	endPoint := Point{
		X: point.X + 1,
		Y: point.Y + math.Tan(float64(angle)),
	}
	return DirectionVector{
		Start: point,
		End:   endPoint,
		Angle: angle,
	}
}

type RaySegment struct {
	StartPoint Point
	Direction  DirectionVector
}

type Ray struct {
	Segments []RaySegment
}

func (ray *Ray) AddSegment(startPoint Point, direction DirectionVector) {
	ray.Segments = append(ray.Segments, RaySegment{StartPoint: startPoint, Direction: direction})
}
