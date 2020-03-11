package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/wrnrlr/foxtrot/style"
)

type Word struct {
	Content                Shape
	Subscript, Superscript Shape
}

func (w *Word) Dimensions(gtx *layout.Context, s style.Style) layout.Dimensions {
	dims := w.Content.Dimensions(gtx, s)
	metrics := s.Shaper.Metrics(gtx, s.Font)
	xHeight := metrics.XHeight.Ceil()
	descent := metrics.Descent.Ceil()
	smallerFont := s
	smallerFont.Font = scaleDownFont(s.Font)
	if w.Subscript != nil {
		d := w.Subscript.Dimensions(gtx, smallerFont)
		dims.Size = dims.Size.Add(d.Size)
		dims.Size.Y -= xHeight
		dims.Baseline += d.Size.Y - xHeight
	}
	if w.Superscript != nil {
		d := w.Superscript.Dimensions(gtx, smallerFont)
		dims.Size = dims.Size.Add(d.Size)
		xHeight := metrics.XHeight.Ceil()
		dims.Size.Y -= xHeight
		dims.Baseline -= d.Size.Y - descent
	}
	return dims
}

func (w *Word) Layout(gtx *layout.Context, s style.Style) {
	var stack op.StackOp
	offset := f32.Point{X: 0, Y: 0}
	metrics := s.Shaper.Metrics(gtx, s.Font)
	xHeight := metrics.XHeight.Ceil()
	descent := metrics.Descent.Ceil()
	smallerFont := s
	smallerFont.Font = scaleDownFont(s.Font)
	contentDimesions := w.Content.Dimensions(gtx, s)
	if w.Superscript != nil {
		superscriptDims := w.Superscript.Dimensions(gtx, smallerFont)
		offset2 := f32.Point{X: float32(contentDimesions.Size.X), Y: 0}
		stack.Push(gtx.Ops)
		op.TransformOp{}.Offset(offset2).Add(gtx.Ops)
		w.Superscript.Layout(gtx, smallerFont)
		offset.Y = float32(superscriptDims.Size.Y - xHeight)
		stack.Pop()
	}
	stack.Push(gtx.Ops)
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	w.Content.Layout(gtx, s)
	offset.X = float32(gtx.Dimensions.Size.X)
	offset.Y += float32(w.Content.Dimensions(gtx, s).Size.Y - descent)
	stack.Pop()
	if w.Subscript != nil {
		stack.Push(gtx.Ops)
		op.TransformOp{}.Offset(offset).Add(gtx.Ops)
		w.Subscript.Layout(gtx, smallerFont)
		stack.Pop()
	}
	gtx.Dimensions = w.Dimensions(gtx, s)
}
