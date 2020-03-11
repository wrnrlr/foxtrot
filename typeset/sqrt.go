package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/style"
)

func Sqrt(body Shape) Shape {
	return &sqrt{Content: body}
}

type sqrt struct {
	Content Shape
}

func (o *sqrt) Dimensions(gtx *layout.Context, s style.Style) layout.Dimensions {
	dims := o.Content.Dimensions(gtx, s)
	d := SqrtSymbol.Dimensions(gtx, s)
	dims.Size.X += d.Size.X
	return dims
}

func (o *sqrt) Layout(gtx *layout.Context, s style.Style) {
	var stack op.StackOp
	offset := f32.Point{X: 0, Y: 0}
	stack.Push(gtx.Ops)
	SqrtSymbol.Layout(gtx, s)
	offset.X += float32(gtx.Dimensions.Size.X)
	offset.Y = 0
	signWidth := float32(gtx.Dimensions.Size.X)
	stack.Pop()

	stack.Push(gtx.Ops)
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	o.Content.Layout(gtx, s)
	offset.X = signWidth
	contentWidth := float32(gtx.Dimensions.Size.X)
	offset.Y = 0
	stack.Pop()

	width := float32(gtx.Px(unit.Sp(1)))
	var p clip.Path
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	stack.Push(gtx.Ops)
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 0, Y: 0})
	p.Line(f32.Point{X: contentWidth, Y: 0})
	p.Line(f32.Point{X: 0, Y: width})
	p.Line(f32.Point{X: -contentWidth, Y: 0})
	p.Line(f32.Point{X: 0, Y: -width})
	p.End().Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{X: lineLen, Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
}
