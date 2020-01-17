package shape

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

type Rectangle struct {
	P1, P2 f32.Point
}

func (r Rectangle) Fill(gtx *layout.Context) f32.Rectangle {
	p1, p2 := r.P1, r.P2
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.End().Add(gtx.Ops)
	box := f32.Rectangle{Min: p1, Max: p2}
	return box
}

func (r Rectangle) Stroke(lineWidth float32, gtx *layout.Context) f32.Rectangle {
	p1, p2 := r.P1, r.P2
	box := f32.Rectangle{Min: p1, Max: p2}
	var path clip.Path
	path.Begin(gtx.Ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.Move(f32.Point{X: lineWidth, Y: lineWidth})
	p2.X -= lineWidth * 2
	p2.Y -= lineWidth * 2
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.End().Add(gtx.Ops)
	return box
}

func StrokeRectangle(p1, p2 f32.Point, lineWidth float32, ops *op.Ops) {

}

func FillRectangle(p1, p2 f32.Point, ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.End().Add(ops)
}
