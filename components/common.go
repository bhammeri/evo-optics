package components

import "github.com/fogleman/gg"

func DrawCross(context *gg.Context, centerX, centerY, originX, originY float64) {
	dimension := 10.0

	var corners [4]Point
	corners[0] = Point{X: originX + centerX - 0.5*dimension, Y: originY + centerY - 0.5*dimension}
	corners[1] = Point{X: originX + centerX - 0.5*dimension, Y: originY + centerY + 0.5*dimension}
	corners[2] = Point{X: originX + centerX + 0.5*dimension, Y: originY + centerY + 0.5*dimension}
	corners[3] = Point{X: originX + centerX + 0.5*dimension, Y: originY + centerY - 0.5*dimension}

	context.SetRGBA(0.0, 0.0, 0.0, 1.0)

	context.Push()
	context.DrawLine(
		corners[0].X,
		corners[0].Y,
		corners[2].X,
		corners[2].Y,
	)
	context.DrawLine(
		corners[1].X,
		corners[1].Y,
		corners[3].X,
		corners[3].Y,
	)
	context.Stroke()
	context.Pop()

}
