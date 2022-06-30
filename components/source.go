package components

import "evo-optics/constants"

type PointSource struct {
	Location Point
	Rays     []Ray
}

func (source *PointSource) InitRays(numberOfRays int, openingAngle Radian) {
	if numberOfRays < 1 {
		numberOfRays = 1
	}
	var angleIncrement, startAngle float64
	if numberOfRays == 1 {
		angleIncrement = 0
		startAngle = 0
	} else {
		angleIncrement = float64(openingAngle) / float64(numberOfRays-1)
		startAngle = float64(-1) * float64(openingAngle) / float64(2)
	}
	var newRay Ray
	var angle Radian
	for i := 0; i < numberOfRays; i++ {
		angle = Radian(startAngle + angleIncrement*float64(i))
		newRay = Ray{WaveLength: 320.0}
		newRay.AddSegment(
			source.Location,
			NewDirectionVector(source.Location, angle),
			constants.REFRACTION_INDEX_OF_VOID,
		)
		source.Rays = append(source.Rays, newRay)
	}
}
