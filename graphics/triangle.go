package graphics

import (
	"errors"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/shape"
)

func toTriangle(e *atoms.Expression) (*Triangle, error) {
	if e.Len() != 3 {
		return nil, errors.New("not valid Triangle[p1,p2,p3]")
	}
	p1, err := toPoint(e.GetPart(1))
	if err != nil {
		return nil, err
	}
	p2, err := toPoint(e.GetPart(2))
	if err != nil {
		return nil, err
	}
	p3, err := toPoint(e.GetPart(3))
	if err != nil {
		return nil, err
	}
	t := Triangle{p1: p1, p2: p2, p3: p3}
	return &t, nil
}

type Triangle struct {
	p1, p2, p3 f32.Point
}

func (t Triangle) Draw(ctx *context, gtx *layout.Context) {
	p1 := t.p1.Mul(float32(gtx.Px(unit.Sp(100))))
	p2 := t.p2.Mul(float32(gtx.Px(unit.Sp(100))))
	p3 := t.p3.Mul(float32(gtx.Px(unit.Sp(100))))
	shape.StrokeTriangle(p1, p2, p3, float32(gtx.Px(unit.Sp(5))), gtx.Ops)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(100), Y: 100}}}.Add(gtx.Ops)
}

func (r Triangle) BoundingBox() (bbox f32.Rectangle) {
	min := f32.Point{X: 0, Y: 0}
	max := f32.Point{X: 1, Y: 1}
	return f32.Rectangle{Min: min, Max: max}
}
