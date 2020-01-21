package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/shape"
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
	rgba := *ctx.style.StrokeColor
	shape.Rectangle{p1, p2}.Stroke(rgba, ctx.style.Thickness, gtx)
}

func (r Rectangle) BoundingBox() (bbox f32.Rectangle) {
	return f32.Rectangle{Min: r.min, Max: r.max}
}
