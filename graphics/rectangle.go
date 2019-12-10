package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/util"
)

func toRectangle(e *atoms.Expression) (*Rectangle, error) {
	r := Rectangle{}
	return &r, nil
}

type Rectangle struct {
	min, max f32.Point
}

func (r Rectangle) Draw(ctx *context, ops *op.Ops) {
	var p clip.Path
	p.Begin(ops)
	p.Move(f32.Point{X: 0, Y: 0})
	p.Line(f32.Point{X: 100, Y: 0})
	p.Line(f32.Point{X: 0, Y: 100})
	p.Line(f32.Point{X: -100, Y: 0})
	p.Line(f32.Point{X: 0, Y: -100})
	p.End().Add(ops)
	paint.ColorOp{util.Black}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(100), Y: 100}}}.Add(ops)
}

func (r Rectangle) BoundingBox() (bbox f32.Rectangle) {
	return bbox
}
