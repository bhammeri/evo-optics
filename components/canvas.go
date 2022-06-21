package components

import (
	"github.com/fogleman/gg"
)

type Canvas struct {
	Width           int
	Height          int
	availableWidth  int
	availableHeight int
	OriginX         int
	OriginY         int
	GGContext       *gg.Context
}

func NewCanvas(width int, height int) *Canvas {
	canvas := Canvas{
		Width:           width,
		Height:          height,
		availableWidth:  width - 60,
		availableHeight: (height / 2) - 60,
		OriginX:         30,
		OriginY:         height / 2,
		GGContext:       gg.NewContext(width, height),
	}
	return &canvas
}

func (canvas *Canvas) DrawBackground() {
	dc := canvas.GGContext
	dc.Push()
	dc.SetRGBA(1, 1, 1, 1)
	dc.DrawRectangle(0, 0, float64(canvas.Width), float64(canvas.Height))
	dc.Fill()
	dc.Pop()
}

func (canvas *Canvas) DrawCoordinateSystem() {
	dc := canvas.GGContext
	dc.Push()
	dc.SetRGB255(125, 125, 125)
	dc.DrawLine(
		float64(canvas.OriginX),
		float64(canvas.OriginY),
		float64(canvas.OriginX+canvas.availableWidth),
		float64(canvas.OriginY),
	)
	dc.DrawLine(
		float64(canvas.OriginX),
		float64(canvas.OriginY),
		float64(canvas.OriginX),
		float64(canvas.OriginY+canvas.availableHeight),
	)
	dc.Stroke()
	dc.Pop()
}

func (canvas *Canvas) SavePNG(filename string) {
	canvas.GGContext.SavePNG(filename)
}
