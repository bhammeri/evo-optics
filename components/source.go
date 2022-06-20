package components

type PointSource struct {
	Location Point
	Rays     []Ray
}

func (source *PointSource) InitRays(numberOfRays int, openingAngle Radiant) {
	angleIncrement := float64(openingAngle) / float64(numberOfRays-1)
	startAngle := float64(-1) * float64(openingAngle) / float64(2)
	var newRay Ray
	var angle Radiant
	for i := 0; i < numberOfRays; i++ {
		angle = Radiant(startAngle + angleIncrement*float64(i))
		newRay = Ray{}
		newRay.AddSegment(source.Location, NewDirectionVector(source.Location, angle))
		source.Rays = append(source.Rays, newRay)
	}
}
