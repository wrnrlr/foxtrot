package graphics

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"golang.org/x/image/math/fixed"
	"image"
	"unicode/utf8"
)

type Tag struct {
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	MaxWidth int
}

type lineIterator struct {
	Lines     []text.Line
	Clip      image.Rectangle
	Alignment text.Alignment
	Width     int
	Offset    image.Point

	y, prevDesc fixed.Int26_6
}

const Inf = 1e6

func (l *lineIterator) Next() (text.String, f32.Point, bool) {
	for len(l.Lines) > 0 {
		line := l.Lines[0]
		l.Lines = l.Lines[1:]
		x := align(l.Alignment, line.Width, l.Width) + fixed.I(l.Offset.X)
		l.y += l.prevDesc + line.Ascent
		l.prevDesc = line.Descent
		// Align baseline and line start to the pixel grid.
		off := fixed.Point26_6{X: fixed.I(x.Floor()), Y: fixed.I(l.y.Ceil())}
		l.y = off.Y
		off.Y += fixed.I(l.Offset.Y)
		if (off.Y + line.Bounds.Min.Y).Floor() > l.Clip.Max.Y {
			break
		}
		if (off.Y + line.Bounds.Max.Y).Ceil() < l.Clip.Min.Y {
			continue
		}
		str := line.Text
		for len(str.Advances) > 0 {
			adv := str.Advances[0]
			if (off.X + adv + line.Bounds.Max.X - line.Width).Ceil() >= l.Clip.Min.X {
				break
			}
			off.X += adv
			_, s := utf8.DecodeRuneInString(str.String)
			str.String = str.String[s:]
			str.Advances = str.Advances[1:]
		}
		n := 0
		endx := off.X
		for i, adv := range str.Advances {
			if (endx + line.Bounds.Min.X).Floor() > l.Clip.Max.X {
				str.String = str.String[:n]
				str.Advances = str.Advances[:i]
				break
			}
			_, s := utf8.DecodeRuneInString(str.String[n:])
			n += s
			endx += adv
		}
		offf := f32.Point{X: float32(off.X) / 64, Y: float32(off.Y) / 64}
		return str, offf, true
	}
	return text.String{}, f32.Point{}, false
}

func (l Tag) Layout(gtx *layout.Context, st *Style, txt string) {
	paint.ColorOp{Color: st.TextColor}.Add(gtx.Ops)

	opts := text.LayoutOptions{MaxWidth: l.MaxWidth}
	textLayout := st.Shaper.Layout(gtx, st.Font, txt, opts)
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
		st.Shaper.Shape(gtx, st.Font, str).Add(gtx.Ops)
		paint.PaintOp{Rect: lclip}.Add(gtx.Ops)
		stack.Pop()
	}
	gtx.Dimensions = dims
}

func (e *Tag) layoutText(txt string, c unit.Converter, s *text.Shaper, font text.Font) ([]text.Line, layout.Dimensions) {
	opts := text.LayoutOptions{MaxWidth: e.MaxWidth}
	textLayout := s.Layout(c, font, txt, opts)
	lines := textLayout.Lines
	dims := linesDimens(lines)
	for i := 0; i < len(lines)-1; i++ {
		s := lines[i].Text.String
		// To avoid layout flickering while editing, assume a soft newline takes
		// up all available space.
		if len(s) > 0 {
			r, _ := utf8.DecodeLastRuneInString(s)
			if r != '\n' {
				dims.Size.X = e.MaxWidth
				break
			}
		}
	}
	return lines, dims
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
