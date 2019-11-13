package foxtrot

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"math/big"
)

func Rational2(num *big.Int, den *big.Int, gtx *layout.Context) {
	var stack op.StackOp
	stack.Push(gtx.Ops)
	txt := fmt.Sprintf("%s\n%s", num.String(), den.String())
	l1 := &Tag{Alignment: text.Middle, MaxWidth: inf}
	l1.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, txt)
	stack.Pop()

	dim1 := gtx.Dimensions

	labelHeight := dim1.Size.Y
	labelWidth := dim1.Size.X

	height := gtx.Config.Px(unit.Sp(2))
	w := float32(labelWidth)
	h := float32(height)

	stack.Push(gtx.Ops)
	offset := f32.Point{X: 0, Y: float32(labelHeight) / 2}
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	s := float32(gtx.Config.Px(unit.Sp(1)))
	var p clip.Path
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 0, Y: 0})
	p.Line(f32.Point{X: w, Y: 0})
	p.Line(f32.Point{X: 0, Y: s})
	p.Line(f32.Point{X: -w, Y: 0})
	p.Line(f32.Point{X: 0, Y: -s})
	p.End()
	paint.ColorOp{black}.Add(gtx.Ops)
	paint.PaintOp{f32.Rectangle{Max: f32.Point{X: w, Y: h}}}.Add(gtx.Ops)
	stack.Pop()
}
