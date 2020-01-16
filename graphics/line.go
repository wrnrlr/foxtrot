package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
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

func (l Line) Draw(ctx *context, gtx *layout.Context) {
	points := l.transformePoints(l.points, ctx)
	points = l.scalePoints(points, float32(gtx.Px(unit.Sp(100))))
	shape.StrokeLine(points, gtx.Px(unit.Sp(1)), gtx.Ops)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(gtx.Ops)
	//p := f32.Point{}
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx.Ops)
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

func (l Line) scalePoints(points []f32.Point, factor float32) []f32.Point {
	ps := make([]f32.Point, len(points))
	for i, p := range points {
		ps[i] = p.Mul(factor)
	}
	return ps
}

func (l Line) transformePoints(points []f32.Point, ctx *context) []f32.Point {
	ps := make([]f32.Point, len(points))
	for i, p := range points {
		ps[i] = ctx.transformPoint(p)
	}
	return ps
}
