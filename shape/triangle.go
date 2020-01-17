package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

type Triangle struct {
	A, B, C f32.Point
}

func (t Triangle) Fill(width float32, gtx *layout.Context) (bbox f32.Rectangle) {
	return bbox
}

func StrokeTriangle(p1, p2, p3 f32.Point, lineWidth float32, ops *op.Ops) {
	p2 = p2.Sub(p1)
	p3 = p3.Sub(p2)
	var path clip.Path
	path.Begin(ops)
	path.Move(p1)
	path.Line(p2)
	path.Line(p3)
	path.End().Add(ops)
}

func FillTriangle(p1, p2, p3 f32.Point, ops *op.Ops) {
	pp2 := p2.Sub(p1)
	pp3 := p3.Sub(p2)
	pp1 := p1.Sub(p3)
	var path clip.Path
	path.Begin(ops)
	path.Move(p1)
	path.Line(pp2)
	path.Line(pp3)
	path.Line(pp1)
	path.End().Add(ops)
}
