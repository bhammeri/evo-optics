package components

import (
	"evo-optics/utils"
	"github.com/fogleman/gg"
	"math"
)

func NewDoubleConvexLens(Center utils.Point, Height, Width, Radius1, Radius2 float64) DoubleConvexLens {
	dcl := DoubleConvexLens{
		Center:       Center,
		Height:       Height,
		Width:        Width,
		FrontSurface: SphericalSurface{Radius: Radius1},
		BackSurface:  SphericalSurface{Radius: Radius2},
	}
	dcl.calculateDimensions()
	return dcl
}

type SphericalSurface struct {
	Radius   float64
	Center   utils.Point
	ArcAngle float64
}

type DoubleConvexLens struct {
	Center       utils.Point
	Height       float64
	Width        float64
	FrontSurface SphericalSurface
	BackSurface  SphericalSurface
	Corners      [4]utils.Point
}

func (lens *DoubleConvexLens) calculateDimensions() {
	lens.Corners[0] = utils.Point{lens.Center.X - lens.Width/2, lens.Center.Y - lens.Height/2}
	lens.Corners[1] = utils.Point{lens.Center.X - lens.Width/2, lens.Center.Y + lens.Height/2}
	lens.Corners[2] = utils.Point{lens.Center.X + lens.Width/2, lens.Center.Y + lens.Height/2}
	lens.Corners[3] = utils.Point{lens.Center.X + lens.Width/2, lens.Center.Y - lens.Height/2}

	lens.FrontSurface.Center = utils.Point{
		X: lens.Corners[1].X + math.Sqrt(math.Pow(lens.FrontSurface.Radius, 2)-math.Pow(lens.Height/2, 2)),
		Y: lens.Center.Y,
	}

	lens.BackSurface.Center = utils.Point{
		X: lens.Corners[2].X - math.Sqrt(math.Pow(lens.BackSurface.Radius, 2)-math.Pow(lens.Height/2, 2)),
		Y: lens.Center.Y,
	}

	lens.FrontSurface.ArcAngle = math.Sin(lens.Height / 2 / lens.FrontSurface.Radius)
	lens.BackSurface.ArcAngle = math.Sin(lens.Height / 2 / lens.BackSurface.Radius)
}

func (lens *DoubleConvexLens) Draw(context *gg.Context, originX float64, originY float64) {
	context.SetRGBA(0.0, 0.0, 0.0, 1.0)

	context.Push()
	context.DrawLine(
		lens.Corners[0].X+originX,
		lens.Corners[0].Y+originY,
		lens.Corners[3].X+originX,
		lens.Corners[3].Y+originY,
	)
	context.DrawLine(
		lens.Corners[1].X+originX,
		lens.Corners[1].Y+originY,
		lens.Corners[2].X+originX,
		lens.Corners[2].Y+originY,
	)
	context.Stroke()
	context.Pop()

	// if context is not pushed and popped around DrawArc it also draws a line from the last point it was before
	context.Push()
	context.DrawArc(
		lens.FrontSurface.Center.X+originX,
		lens.FrontSurface.Center.Y+originY,
		-lens.FrontSurface.Radius,
		-lens.FrontSurface.ArcAngle,
		lens.FrontSurface.ArcAngle,
	)
	context.Stroke()
	context.Pop()

	context.Push()
	context.DrawArc(
		lens.BackSurface.Center.X+originX,
		lens.BackSurface.Center.Y+originY,
		lens.BackSurface.Radius,
		-lens.BackSurface.ArcAngle,
		lens.BackSurface.ArcAngle,
	)

	context.Stroke()
	context.Pop()
}
