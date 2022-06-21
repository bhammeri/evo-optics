package components

import (
	"github.com/fogleman/gg"
)

type Drawable interface {
	Draw(context *gg.Context, originX float64, originY float64)
}

type Canvas struct {
	Width           int
	Height          int
	padding         int
	availableWidth  int
	availableHeight int
	OriginX         int
	OriginY         int
	GGContext       *gg.Context
}

func NewCanvas(width int, height int, padding int) *Canvas {
	canvas := Canvas{
		Width:           width,
		Height:          height,
		padding:         padding,
		availableWidth:  width - (padding * 2),
		availableHeight: (height / 2) - (padding * 2),
		OriginX:         padding,
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
		float64(canvas.padding),
	)
	dc.Stroke()
	dc.Pop()
}

func (canvas *Canvas) SavePNG(filename string) {
	canvas.GGContext.SavePNG(filename)
}

func (canvas *Canvas) Draw(element Drawable) {
	element.Draw(canvas.GGContext, float64(canvas.OriginX), float64(canvas.OriginY))
}
