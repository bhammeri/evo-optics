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
	source.InitRays(1, 1)

	for ray := range source.Rays {
		fmt.Printf("%+v", source.Rays[ray])
	}

	canvas := components.NewCanvas(1024, 1024)
	canvas.DrawBackground()
	canvas.DrawCoordinateSystem()
	canvas.SavePNG("out.png")
}
