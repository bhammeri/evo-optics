package utils

type Vector struct {
	Start Point
	End   Point
}

type DirectionVector struct {
	Start   Point
	End     Point
	LengthX float64
	LengthY float64
	Angle   Radian
}
