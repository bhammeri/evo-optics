package components

type PointSource struct {
	Location Point
	Rays     []Ray
}

func (source *PointSource) InitRays(numberOfRays int, openingAngle Radiant) {
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
	var angle Radiant
	for i := 0; i < numberOfRays; i++ {
		angle = Radiant(startAngle + angleIncrement*float64(i))
		newRay = Ray{}
		newRay.AddSegment(source.Location, NewDirectionVector(source.Location, angle))
		directionVector := NewDirectionVector(source.Location, angle)
		point := Point{
			X: directionVector.End.X * 20,
			Y: directionVector.End.Y * 20,
		}
		newRay.AddSegment(point, NewDirectionVector(source.Location, angle))
		source.Rays = append(source.Rays, newRay)
	}
}
