package foxtrot

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/key"
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
	//checked  bool
}

//func (m *Margin) SetChecked(value bool) {
//	m.checked = value
//}

func (m *Margin) Event(gtx *layout.Context) interface{} {
	for _, e := range m.click.Events(gtx) {
		fmt.Printf("Margin clicked %v\n", e)
		switch e.Type {
		case gesture.TypeClick:
			//m.checked = !m.checked
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

type Selection struct {
	eventKey     int
	count        int
	events       []interface{}
	prevEvents   int
	focused      bool
	requestFocus bool

	begin, end int
}

func (s *Selection) Event(gtx *layout.Context) []interface{} {
	s.processEvents(gtx)
	return s.flushEvents()
}

func (s *Selection) flushEvents() []interface{} {
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *Selection) processEvents(gtx *layout.Context) {
	key.InputOp{Key: &s.eventKey, Focus: s.requestFocus}.Add(gtx.Ops)
	s.requestFocus = false
	for _, e := range gtx.Events(&s.eventKey) {
		switch ke := e.(type) {
		case key.Event:
			fmt.Printf("Selection Key Event\n")
			if ke.Name == key.NameDeleteBackward || ke.Name == key.NameDeleteForward {
				s.events = append(s.events, DeleteSelected{})
			}
		case key.FocusEvent:
			fmt.Printf("Selection Focus Event: %b\n", ke.Focus)
			s.focused = ke.Focus
		}
	}
}

func (s *Selection) RequestFocus(b bool, gtx *layout.Context) {
	fmt.Printf("Selection Request Focus: %b\n", b)
	s.requestFocus = b
	s.processEvents(gtx)
}

func (s *Selection) Clear() {
	s.begin = -1
	s.end = -1
}

func (s *Selection) SetBegin(i int) {
	s.begin = i
	s.end = i
}

func (s *Selection) SetEnd(i int) {
	s.end = i
}

func (s *Selection) IsSelected(i int) bool {
	return s.begin != -1 && s.min() >= i && s.max() <= i
}

func (s *Selection) min() int {
	if s.begin < s.end {
		return s.begin
	} else {
		return s.end
	}
}

func (s *Selection) max() int {
	if s.begin > s.end {
		return s.begin
	} else {
		return s.end
	}
}

type SelectCellEvent struct{}

type DeleteSelected struct{}
