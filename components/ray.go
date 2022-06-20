package components

type Point struct {
	X float64
	Y float64
}

type DirectionVector struct {
	Start Point
	End   Point
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
