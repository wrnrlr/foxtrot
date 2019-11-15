package notebook

import (
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
)

type Margin struct {
	eventKey int
	click    gesture.Click
	scroller gesture.Scroll
}

func (m *Margin) Event(gtx *layout.Context) interface{} {
	for _, e := range m.click.Events(gtx) {
		switch e.Type {
		case gesture.TypeClick:
			return SelectCellEvent{}
		}
	}
	return nil
}

func (m *Margin) Layout(gtx *layout.Context, checked bool, widget layout.Widget) {
	dim := gtx.Dimensions

	marginWidth := gtx.Config.Px(unit.Sp(15))
	editorWidth := gtx.Constraints.Width.Max - marginWidth
	gtx.Constraints.Width.Max = editorWidth
	widget()
	editorHeight := gtx.Dimensions.Size.Y

	var stack op.StackOp
	stack.Push(gtx.Ops)
	gtx.Constraints = layout.RigidConstraints(image.Point{X: marginWidth, Y: editorHeight})
	offset := image.Point{X: editorWidth, Y: 0}
	op.TransformOp{}.Offset(toPointF(offset)).Add(gtx.Ops)
	m.layoutMargin(checked, gtx)
	pointer.RectAreaOp{Rect: image.Rectangle{Max: image.Point{X: marginWidth, Y: editorHeight}}}.Add(gtx.Ops)
	m.scroller.Add(gtx.Ops)
	m.click.Add(gtx.Ops)
	stack.Pop()

	dim.Size.Y = editorHeight
	gtx.Dimensions = dim
}

func (m *Margin) layoutMargin(checked bool, gtx *layout.Context) {
	cs := gtx.Constraints

	if checked {
		d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
		dr := f32.Rectangle{
			Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
		}
		paint.ColorOp{Color: selectedColor}.Add(gtx.Ops)
		paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	}

	s := float32(gtx.Config.Px(unit.Sp(1)))
	w := float32(gtx.Constraints.Width.Max)
	h := float32(gtx.Constraints.Height.Max)
	var p clip.Path
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 2 * s, Y: 2 * s})
	p.Line(f32.Point{X: w - 4*s, Y: 0})
	p.Line(f32.Point{X: 0, Y: h - 4*s})
	p.Line(f32.Point{X: 4*s - w, Y: 0})
	p.Line(f32.Point{X: 0, Y: -s})
	p.Line(f32.Point{X: w - 5*s, Y: 0})
	p.Line(f32.Point{X: 0, Y: 6*s - h})
	p.Line(f32.Point{X: 5*s - w, Y: 0})
	p.Line(f32.Point{X: 0, Y: -s})
	p.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: w, Y: h}}}.Add(gtx.Ops)
}

type SelectCellEvent struct{}
