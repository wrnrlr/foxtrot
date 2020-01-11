package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/shape"
)

type Line struct {
	points []f32.Point
}

func toLine(e *atoms.Expression) (*Line, error) {
	points, err := toPoints(e.GetPart(1))
	if err != nil {
		return nil, err
	}
	line := Line{points: points}
	return &line, nil
}

func (l Line) Draw(ctx *context, ops *op.Ops) {
	points := l.transformedPoints(ctx)
	shape.StrokeLine(points, ops)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
}

func (l Line) BoundingBox() (bb f32.Rectangle) {
	for _, p := range l.points {
		if bb.Min.X > p.X {
			bb.Min.X = p.X
		}
		if bb.Min.Y > p.Y {
			bb.Min.Y = p.Y
		}
		if bb.Max.X < p.X {
			bb.Max.X = p.X
		}
		if bb.Max.Y < p.Y {
			bb.Max.Y = p.Y
		}
	}
	return bb
}

func (l Line) transformedPoints(ctx *context) []f32.Point {
	points := make([]f32.Point, len(l.points))
	for i, p := range l.points {
		points[i] = scale(ctx.transform(p))
	}
	return points
}

func scale(p f32.Point) f32.Point {
	return f32.Point{X: p.X * 100, Y: p.Y * 100}
}
