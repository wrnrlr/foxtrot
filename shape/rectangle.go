package shape

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
)

func StrokeRectangle(p1, p2 f32.Point, lineWidth float32, ops *op.Ops) {
	scale := float32(0.80) // base value on lineWidth
	var path clip.Path
	path.Begin(ops)
	path.Move(p1)
	path.Line(f32.Point{X: p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: p2.Y})
	path.Line(f32.Point{X: -p2.X, Y: 0})
	path.Line(f32.Point{X: 0, Y: -p2.Y})
	path.Move(f32.Point{X: p2.X * (1 - scale) * 0.5, Y: p2.Y * (1 - scale) * 0.5})
	p2.X *= scale
	p2.Y *= scale
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
