package utils

type Point struct {
	X float64
	Y float64
}

func (point *Point) Translate(deltaX, deltaY float64) {
	point.X += deltaX
	point.Y += deltaY
}

func (point *Point) EqualTo(otherPoint *Point) bool {
	return isClose(point.X, otherPoint.X) && isClose(point.Y, otherPoint.Y)
}

func DeltaBetweenPoints(a Point, b Point) (deltaX, deltaY float64) {
	deltaX = b.X - a.X
	deltaY = b.Y - a.Y

	return deltaX, deltaY
}
