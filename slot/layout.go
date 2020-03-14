package slot

import (
	"gioui.org/f32"
	. "gioui.org/io/key"
	"gioui.org/io/pointer"
	. "gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/util"
	"github.com/wrnrlr/shape"
	"image"
	"time"
)

func (s *slot) Layout(isLast bool, gtx *Context) {
	// Flush events from before the previous frame.
	copy(s.events, s.events[s.prevEvents:])
	s.events = s.events[:len(s.events)-s.prevEvents]
	s.prevEvents = len(s.events)
	s.processEvents(gtx)
	s.layout(isLast, gtx)
}

func (s *slot) layout(isLast bool, gtx *Context) {
	InputOp{Key: &s.eventKey, Focus: s.requestFocus}.Add(gtx.Ops)
	s.requestFocus = false
	s.layoutHeight(isLast, gtx)
	st := Stack{Alignment: NW}
	c := Expanded(func() {
		PlusButton{}.Layout(gtx, s.plusButton)
	})
	l := Expanded(func() {
		s.drawLine(gtx)
		s.drawCursor(gtx)
	})
	st.Layout(gtx, l, c)
}

func (s *slot) layoutHeight(isLast bool, gtx *Context) {
	if isLast {
		gtx.Constraints.Height.Min = 2000
	} else {
		px := gtx.Px(unit.Dp(20))
		constraint := Constraint{Min: px, Max: px}
		gtx.Constraints.Height = constraint
	}
}

func (s slot) placeholderLayout(gtx *Context) {
	width := gtx.Constraints.Width.Max
	height := gtx.Constraints.Height.Max
	dr := f32.Rectangle{Max: f32.Point{X: float32(width), Y: float32(height)}}
	paint.ColorOp{Color: util.White}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = Dimensions{Size: image.Point{X: width, Y: height}}
	r := image.Rectangle{Max: gtx.Dimensions.Size}
	pointer.Rect(r).Add(gtx.Ops)
	//s.backgroundButton.Layout(gtx)
}

func (s slot) drawLine(gtx *Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Dp(20))
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(px / 2)
	line := shape.Line{{0, merginTop}, {lineLen, merginTop}}
	line.Stroke(util.LightGrey, width, gtx)
}

func (s *slot) drawCursor(gtx *Context) {
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
