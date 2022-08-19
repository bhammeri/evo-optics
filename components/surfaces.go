package components

import (
	"evo-optics/utils"
	"fmt"
	"math"
)

type SphericalSurface struct {
	Radius   float64
	Center   utils.Point
	ArcAngle float64
}

func (surface *SphericalSurface) NormalVector(pointOnSurface utils.Point) utils.Vector {
	unitVector := utils.NewUnitVector(surface.Center, pointOnSurface)
	unitVector.TranslateTo(pointOnSurface)
	return unitVector
}

func (surface *SphericalSurface) RayIntersect(ray *RaySegment) *utils.Point {
	cX := ray.StartPoint.X - surface.Center.X
	cY := ray.StartPoint.Y - surface.Center.Y

	fmt.Println(ray)

	a := math.Pow(ray.Direction.LengthX, 2) + math.Pow(ray.Direction.LengthY, 2)
	b := 2*ray.Direction.LengthX*cX + 2*ray.Direction.LengthY*cY
	c := math.Pow(cX, 2) + math.Pow(cY, 2) - math.Pow(surface.Radius, 2)

	if math.Pow(b, 2) <= 4*a*c {
		return nil
	}

	lambda1 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	lambda2 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)

	if lambda1 == lambda2 {
		fmt.Println("Lambda1 == Lambda2")
	}

	lambda := lambda1
	if lambda1 > lambda2 {
		lambda = lambda2
	}

	result := utils.Point{
		X: ray.StartPoint.X + lambda*ray.Direction.LengthX,
		Y: ray.StartPoint.Y + lambda*ray.Direction.LengthY,
	}
	fmt.Println(lambda1, lambda2, result)
	return &result
}
