package main

import (
	"evo-optics/components"
	"evo-optics/utils"
)

func main() {
	source := components.PointSource{
		Location: utils.Point{0, 0},
		Rays:     []components.Ray{},
	}
	source.InitRays(10, 1)

	detector := components.Detector{
		X:    500.0,
		Size: 250.0,
	}

	cuboid1 := components.Cuboid{
		Center:          utils.Point{X: 100, Y: 0},
		Width:           10,
		Height:          300,
		RefractionIndex: 10,
	}

	cuboid2 := components.Cuboid{
		Center:          utils.Point{X: 300, Y: 0},
		Width:           50,
		Height:          300,
		RefractionIndex: 10,
	}

	lens1 := components.NewDoubleConvexLens(
		utils.Point{X: 400, Y: 0},
		200,
		50,
		400,
		400,
	)

	canvas := components.NewCanvas(1024, 1024, 30)
	canvas.DrawBackground()
	canvas.DrawCoordinateSystem()
	for index := range source.Rays {
		cuboid1.InteractWithRay(&source.Rays[index])
		cuboid2.InteractWithRay(&source.Rays[index])
		lens1.InteractWithRay(&source.Rays[index], canvas.GGContext, float64(canvas.OriginX), float64(canvas.OriginY))
		detector.InteractWithRay(&source.Rays[index])
	}
	for _, ray := range source.Rays {
		canvas.Draw(&ray)
	}
	canvas.Draw(&detector)
	canvas.Draw(&cuboid1)
	canvas.Draw(&cuboid2)
	canvas.Draw(&lens1)

	canvas.SavePNG("out.png")

	detector.ScoreObservations()
}
