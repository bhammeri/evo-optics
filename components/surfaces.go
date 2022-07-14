package components

import "evo-optics/utils"

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
