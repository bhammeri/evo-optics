package main

import (
	"evo_optics/components"
	"fmt"
	"github.com/fogleman/gg"
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

	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(1, 1, 1, 1)
	dc.DrawRectangle(0, 0, S, S)
	dc.Fill()
	for _, ray := range source.Rays {
		ray.Draw(dc)
	}
	dc.SavePNG("out.png")
}
