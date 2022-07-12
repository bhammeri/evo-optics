package components

import (
	"evo-optics/constants"
	"github.com/fogleman/gg"
	"math"
)

type Observation struct {
	X   float64
	Y   float64
	Ray *Ray
}

type Detector struct {
	X            float64
	Size         float64
	Observations []Observation
}

func (detector *Detector) InteractWithRay(ray *Ray) {
	lastRaySegment := ray.Segments[len(ray.Segments)-1]

	// find intersection point
	var y float64
	y = lastRaySegment.StartPoint.Y + (detector.X-lastRaySegment.StartPoint.X)/lastRaySegment.Direction.LengthX*lastRaySegment.Direction.LengthY
	ray.AddSegment(Point{detector.X, y}, lastRaySegment.Direction, constants.REFRACTION_INDEX_OF_VOID)

	if math.Abs(y) <= detector.Size/2 {
		detector.AddObservation(detector.X, y, ray)
	}
}

func (detector *Detector) AddObservation(x float64, y float64, ray *Ray) {
	detector.Observations = append(
		detector.Observations,
		Observation{X: x, Y: y, Ray: ray},
	)
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

	detector.drawObservations(context, originX, originY)
}

func (detector *Detector) drawObservations(context *gg.Context, originX float64, originY float64) {
	for _, observation := range detector.Observations {
		DrawCross(context, observation.X, observation.Y, originX, originY)
	}
}
