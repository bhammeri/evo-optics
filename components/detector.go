package components

import (
	"evo-optics/constants"
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

type Observation struct {
	X   float64
	Y   float64
	Ray *Ray
}

type Observations []Observation

func (observations Observations) CenterOfMass() float64 {
	// center of mass
	totalNumberObservations := len(observations)
	weightedSum := 0.0

	for _, observation := range observations {
		weightedSum += observation.Y
	}

	return weightedSum / float64(totalNumberObservations)
}

func (observations Observations) AbsDistance() float64 {
	// absolute distance
	totalDistance := 0.0

	for _, observation := range observations {
		totalDistance += math.Abs(observation.Y)
	}

	return totalDistance
}

func (observations Observations) RMSDistance() float64 {
	// root mean square distance
	totalDistance := 0.0

	for _, observation := range observations {
		totalDistance += math.Pow(observation.Y, 2)
	}

	return math.Sqrt(totalDistance / float64(len(observations)))
}

type Detector struct {
	X            float64
	Size         float64
	Observations Observations
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

func (detector *Detector) ScoreObservations() {
	cm := detector.Observations.CenterOfMass()
	absDistance := detector.Observations.AbsDistance()
	rmsDistance := detector.Observations.RMSDistance()

	fmt.Printf("Center of Mass: %.2f \n", cm)
	fmt.Printf("Absolute Distance: %.2f \n", absDistance)
	fmt.Printf("RMS Distance: %.2f \n", rmsDistance)
}
