package components

import (
	"evo-optics/constants"
	"github.com/fogleman/gg"
	"math"
)

type Cuboid struct {
	Center          Point
	Height          float64
	Width           float64
	RefractionIndex float64
}

func (cuboid *Cuboid) InteractWithRay(ray *Ray) {
	cuboid.FrontInteractWithRay(ray)
	cuboid.BacksideInteractWithRay(ray)
}

func (cuboid *Cuboid) FrontInteractWithRay(ray *Ray) {
	bottomLeft := Point{
		X: cuboid.Center.X - (cuboid.Width / 2),
		Y: cuboid.Center.Y - (cuboid.Height / 2),
	}

	topLeft := Point{
		X: bottomLeft.X,
		Y: cuboid.Center.Y + (cuboid.Height / 2),
	}

	// find intersection
	var intersectY float64
	lastRaySegment := ray.Segments[len(ray.Segments)-1]
	intersectY = lastRaySegment.StartPoint.Y + (bottomLeft.X-lastRaySegment.StartPoint.X)/lastRaySegment.Direction.LengthX*lastRaySegment.Direction.LengthY

	if intersectY >= bottomLeft.Y && intersectY <= topLeft.Y {
		// calculate new direction
		newDirectionAngle := Radian(
			math.Asin(
				math.Sin(float64(lastRaySegment.Direction.Angle)) * lastRaySegment.RefractionIndex / cuboid.RefractionIndex,
			),
		)
		newStartPoint := Point{bottomLeft.X, intersectY}
		newDirection := NewDirectionVector(newStartPoint, newDirectionAngle)
		ray.AddSegment(newStartPoint, newDirection, cuboid.RefractionIndex)
	}
}

func (cuboid *Cuboid) BacksideInteractWithRay(ray *Ray) {
	bottomRight := Point{
		X: cuboid.Center.X + (cuboid.Width / 2),
		Y: cuboid.Center.Y - (cuboid.Height / 2),
	}

	topRight := Point{
		X: bottomRight.X,
		Y: cuboid.Center.Y + (cuboid.Height / 2),
	}

	// find intersection
	var intersectY float64
	lastRaySegment := ray.Segments[len(ray.Segments)-1]
	intersectY = lastRaySegment.StartPoint.Y + (bottomRight.X-lastRaySegment.StartPoint.X)/lastRaySegment.Direction.LengthX*lastRaySegment.Direction.LengthY
	if intersectY >= bottomRight.Y && intersectY <= topRight.Y {
		// calculate new direction
		newDirectionAngle := Radian(math.Asin(float64(lastRaySegment.Direction.Angle) * cuboid.RefractionIndex / constants.REFRACTION_INDEX_OF_VOID))
		newStartPoint := Point{bottomRight.X, intersectY}
		newDirection := NewDirectionVector(newStartPoint, newDirectionAngle)
		ray.AddSegment(newStartPoint, newDirection, constants.REFRACTION_INDEX_OF_VOID)
	}
}

func (cuboid *Cuboid) Draw(context *gg.Context, originX float64, originY float64) {
	bottomLeft := Point{
		X: cuboid.Center.X + originX - (cuboid.Width / 2),
		Y: cuboid.Center.Y + originY - (cuboid.Height / 2),
	}
	topRight := Point{
		X: cuboid.Center.X + originX + (cuboid.Width / 2),
		Y: cuboid.Center.Y + originY + (cuboid.Height / 2),
	}

	context.Push()
	context.DrawLine(
		bottomLeft.X,
		bottomLeft.Y,
		bottomLeft.X,
		topRight.Y,
	)
	context.DrawLine(
		bottomLeft.X,
		topRight.Y,
		topRight.X,
		topRight.Y,
	)
	context.DrawLine(
		topRight.X,
		topRight.Y,
		topRight.X,
		bottomLeft.Y,
	)
	context.DrawLine(
		topRight.X,
		bottomLeft.Y,
		bottomLeft.X,
		bottomLeft.Y,
	)
	context.Stroke()
	context.Pop()
}
