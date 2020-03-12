package notebook

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/gesture"
	. "gioui.org/io/key"
	"gioui.org/io/pointer"
	. "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	. "github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/colors"
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

func (s *Slot) Event(isActive bool, gtx *Context) []interface{} {
	s.processEvents(isActive, gtx)
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *Slot) processEvents(isActive bool, gtx *Context) {
	for s.plusButton.Clicked(gtx) {
		s.events = append(s.events, AddCell{Type: Input})
	}
	for s.backgroundButton.Clicked(gtx) {
		s.Focus(true)
		s.events = append(s.events, SelectSlot{})
	}
	s.processKey(isActive, gtx)
}

func (s *Slot) processPointer(gtx *Context) interface{} {
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

func (s *Slot) processKey(isActive bool, gtx *Context) {
	if !isActive {
		return
	}
	for _, ke := range gtx.Events(&s.eventKey) {
		s.blinkStart = gtx.Now()
		switch ke := ke.(type) {
		case FocusEvent:
			s.focused = ke.Focus
		case Event:
			if !s.focused {
				break
			}
			e := s.switchKey(ke)
			if e != nil {
				s.events = append(s.events, e)
			}
		case EditEvent:
			fmt.Println("Slot: key.EditEvent")
		}
	}
}

func (Slot) switchKey(ke Event) interface{} {
	switch {
	case isKey(ke, NameUpArrow, ModShift):
		return SelectPreviousCell{}
	case isKey(ke, NameDownArrow, ModShift):
		return SelectNextCell{}
	case isKey(ke, NameReturn), isKey(ke, NameEnter):
		return AddCell{Type: Input}
	case isKey(ke, NameUpArrow), isKey(ke, NameLeftArrow):
		return FocusPreviousCell{}
	case isKey(ke, NameDownArrow), isKey(ke, NameRightArrow):
		return FocusNextCell{}
	case isKey(ke, "1", ModCommand):
		return AddCell{Type: H1}
	case isKey(ke, "2", ModCommand):
		return AddCell{Type: H2}
	case isKey(ke, "3", ModCommand):
		return AddCell{Type: H3}
	case isKey(ke, "4", ModCommand):
		return AddCell{Type: H4}
	case isKey(ke, "5", ModCommand):
		return AddCell{Type: Text}
	case isKey(ke, "6", ModCommand):
		return AddCell{Type: Code}
	default:
		return nil
	}
}

func isKey(e Event, k string, ms ...Modifiers) bool {
	return e.Name == k && hasMod(e, ms)
}

func hasMod(e Event, ms []Modifiers) bool {
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

func (s *Slot) Layout(isActive, isLast bool, gtx *Context) {
	// Flush events from before the previous frame.
	copy(s.events, s.events[s.prevEvents:])
	s.events = s.events[:len(s.events)-s.prevEvents]
	s.prevEvents = len(s.events)
	s.processEvents(isActive, gtx)
	s.layout(isActive, isLast, gtx)
}

func (s *Slot) layout(isActive, isLast bool, gtx *Context) {
	InputOp{Key: &s.eventKey, Focus: s.requestFocus}.Add(gtx.Ops)
	s.requestFocus = false
	if isLast {
		gtx.Constraints.Height.Min = 2000
	} else {
		px := gtx.Px(unit.Dp(20))
		constraint := Constraint{Min: px, Max: px}
		gtx.Constraints.Height = constraint
	}
	if isActive {
		st := Stack{Alignment: NW}
		c := Expanded(func() {
			PlusButton{}.Layout(gtx, s.plusButton)
		})
		l := Expanded(func() {
			//s.drawLine(gtx)
			s.drawCursor(gtx)
		})
		st.Layout(gtx, l, c)
	} else {
		s.placeholderLayout(gtx)
	}
}

func (s Slot) placeholderLayout(gtx *Context) {
	width := gtx.Constraints.Width.Max
	height := gtx.Constraints.Height.Max
	dr := f32.Rectangle{Max: f32.Point{X: float32(width), Y: float32(height)}}
	paint.ColorOp{Color: util.White}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = Dimensions{Size: image.Point{X: width, Y: height}}
	r := image.Rectangle{Max: gtx.Dimensions.Size}
	pointer.Rect(r).Add(gtx.Ops)
	s.backgroundButton.Layout(gtx)
}

func (s Slot) drawLine(gtx *Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Dp(20))
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(px / 2)
	line := shape.Line{{0, merginTop}, {lineLen, merginTop}}
	line.Stroke(util.LightGrey, width, gtx)
}

func (s *Slot) drawCursor(gtx *Context) {
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
	merginTop := float32(px / 2)
	merginLeft := float32(gtx.Px(unit.Sp(60)))
	line := shape.Line{{merginLeft, merginTop}, {length, merginTop}}
	line.Stroke(util.Black, width, gtx)
}

type PlusButton struct{}

func (b PlusButton) Layout(gtx *Context, button *widget.Button) {
	inset := Inset{Left: unit.Sp(20)}
	inset.Layout(gtx, func() {
		size := gtx.Px(unit.Sp(20))
		gtx.Constraints = RigidConstraints(image.Point{size, size})
		//b.drawCircle(gtx)
		b.drawPlus(gtx)
		r := image.Rectangle{Max: gtx.Dimensions.Size}
		pointer.Ellipse(r).Add(gtx.Ops)
		button.Layout(gtx)
	})
}

func (b PlusButton) drawCircle(gtx *Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Sp(20))
	size := float32(px)
	rr := float32(size) * .5
	c1 := shape.Circle{f32.Point{rr / 2, rr / 2}, rr}
	c1.Fill(util.White, gtx)
	c1.Stroke(util.LightGrey, width, gtx)
}

func (b PlusButton) drawPlus(gtx *Context) {
	s := gtx.Constraints
	w := float32(gtx.Px(unit.Sp(1)))
	yc := float32(s.Height.Min / 2)
	xc := float32(s.Width.Min / 2)
	offset := float32(s.Width.Min) / 4
	length := float32(gtx.Constraints.Width.Min) - offset
	line1 := shape.Line{{offset, yc}, {length, yc}}
	line1.Stroke(colors.Black, w, gtx)
	line2 := shape.Line{{xc, offset}, {xc, length}}
	line2.Stroke(colors.Black, w, gtx)
}

type AddCell struct {
	Type Type
}

type SelectSlot struct{}
type FocusNextCell struct{}
type FocusPreviousCell struct{}
type SelectNextCell struct{}
type SelectPreviousCell struct{}
type Slots struct{}
