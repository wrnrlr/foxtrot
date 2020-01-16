package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/shape"
)

func toRectangle(e *atoms.Expression) (*Rectangle, error) {
	min := f32.Point{}
	var err error
	max := f32.Point{X: 1, Y: 1}
	length := e.Len()
	if length == 1 {
		max, err = toPoint(e.GetPart(1))
		if err != nil {
			return nil, err
		}
	} else if length == 2 {
		min, err = toPoint(e.GetPart(1))
		if err != nil {
			return nil, err
		}
		max, err = toPoint(e.GetPart(2))
		if err != nil {
			return nil, err
		}
	}
	r := Rectangle{min: min, max: max}
	return &r, nil
}

type Rectangle struct {
	min, max f32.Point
}

func (r Rectangle) Draw(ctx *context, gtx *layout.Context) {
	p1 := r.min.Mul(100)
	p2 := r.max.Mul(100)
	shape.StrokeRectangle(p1, p2, ctx.style.Thickness, gtx.Ops)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(100), Y: 100}}}.Add(gtx.Ops)
}

func (r Rectangle) BoundingBox() (bbox f32.Rectangle) {
	return f32.Rectangle{Min: r.min, Max: r.max}
}
