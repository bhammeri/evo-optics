package components

import (
	"evo-optics/constants"
	"github.com/fogleman/gg"
)

type Detector struct {
	X    float64
	Size float64
}

func (detector *Detector) InteractWithRay(ray *Ray) {
	lastRaySegment := ray.Segments[len(ray.Segments)-1]

	// find intersection point
	var y float64
	y = lastRaySegment.StartPoint.Y + (detector.X-lastRaySegment.StartPoint.X)/lastRaySegment.Direction.LengthX*lastRaySegment.Direction.LengthY
	ray.AddSegment(Point{detector.X, y}, lastRaySegment.Direction, constants.REFRACTION_INDEX_OF_VOID)
}

func (detector *Detector) Draw(context *gg.Context, originX float64, originY float64) {
	context.SetRGBA(0.0, 0.0, 0.0, 1.0)

	context.Push()
	context.DrawLine(
		detector.X+originX,
		originY-detector.Size/2,
		detector.X+originX,
		originY+detector.Size/2,
	)
	context.Stroke()
	context.Pop()
}
