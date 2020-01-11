package shape

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
)

func StrokeRectangle(p1, p2 f32.Point, lineWidth float32, ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
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
	path.End().Add(ops)
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
