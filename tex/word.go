package tex

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
)

type Word struct {
	Content                Shape
	Subscript, Superscript Shape
}

func (w *Word) Dimensions(c unit.Converter, s *text.Shaper, font text.Font) layout.Dimensions {
	dims := w.Content.Dimensions(c, s, font)
	scaledFont := scaleDownFont(font)
	if w.Subscript != nil {
		d := w.Subscript.Dimensions(c, s, scaledFont)
		dims.Size = dims.Size.Add(d.Size)
		dims.Baseline += d.Size.Y
	}
	if w.Superscript != nil {
		d := w.Superscript.Dimensions(c, s, scaledFont)
		dims.Size = dims.Size.Add(d.Size)
		dims.Baseline -= d.Size.Y
	}
	return dims
}

func (w *Word) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	var stack op.StackOp
	offset := f32.Point{X: 0, Y: 0}
	scaledFont := scaleDownFont(font)
	contentDimesions := w.Content.Dimensions(gtx, s, font)
	if w.Superscript != nil {
		superscriptDims := w.Superscript.Dimensions(gtx, s, scaledFont)
		offset2 := f32.Point{X: float32(contentDimesions.Size.X), Y: 0}
		stack.Push(gtx.Ops)
		op.TransformOp{}.Offset(offset2).Add(gtx.Ops)
		w.Superscript.Layout(gtx, s, scaledFont)
		offset.Y = float32(superscriptDims.Size.Y)
		stack.Pop()
	}
	stack.Push(gtx.Ops)
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	w.Content.Layout(gtx, s, font)
	offset.X = float32(gtx.Dimensions.Size.X)
	offset.Y += float32(w.Content.Dimensions(gtx, s, font).Size.Y)
	stack.Pop()
	if w.Subscript != nil {
		stack.Push(gtx.Ops)
		op.TransformOp{}.Offset(offset).Add(gtx.Ops)
		w.Subscript.Layout(gtx, s, scaledFont)
		offset.Y = float32(gtx.Dimensions.Size.Y)
		stack.Pop()
	}
	gtx.Dimensions = w.Dimensions(gtx, s, font)
}
