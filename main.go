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
	source.InitRays(3, 1)

	for ray := range source.Rays {
		fmt.Printf("%+v", source.Rays[ray])
	}
}
