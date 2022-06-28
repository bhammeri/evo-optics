package main

import (
	"evo-optics/components"
	"fmt"
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

	canvas := components.NewCanvas(1024, 1024, 30)
	canvas.DrawBackground()
	canvas.DrawCoordinateSystem()
	for index := range source.Rays {
		detector.InteractWithRay(&source.Rays[index])
	}
	for ray := range source.Rays {
		fmt.Printf("ray: %+v \n\n", source.Rays[ray])
	}
	for _, ray := range source.Rays {
		canvas.Draw(&ray)
	}
	canvas.Draw(&detector)
	canvas.SavePNG("out.png")
}
