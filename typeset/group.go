package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
)

type Group struct {
	Parts                  []Shape
	Subscript, SuperScript *Shape
}

func (g *Group) Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions {
	var dims layout.Dimensions
	for _, l := range g.lines(c, s, font) {
		d := l.dimensions(c, s, font)
		dims.Size.X = max(dims.Size.X, d.Size.X)
		dims.Size.Y += d.Size.Y
	}
	return dims
}

func (g *Group) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	var stack op.StackOp
	var dims layout.Dimensions
	lineOffset := 0
	for _, l := range g.lines(gtx, s, font) {
		ld := l.dimensions(gtx, s, font)
		for _, p := range l.shapes {
			stack.Push(gtx.Ops)
			d := p.Dimensions(gtx, s, font)
			offset := f32.Point{X: float32(dims.Size.X), Y: float32((ld.Size.Y-d.Size.Y)/2 + lineOffset)}
			op.TransformOp{}.Offset(offset).Add(gtx.Ops)
			p.Layout(gtx, s, font)
			dims.Size.X += d.Size.X
			stack.Pop()
		}
		dims.Size.X = 0
		lineOffset += ld.Size.Y
	}
	gtx.Dimensions = g.Dimensions(gtx, s, font)
}

func (g *Group) lines(gtx *layout.Context, s *text.Shaper, font text.Font) []line {
	maxWidth := gtx.Constraints.Width.Max
	lineWidth := 0
	lines := []line{}
	i := 0
	for _, p := range g.Parts {
		d := p.Dimensions(gtx, s, font)
		if lineWidth+d.Size.X > maxWidth {
			l := line{[]Shape{p}}
			lines = append(lines, l)
			lineWidth = d.Size.X
			i++
		} else {
			if len(lines) == 0 {
				lines = append(lines, line{})
			}
			shapes := lines[i].shapes
			if shapes == nil {
				shapes = []Shape{}
			}
			lines[i].shapes = append(shapes, p)
			lineWidth += d.Size.X
		}
	}
	return lines
}

type line struct {
	shapes []Shape
}

func (l line) dimensions(gtx *layout.Context, sh *text.Shaper, font text.Font) layout.Dimensions {
	dims := layout.Dimensions{}
	for _, s := range l.shapes {
		d := s.Dimensions(gtx, sh, font)
		dims.Size.X += d.Size.X
		dims.Size.Y = max(dims.Size.Y, d.Size.Y)
		dims.Baseline = max(dims.Baseline, d.Baseline)
	}
	return dims
}
