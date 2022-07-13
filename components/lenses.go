package components

import (
	"github.com/fogleman/gg"
	"math"
)

func NewDoubleConvexLens(Center Point, Height, Width, Radius1, Radius2 float64) DoubleConvexLens {
	dcl := DoubleConvexLens{
		Center:  Center,
		Height:  Height,
		Width:   Width,
		Radius1: Radius1,
		Radius2: Radius2,
	}
	dcl.calculateDimensions()
	return dcl
}

type DoubleConvexLens struct {
	Center    Point
	Height    float64
	Width     float64
	Radius1   float64
	Center1   Point
	ArcAngle1 float64
	Radius2   float64
	Center2   Point
	ArcAngle2 float64
	Corners   [4]Point
}

func (lens *DoubleConvexLens) calculateDimensions() {
	lens.Corners[0] = Point{lens.Center.X - lens.Width/2, lens.Center.Y - lens.Height/2}
	lens.Corners[1] = Point{lens.Center.X - lens.Width/2, lens.Center.Y + lens.Height/2}
	lens.Corners[2] = Point{lens.Center.X + lens.Width/2, lens.Center.Y + lens.Height/2}
	lens.Corners[3] = Point{lens.Center.X + lens.Width/2, lens.Center.Y - lens.Height/2}

	lens.Center1 = Point{
		X: lens.Corners[1].X + math.Sqrt(math.Pow(lens.Radius1, 2)-math.Pow(lens.Height/2, 2)),
		Y: lens.Center.Y,
	}

	lens.Center2 = Point{
		X: lens.Corners[2].X - math.Sqrt(math.Pow(lens.Radius2, 2)-math.Pow(lens.Height/2, 2)),
		Y: lens.Center.Y,
	}

	lens.ArcAngle1 = math.Sin(lens.Height / 2 / lens.Radius1)
	lens.ArcAngle2 = math.Sin(lens.Height / 2 / lens.Radius2)
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
		lens.Center1.X+originX,
		lens.Center1.Y+originY,
		-lens.Radius1,
		-lens.ArcAngle1,
		lens.ArcAngle1,
	)
	context.Stroke()
	context.Pop()

	context.Push()
	context.DrawArc(
		lens.Center2.X+originX,
		lens.Center2.Y+originY,
		lens.Radius2,
		-lens.ArcAngle2,
		lens.ArcAngle2,
	)

	context.Stroke()
	context.Pop()
}
