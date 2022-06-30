package main

import (
	"evo-optics/components"
)

func main() {
	source := components.PointSource{
		Location: components.Point{0, 0},
		Rays:     []components.Ray{},
	}
	source.InitRays(10, 1)

	detector := components.Detector{
		X:    500.0,
		Size: 250.0,
	}

	cuboid1 := components.Cuboid{
		Center:          components.Point{X: 100, Y: 0},
		Width:           10,
		Height:          300,
		RefractionIndex: 10,
	}

	cuboid2 := components.Cuboid{
		Center:          components.Point{X: 300, Y: 0},
		Width:           50,
		Height:          300,
		RefractionIndex: 10,
	}

	canvas := components.NewCanvas(1024, 1024, 30)
	canvas.DrawBackground()
	canvas.DrawCoordinateSystem()
	for index := range source.Rays {
		cuboid1.InteractWithRay(&source.Rays[index])
		cuboid2.InteractWithRay(&source.Rays[index])
		detector.InteractWithRay(&source.Rays[index])
	}
	for _, ray := range source.Rays {
		canvas.Draw(&ray)
	}
	canvas.Draw(&detector)
	canvas.Draw(&cuboid1)
	canvas.Draw(&cuboid2)
	canvas.SavePNG("out.png")
}
