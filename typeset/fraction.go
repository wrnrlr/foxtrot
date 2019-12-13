package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"image"
)

type Fraction struct {
	Numerator, Denominator Shape
}

func (f *Fraction) Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions {
	dN := f.Numerator.Dimensions(c, s, font)
	dD := f.Denominator.Dimensions(c, s, font)
	width := max(dN.Size.X, dD.Size.X)
	height := dN.Size.Y + dD.Size.Y + int(unit.Sp(4).V)
	dims := layout.Dimensions{
		Size:     image.Point{X: width, Y: height},
		Baseline: height / 2}
	return dims
}

func (f *Fraction) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	dims := f.Dimensions(gtx, s, font)
	var stack op.StackOp

	stack.Push(gtx.Ops)
	dN := f.Numerator.Dimensions(gtx, s, font)
	leftOffset := float32(dims.Size.X-dN.Size.X) / 2
	offset := f32.Point{X: leftOffset, Y: 0}
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	f.Numerator.Layout(gtx, s, font)
	stack.Pop()

	topOffset := float32(dN.Size.Y + gtx.Px(unit.Sp(1)))
	height := gtx.Px(unit.Sp(1))
	w := float32(dims.Size.X)
	h := float32(height)
	stack.Push(gtx.Ops)
	offset = f32.Point{X: 0, Y: topOffset}
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	size := float32(gtx.Px(unit.Sp(1)))
	var p clip.Path
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 0, Y: 0})
	p.Line(f32.Point{X: w, Y: 0})
	p.Line(f32.Point{X: 0, Y: size})
	p.Line(f32.Point{X: -w, Y: 0})
	p.Line(f32.Point{X: 0, Y: -size})
	p.End()
	paint.PaintOp{f32.Rectangle{Max: f32.Point{X: w, Y: h}}}.Add(gtx.Ops)
	stack.Pop()

	topOffset += float32(gtx.Px(unit.Sp(1)))
	stack.Push(gtx.Ops)
	dD := f.Denominator.Dimensions(gtx, s, font)
	leftOffset = float32(dims.Size.X-dD.Size.X) / 2
	offset = f32.Point{X: leftOffset, Y: topOffset}
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	f.Denominator.Layout(gtx, s, font)
	stack.Pop()

	gtx.Dimensions = dims
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
