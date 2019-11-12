package foxtrot

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"math/big"
)

type Rational struct {
	// Face defines the text style.
	Font text.Font
	// Color is the text color.
	Color color.RGBA

	shaper *text.Shaper
}

func (ra Rational) Layout(num *big.Int, den *big.Int, gtx *layout.Context) {
	//cs := gtx.Constraints
	//textLayout := ra.shaper.Layout(gtx, ra.Font, num.String(), text.LayoutOptions{MaxWidth: cs.Width.Max})
	//ra.shaper.
}

const inf = 1e6

func Rational2(num *big.Int, den *big.Int, gtx *layout.Context) {
	var stack op.StackOp
	stack.Push(gtx.Ops)
	l1 := theme.Label(_defaultFontSize, num.String())
	l1.Font.Variant = "Mono"
	gtx.Constraints.Width.Max = inf
	l1.Layout(gtx)
	stack.Pop()

	dim1 := gtx.Dimensions

	labelHeight := dim1.Size.Y
	labelWidth := dim1.Size.X

	stack.Push(gtx.Ops)
	offset := image.Point{X: 0, Y: labelHeight}
	op.TransformOp{}.Offset(toPointF(offset)).Add(gtx.Ops)

	l2 := theme.Label(_defaultFontSize, den.String())
	l2.Font.Variant = "Mono"
	l2.Layout(gtx)
	stack.Pop()

	dim2 := gtx.Dimensions

	if labelWidth < gtx.Dimensions.Size.Y {
		labelWidth = gtx.Dimensions.Size.Y
	}

	height := gtx.Config.Px(unit.Sp(1))
	w := float32(labelWidth)
	h := float32(height)

	stack.Push(gtx.Ops)
	offset = image.Point{X: 0, Y: labelHeight}
	op.TransformOp{}.Offset(toPointF(offset)).Add(gtx.Ops)
	s := float32(gtx.Config.Px(unit.Sp(1)))
	var p paint.Path
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

	totalHeight := dim1.Size.Y + dim2.Size.Y
	gtx.Dimensions.Size = image.Point{X: labelWidth, Y: totalHeight}
}

func linesDimens(lines []text.Line) layout.Dimensions {
	var width fixed.Int26_6
	var h int
	var baseline int
	if len(lines) > 0 {
		baseline = lines[0].Ascent.Ceil()
		var prevDesc fixed.Int26_6
		for _, l := range lines {
			h += (prevDesc + l.Ascent).Ceil()
			prevDesc = l.Descent
			if l.Width > width {
				width = l.Width
			}
		}
		h += lines[len(lines)-1].Descent.Ceil()
	}
	w := width.Ceil()
	return layout.Dimensions{
		Size: image.Point{
			X: w,
			Y: h,
		},
		Baseline: h - baseline,
	}
}

func textPadding(lines []text.Line) (padding image.Rectangle) {
	if len(lines) == 0 {
		return
	}
	first := lines[0]
	if d := first.Ascent + first.Bounds.Min.Y; d < 0 {
		padding.Min.Y = d.Ceil()
	}
	last := lines[len(lines)-1]
	if d := last.Bounds.Max.Y - last.Descent; d > 0 {
		padding.Max.Y = d.Ceil()
	}
	if d := first.Bounds.Min.X; d < 0 {
		padding.Min.X = d.Ceil()
	}
	if d := first.Bounds.Max.X - first.Width; d > 0 {
		padding.Max.X = d.Ceil()
	}
	return
}
