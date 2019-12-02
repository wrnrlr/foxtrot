package typeset

import (
	"fmt"
	"image"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"

	"golang.org/x/image/math/fixed"
)

const FitContent = 1e6

// Label is a widget for laying out and drawing text.
type Label struct {
	Text string
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxWith limits the with of a label, FitContent to limits the width to fit the content.
	MaxWidth int
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
}

func (l Label) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	//cs := gtx.Constraints
	options := text.LayoutOptions{MaxWidth: l.MaxWidth}
	textLayout := s.Layout(gtx, font, l.Text, options)
	//lines := textLayout.Lines
	//if max := l.MaxLines; max > 0 && len(lines) > max {
	//	lines = lines[:max]
	//}
	dims := linesDimens(textLayout.Lines)
	//dims.Size = cs.Constrain(dims.Size)
	clip := textPadding(textLayout.Lines)
	clip.Max = clip.Max.Add(dims.Size)
	it := lineIterator{
		Lines:     textLayout.Lines,
		Clip:      clip,
		Alignment: l.Alignment,
		Width:     dims.Size.X,
	}
	for {
		str, off, ok := it.Next()
		if !ok {
			break
		}
		lclip := toRectF(clip).Sub(off)
		var stack op.StackOp
		stack.Push(gtx.Ops)
		op.TransformOp{}.Offset(off).Add(gtx.Ops)
		s.Shape(gtx, font, str).Add(gtx.Ops)
		paint.PaintOp{Rect: lclip}.Add(gtx.Ops)
		stack.Pop()
	}
	gtx.Dimensions = dims
}

func (l *Label) Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions {
	options := text.LayoutOptions{MaxWidth: l.MaxWidth}
	textLayout := s.Layout(c, font, l.Text, options)
	lines := textLayout.Lines
	if max := l.MaxLines; max > 0 && len(lines) > max {
		lines = lines[:max]
	}
	dims := linesDimens(lines)
	return dims
}

func toRectF(r image.Rectangle) f32.Rectangle {
	return f32.Rectangle{
		Min: f32.Point{X: float32(r.Min.X), Y: float32(r.Min.Y)},
		Max: f32.Point{X: float32(r.Max.X), Y: float32(r.Max.Y)},
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

func align(align text.Alignment, width fixed.Int26_6, maxWidth int) fixed.Int26_6 {
	mw := fixed.I(maxWidth)
	switch align {
	case text.Middle:
		return fixed.I(((mw - width) / 2).Floor())
	case text.End:
		return fixed.I((mw - width).Floor())
	case text.Start:
		return 0
	default:
		panic(fmt.Errorf("unknown alignment %v", align))
	}
}
