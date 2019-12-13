package notebook

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
	"gioui.org/widget"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/util"
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

			if ke.Name == key.NameUpArrow && ke.Modifiers.Contain(key.ModShift) {
				s.events = append(s.events, SelectPreviousCellEvent{})
			} else if ke.Name == key.NameDownArrow && ke.Modifiers.Contain(key.ModShift) {
				s.events = append(s.events, SelectNextCellEvent{})
			} else if ke.Name == key.NameReturn || ke.Name == key.NameEnter {
				s.events = append(s.events, AddCellEvent{Type: cell.Input})
			} else if ke.Name == key.NameUpArrow || ke.Name == key.NameLeftArrow {
				s.events = append(s.events, FocusPreviousCellEvent{})
			} else if ke.Name == key.NameDownArrow || ke.Name == key.NameRightArrow {
				s.events = append(s.events, FocusNextCellEvent{})
			} else if ke.Name == "1" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.H1})
			} else if ke.Name == "2" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.H2})
			} else if ke.Name == "3" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.H3})
			} else if ke.Name == "4" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.H4})
			} else if ke.Name == "5" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.Text})
			} else if ke.Name == "6" && ke.Modifiers.Contain(key.ModCommand) {
				s.events = append(s.events, AddCellEvent{Type: cell.Code})
			}
		case key.EditEvent:
			fmt.Println("Slot: key.EditEvent")
		}
	}
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

func (s *Slot) placeholderLayout(gtx *layout.Context) {
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

func (s *Slot) drawLine(gtx *layout.Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	var path clip.Path
	px := gtx.Px(unit.Dp(20))
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(px / 2)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: 0, Y: merginTop})
	path.Line(f32.Point{X: lineLen, Y: 0})
	path.Line(f32.Point{X: 0, Y: width})
	path.Line(f32.Point{X: -lineLen, Y: 0})
	path.Line(f32.Point{X: 0, Y: -width})
	path.End().Add(gtx.Ops)
	paint.ColorOp{util.LightGrey}.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{X: lineLen, Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
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
	var path clip.Path
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var merginLeft = float32(gtx.Px(unit.Sp(60)))
	var stack op.StackOp
	stack.Push(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: merginLeft, Y: merginTop})
	path.Line(f32.Point{X: length, Y: 0})
	path.Line(f32.Point{X: 0, Y: width})
	path.Line(f32.Point{X: -length, Y: 0})
	path.Line(f32.Point{X: 0, Y: -width})
	path.End().Add(gtx.Ops)
	paint.ColorOp{util.Black}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(length), Y: merginTop + width}}}.Add(gtx.Ops)
	stack.Pop()
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
	rrect(gtx.Ops, size, size, rr, rr, rr, rr)
	fill(gtx, util.LightGrey)
	stack.Pop()
}

func (b PlusButton) drawPlus(gtx *layout.Context) {
	width := float32(gtx.Px(unit.Sp(2)))
	offset := float32(gtx.Constraints.Width.Min) / 4
	length := float32(gtx.Constraints.Width.Min) - offset
	var p1 clip.Path
	var xcenter = float32(gtx.Constraints.Width.Min/2) - float32(gtx.Px(unit.Sp(1)))
	var ycenter = float32(gtx.Constraints.Height.Min/2) - float32(gtx.Px(unit.Sp(1)))
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p1.Begin(gtx.Ops)
	p1.Move(f32.Point{X: offset, Y: ycenter})
	p1.Line(f32.Point{X: length, Y: 0})
	p1.Line(f32.Point{X: 0, Y: width})
	p1.Line(f32.Point{X: -length, Y: 0})
	p1.Line(f32.Point{X: 0, Y: -width})
	p1.End().Add(gtx.Ops)
	paint.ColorOp{util.LightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: length, Y: length}}}.Add(gtx.Ops)
	stack.Pop()
	stack.Push(gtx.Ops)
	var p2 clip.Path
	p2.Begin(gtx.Ops)
	p2.Move(f32.Point{X: xcenter, Y: offset})
	p2.Line(f32.Point{X: 0, Y: length})
	p2.Line(f32.Point{X: width, Y: 0})
	p2.Line(f32.Point{X: 0, Y: -length})
	p2.Line(f32.Point{X: -width, Y: 0})
	p2.End().Add(gtx.Ops)
	paint.ColorOp{util.LightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: length, Y: length}}}.Add(gtx.Ops)
	stack.Pop()
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
