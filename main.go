package main

import (
	"evo_optics/components"
	"fmt"
)

func main() {
	source := components.PointSource{
		Location: components.Point{0, 0},
		Rays:     []components.Ray{},
	}
	source.InitRays(10, 1)

	for ray := range source.Rays {
		fmt.Printf("%+v", source.Rays[ray])
	}

	canvas := components.NewCanvas(1024, 1024, 30)
	canvas.DrawBackground()
	canvas.DrawCoordinateSystem()
	for _, ray := range source.Rays {
		canvas.Draw(&ray)
	}
	canvas.SavePNG("out.png")
}
