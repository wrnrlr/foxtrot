package foxtrot

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
)

type Margin struct {
	eventKey int
	click    gesture.Click
	checked  bool
}

func (m *Margin) SetChecked(value bool) {
	m.checked = value
}

func (m *Margin) Checked(gtx *layout.Context) bool {
	for _, e := range m.click.Events(gtx) {
		fmt.Printf("Margin clicked %v\n", e)
		switch e.Type {
		case gesture.TypeClick:
			m.checked = !m.checked
		}
	}
	return m.checked
}

func (m *Margin) Layout(gtx *layout.Context, widget layout.Widget) {
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
	m.layoutMargin(gtx)
	pointer.RectAreaOp{Rect: image.Rectangle{Max: image.Point{X: marginWidth, Y: editorHeight}}}.Add(gtx.Ops)
	m.click.Add(gtx.Ops)
	stack.Pop()

	dim.Size.Y = editorHeight
	gtx.Dimensions = dim
}

func (m *Margin) layoutMargin(gtx *layout.Context) {

	cs := gtx.Constraints

	if m.checked {
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
	var p paint.Path
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
