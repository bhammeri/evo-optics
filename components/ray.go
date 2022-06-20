package components

import (
	"github.com/fogleman/gg"
	"math"
)

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
	StartPoint  Point
	Direction   DirectionVector
	Terminating bool
}

type Ray struct {
	Segments []RaySegment
}

func (ray *Ray) AddSegment(startPoint Point, direction DirectionVector) {
	ray.Segments = append(ray.Segments, RaySegment{StartPoint: startPoint, Direction: direction})
}

func (ray *Ray) Draw(context *gg.Context) {
	numberOfSegments := len(ray.Segments)

	context.SetRGBA(0, 0, 0, 1)
	// last point is the final intersect
	var currentRaySegment, nextRaySegment RaySegment
	for i := 0; i < numberOfSegments-1; i++ {
		currentRaySegment = ray.Segments[i]
		nextRaySegment = ray.Segments[i+1]
		context.Push()
		context.DrawLine(
			currentRaySegment.StartPoint.X,
			currentRaySegment.StartPoint.Y,
			nextRaySegment.StartPoint.X,
			nextRaySegment.StartPoint.Y,
		)
		context.Fill()
		context.Pop()
	}
}
