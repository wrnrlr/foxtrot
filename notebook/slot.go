package notebook

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/util"
	"github.com/wrnrlr/shape"
	"image"
	"time"
)

const (
	blinksPerSecond  = 1
	maxBlinkDuration = 10 * time.Second
)

// Every cell has a slot above and below it that allow new Cells to be inserted insert.
type Slot struct {
	Active           bool
	plusButton       *widget.Button
	backgroundButton *widget.Button

	eventKey     int
	blinkStart   time.Time
	focused      bool
	caretOn      bool
	requestFocus bool

	clicker gesture.Click

	events     []interface{}
	prevEvents int
}

func NewSlot() *Slot {
	return &Slot{
		plusButton:       new(widget.Button),
		backgroundButton: new(widget.Button),
		blinkStart:       time.Now()}
}

func (s *Slot) Event(isActive bool, gtx *layout.Context) []interface{} {
	s.processEvents(isActive, gtx)
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *Slot) processEvents(isActive bool, gtx *layout.Context) {
	for s.plusButton.Clicked(gtx) {
		s.events = append(s.events, AddCellEvent{Type: cell.Input})
	}
	for s.backgroundButton.Clicked(gtx) {
		s.Focus(true)
		s.events = append(s.events, SelectSlotEvent{})
	}
	s.processKey(isActive, gtx)
}

func (s *Slot) processPointer(gtx *layout.Context) interface{} {
	if !s.focused {
		return nil
	}
	for _, evt := range s.clicker.Events(gtx) {
		switch {
		case evt.Type == gesture.TypePress && evt.Source == pointer.Mouse,
			evt.Type == gesture.TypeClick && evt.Source == pointer.Touch:
			s.blinkStart = gtx.Now()
			s.requestFocus = true
		}
	}
	return nil
}

func (s *Slot) processKey(isActive bool, gtx *layout.Context) {
	if !isActive {
		return
	}
	for _, ke := range gtx.Events(&s.eventKey) {
		s.blinkStart = gtx.Now()
		switch ke := ke.(type) {
		case key.FocusEvent:
			s.focused = ke.Focus
		case key.Event:
			if !s.focused {
				break
			}
			e := s.switchKey(ke)
			if e != nil {
				s.events = append(s.events, e)
			}
		case key.EditEvent:
			fmt.Println("Slot: key.EditEvent")
		}
	}
}

func (Slot) switchKey(ke key.Event) (e interface{}) {
	switch {
	case isKey(ke, key.NameUpArrow, key.ModShift):
		e = SelectPreviousCellEvent{}
	case isKey(ke, key.NameDownArrow, key.ModShift):
		e = SelectNextCellEvent{}
	case isKey(ke, key.NameReturn), isKey(ke, key.NameEnter):
		e = AddCellEvent{Type: cell.Input}
	case isKey(ke, key.NameUpArrow), isKey(ke, key.NameLeftArrow):
		e = FocusPreviousCellEvent{}
	case isKey(ke, key.NameDownArrow), isKey(ke, key.NameRightArrow):
		e = FocusNextCellEvent{}
	case isKey(ke, "1", key.ModCommand):
		e = AddCellEvent{Type: cell.H1}
	case isKey(ke, "2", key.ModCommand):
		e = AddCellEvent{Type: cell.H2}
	case isKey(ke, "3", key.ModCommand):
		e = AddCellEvent{Type: cell.H3}
	case isKey(ke, "4", key.ModCommand):
		e = AddCellEvent{Type: cell.H4}
	case isKey(ke, "5", key.ModCommand):
		e = AddCellEvent{Type: cell.Text}
	case isKey(ke, "6", key.ModCommand):
		e = AddCellEvent{Type: cell.Code}
	}
	return e
}

func isKey(e key.Event, k string, ms ...key.Modifiers) bool {
	return e.Name == k && hasMod(e, ms)
}

func hasMod(e key.Event, ms []key.Modifiers) bool {
	for _, m := range ms {
		if !e.Modifiers.Contain(m) {
			return false
		}
	}
	return true
}

func (s *Slot) Focus(requestFocus bool) {
	s.requestFocus = requestFocus
}

func (s *Slot) Layout(isActive, isLast bool, gtx *layout.Context) {
	// Flush events from before the previous frame.
	copy(s.events, s.events[s.prevEvents:])
	s.events = s.events[:len(s.events)-s.prevEvents]
	s.prevEvents = len(s.events)
	s.processEvents(isActive, gtx)
	s.layout(isActive, isLast, gtx)
}

func (s *Slot) layout(isActive, isLast bool, gtx *layout.Context) {
	key.InputOp{Key: &s.eventKey, Focus: s.requestFocus}.Add(gtx.Ops)
	s.requestFocus = false
	if isLast {
		gtx.Constraints.Height.Min = 2000
	} else {
		px := gtx.Px(unit.Dp(20))
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Height = constraint
	}
	if isActive {
		st := layout.Stack{Alignment: layout.NW}
		c := layout.Expanded(func() {
			PlusButton{}.Layout(gtx, s.plusButton)
		})
		l := layout.Expanded(func() {
			s.drawLine(gtx)
			s.drawCursor(gtx)
		})
		st.Layout(gtx, l, c)
	} else {
		s.placeholderLayout(gtx)
	}
}

func (s Slot) placeholderLayout(gtx *layout.Context) {
	width := gtx.Constraints.Width.Max
	height := gtx.Constraints.Height.Max
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(width), Y: float32(height)},
	}
	paint.ColorOp{Color: util.White}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: image.Point{X: width, Y: height}}
	r := image.Rectangle{Max: gtx.Dimensions.Size}
	pointer.Rect(r).Add(gtx.Ops)
	s.backgroundButton.Layout(gtx)
}

func (s Slot) drawLine(gtx *layout.Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Dp(20))
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(px / 2)
	line := shape.Line{{0, merginTop}, {lineLen, merginTop}}
	line.Stroke(util.LightGrey, width, gtx)
}

func (s *Slot) drawCursor(gtx *layout.Context) {
	if !s.focused {
		return
	}
	s.caretOn = false
	now := gtx.Now()
	dt := now.Sub(s.blinkStart)
	blinking := dt < maxBlinkDuration
	const timePerBlink = time.Second / blinksPerSecond
	nextBlink := now.Add(timePerBlink/2 - dt%(timePerBlink/2))
	if blinking {
		redraw := op.InvalidateOp{At: nextBlink}
		redraw.Add(gtx.Ops)
	}
	s.caretOn = s.focused && (!blinking || dt%timePerBlink < timePerBlink/2)
	if !s.caretOn {
		return
	}
	length := float32(gtx.Px(unit.Sp(100)))
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Dp(20))
	var merginTop = float32(px / 2)
	var merginLeft = float32(gtx.Px(unit.Sp(60)))
	line := shape.Line{{merginLeft, merginTop}, {length, merginTop}}
	line.Stroke(util.Black, width, gtx)
}

type PlusButton struct{}

func (b PlusButton) Layout(gtx *layout.Context, button *widget.Button) {
	inset := layout.Inset{Left: unit.Sp(20)}
	inset.Layout(gtx, func() {
		size := gtx.Px(unit.Sp(20))
		gtx.Constraints = layout.RigidConstraints(image.Point{size, size})
		b.drawCircle(gtx)
		b.drawPlus(gtx)
		r := image.Rectangle{Max: gtx.Dimensions.Size}
		pointer.Ellipse(r).Add(gtx.Ops)
		button.Layout(gtx)
	})
}

func (b PlusButton) drawCircle(gtx *layout.Context) {
	px := gtx.Px(unit.Sp(20))
	size := float32(px)
	rr := float32(size) * .5
	var stack op.StackOp
	stack.Push(gtx.Ops)
	//shape.StrokeCircle(f32.Point{rr,rr}, rr, float32(gtx.Px(unit.Sp(0.01))), gtx.Ops)
	//paint.ColorOp{util.LightGrey}.Add(gtx.Ops)
	rrect(gtx.Ops, size, size, rr, rr, rr, rr)
	fill(gtx, util.LightGrey)
	stack.Pop()
}

func (b PlusButton) drawPlus(gtx *layout.Context) {
	size := gtx.Constraints
	width := float32(gtx.Px(unit.Sp(1)))
	ycenter := float32(size.Height.Min / 2)
	xcenter := float32(size.Width.Min / 2)
	offset := float32(size.Width.Min) / 4
	length := float32(gtx.Constraints.Width.Min) - offset
	line1 := shape.Line{{offset, ycenter}, {length, ycenter}}
	line1.Stroke(util.LightGrey, width, gtx)
	line2 := shape.Line{{xcenter, offset}, {xcenter, length}}
	line2.Stroke(util.LightGrey, width, gtx)
}

type AddCellEvent struct {
	Type cell.Type
}

type SelectSlotEvent struct{}

type FocusNextCellEvent struct{}

type FocusPreviousCellEvent struct{}

type SelectNextCellEvent struct{}

type SelectPreviousCellEvent struct{}

type Slots struct {
}
