package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
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

func (r Rectangle) Draw(ctx *context, ops *op.Ops) {
	x0 := float32(0)
	y0 := float32(0)
	x1 := r.max.X * 100
	y1 := r.max.Y * 100
	var p clip.Path
	p.Begin(ops)
	p.Move(f32.Point{X: x0, Y: y0})
	p.Line(f32.Point{X: x1, Y: 0})
	p.Line(f32.Point{X: 0, Y: y1})
	p.Line(f32.Point{X: -x1, Y: 0})
	p.Line(f32.Point{X: 0, Y: -y1})
	p.End().Add(ops)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(100), Y: 100}}}.Add(ops)
}

func (r Rectangle) BoundingBox() (bbox f32.Rectangle) {
	min := f32.Point{X: 0, Y: 0}
	max := f32.Point{X: 1, Y: 1}
	return f32.Rectangle{Min: min, Max: max}
}
