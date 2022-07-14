package components

import (
	"evo-optics/utils"
	"github.com/fogleman/gg"
	"math"
)

type DirectionVector struct {
	Start   utils.Point
	End     utils.Point
	LengthX float64
	LengthY float64
	Angle   utils.Radian
}

func NewDirectionVector(point utils.Point, angle utils.Radian) DirectionVector {
	lengthX := 1.0
	lengthY := math.Tan(float64(angle))

	endPoint := utils.Point{
		X: point.X + lengthX,
		Y: point.Y + lengthY,
	}
	return DirectionVector{
		Start:   point,
		End:     endPoint,
		LengthX: lengthX,
		LengthY: lengthY,
		Angle:   angle,
	}
}

type RaySegment struct {
	StartPoint      utils.Point
	Direction       DirectionVector
	Terminating     bool
	RefractionIndex float64
}

type Ray struct {
	Segments   []RaySegment
	WaveLength float64 `unit:"nm"`
}

func (ray *Ray) AddSegment(startPoint utils.Point, direction DirectionVector, refractionIndex float64) {
	ray.Segments = append(
		ray.Segments,
		RaySegment{StartPoint: startPoint, Direction: direction, RefractionIndex: refractionIndex},
	)
}

func (ray *Ray) Draw(context *gg.Context, originX float64, originY float64) {
	numberOfSegments := len(ray.Segments)
	r, g, b := ray.RGB()
	context.SetRGBA(r, g, b, 1)
	// last point is the final intersect
	var currentRaySegment, nextRaySegment RaySegment
	for i := 0; i < numberOfSegments-1; i++ {
		currentRaySegment = ray.Segments[i]
		nextRaySegment = ray.Segments[i+1]
		context.Push()
		context.DrawLine(
			currentRaySegment.StartPoint.X+originX,
			currentRaySegment.StartPoint.Y+originY,
			nextRaySegment.StartPoint.X+originX,
			nextRaySegment.StartPoint.Y+originY,
		)
		context.Stroke()
		context.Pop()
	}
}

func (ray *Ray) RGB() (red, green, blue float64) {
	// https://stackoverflow.com/questions/1472514/convert-light-frequency-to-rgb
	waveLength := ray.WaveLength
	switch {
	case waveLength < 380:
		red = 0.0
		green = 0.0
		blue = 1.0
	case waveLength < 440:
		red = -(waveLength - 440) / (440 - 380)
		green = 0.0
		blue = 1.0
	case waveLength < 490:
		red = 0.0
		green = (waveLength - 440) / (490 - 440)
		blue = 1.0
	case waveLength < 510:
		red = 0.0
		green = 1.0
		blue = -(waveLength - 510) / (580 - 510)
	case waveLength < 580:
		red = (waveLength - 510) / (580 - 510)
		green = 1.0
		blue = 0.0
	case waveLength < 645:
		red = 1.0
		green = -(waveLength - 645) / (645 - 580)
		blue = 0.0
	case waveLength >= 645:
		red = 1.0
		green = 0.0
		blue = 0.0
	}

	var factor float64
	switch {
	case waveLength > 250 && waveLength < 420:
		factor = 0.3 + 0.7*(waveLength-250)/(420-250)
	case waveLength < 701:
		factor = 1.0
	case waveLength < 781:
		factor = 0.3 + 0.7*(780-waveLength)/(780-700)
	default:
		factor = 0.0
	}

	gamma := 0.8
	if red != 0.0 {
		red = math.Pow(red*factor, gamma)
	}
	if green != 0.0 {
		green = math.Pow(green*factor, gamma)
	}
	if blue != 0.0 {
		blue = math.Pow(blue*factor, gamma)
	}

	return red, green, blue
}
