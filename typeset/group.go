package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/wrnrlr/foxtrot/style"
)

type Group struct {
	Parts                  []Shape
	Subscript, SuperScript *Shape
}

func (g *Group) Dimensions(gtx *layout.Context, s style.Style) layout.Dimensions {
	var dims layout.Dimensions
	for _, l := range g.lines(gtx, s) {
		d := l.dimensions(gtx, s)
		dims.Size.X = max(dims.Size.X, d.Size.X)
		dims.Size.Y += d.Size.Y
	}
	return dims
}

func (g *Group) Layout(gtx *layout.Context, s style.Style) {
	var stack op.StackOp
	var dims layout.Dimensions
	lineOffset := 0
	for _, l := range g.lines(gtx, s) {
		ld := l.dimensions(gtx, s)
		for _, p := range l.shapes {
			stack.Push(gtx.Ops)
			d := p.Dimensions(gtx, s)
			offset := f32.Point{X: float32(dims.Size.X), Y: float32((ld.Size.Y-d.Size.Y)/2 + lineOffset)}
			op.TransformOp{}.Offset(offset).Add(gtx.Ops)
			p.Layout(gtx, s)
			dims.Size.X += d.Size.X
			stack.Pop()
		}
		dims.Size.X = 0
		lineOffset += ld.Size.Y
	}
	gtx.Dimensions = g.Dimensions(gtx, s)
}

func (g *Group) lines(gtx *layout.Context, s style.Style) []line {
	maxWidth := gtx.Constraints.Width.Max
	lineWidth := 0
	lines := []line{}
	i := 0
	for _, p := range g.Parts {
		d := p.Dimensions(gtx, s)
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

func (l line) dimensions(gtx *layout.Context, s style.Style) layout.Dimensions {
	dims := layout.Dimensions{}
	for _, shape := range l.shapes {
		d := shape.Dimensions(gtx, s)
		dims.Size.X += d.Size.X
		dims.Size.Y = max(dims.Size.Y, d.Size.Y)
		dims.Baseline = max(dims.Baseline, d.Baseline)
	}
	return dims
}
