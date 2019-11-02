package foxtrot

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"image"
	"time"
)

const (
	blinksPerSecond  = 1
	maxBlinkDuration = 10 * time.Second
)

type Add struct {
	Index int
	//Focus bool
	button       *widget.Button
	input        *widget.Editor
	caretOn      bool
	eventKey     int
	requestFocus bool
	focused      bool
	blinkStart   time.Time
}

func NewAdd() *Add {
	return &Add{
		Index:      0,
		button:     new(widget.Button),
		blinkStart: time.Now()}
}

type AddCellEvent struct{}

func (a *Add) Event(gtx *layout.Context) interface{} {
	for a.button.Clicked(gtx) {
		fmt.Println("Add Button Clicked")
		return AddCellEvent{}
	}
	//for _, e := range a.input.Events(gtx) {
	//	if _, ok := e.(widget.SubmitEvent); ok {
	//		fmt.Println("Submit Cell")
	//		return AddCellEvent{}
	//	}
	//}
	a.processKey(gtx)
	return nil
}

func (a *Add) processKey(gtx *layout.Context) {
	for _, ke := range gtx.Events(&a.eventKey) {
		a.blinkStart = gtx.Now()
		switch ke.(type) {
		case key.FocusEvent:
			fmt.Println("Add: key.FocusEvent")
		case key.Event:
			fmt.Println("Add: key.Event")
		case key.EditEvent:
			fmt.Println("Add: key.EditEvent")
		}
	}
}

func (a *Add) Focus(i int) {
	a.Index = i
	a.focused = true
	a.blinkStart = time.Now()
}

func (a *Add) Layout(i int, gtx *layout.Context) {
	key.InputOp{Key: &a.eventKey, Focus: true}.Add(gtx.Ops)
	px := gtx.Config.Px(unit.Dp(20))
	constraint := layout.Constraint{Min: px, Max: px}
	gtx.Constraints.Height = constraint
	st := layout.Stack{Alignment: layout.NW}
	c := st.Expand(gtx, func() {
		PlusButton{}.Layout(gtx, a.button)
	})
	l := st.Expand(gtx, func() {
		a.line(gtx)
		a.cursor(gtx)
	})
	st.Layout(gtx, l, c)
}

func (a *Add) line(gtx *layout.Context) {
	width := float32(gtx.Config.Px(unit.Sp(1)))
	var p paint.Path
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 0, Y: merginTop})
	p.Line(f32.Point{X: lineLen, Y: 0})
	p.Line(f32.Point{X: 0, Y: width})
	p.Line(f32.Point{X: -lineLen, Y: 0})
	p.Line(f32.Point{X: 0, Y: -width})
	p.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{X: lineLen, Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
}

func (a *Add) cursor(gtx *layout.Context) {
	a.caretOn = false
	now := gtx.Now()
	dt := now.Sub(a.blinkStart)
	const timePerBlink = time.Second / blinksPerSecond
	a.caretOn = dt%timePerBlink < timePerBlink/2
	if !a.caretOn {
		return
	}
	length := float32(gtx.Config.Px(unit.Sp(100)))
	width := float32(gtx.Config.Px(unit.Sp(1)))
	var p paint.Path
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var merginLeft = float32(gtx.Config.Px(unit.Sp(60)))
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: merginLeft, Y: merginTop})
	p.Line(f32.Point{X: length, Y: 0})
	p.Line(f32.Point{X: 0, Y: width})
	p.Line(f32.Point{X: -length, Y: 0})
	p.Line(f32.Point{X: 0, Y: -width})
	p.End().Add(gtx.Ops)
	paint.ColorOp{black}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(length), Y: merginTop + width}}}.Add(gtx.Ops)
	stack.Pop()
}

func (a *Add) IsIndex(i int) bool {
	return a.Index == i
}

type PlusButton struct{}

func (b PlusButton) Layout(gtx *layout.Context, button *widget.Button) {
	inset := layout.Inset{Left: unit.Sp(20)}
	inset.Layout(gtx, func() {
		px := gtx.Config.Px(unit.Sp(20))
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		gtx.Constraints.Height = constraint
		b.circle(gtx)
		b.plus(gtx)
		pointer.EllipseAreaOp{Rect: image.Rectangle{Max: gtx.Dimensions.Size}}.Add(gtx.Ops)
		button.Layout(gtx)
	})
}

func (b PlusButton) circle(gtx *layout.Context) {
	px := gtx.Config.Px(unit.Sp(20))
	size := float32(px)
	rr := float32(size) * .5
	var stack op.StackOp
	stack.Push(gtx.Ops)
	rrect(gtx.Ops, size, size, rr, rr, rr, rr)
	fill(gtx, lightGrey)
	stack.Pop()
}

func (b PlusButton) plus(gtx *layout.Context) {
	width := float32(gtx.Config.Px(unit.Sp(2)))
	offset := float32(gtx.Constraints.Width.Min) / 4
	length := float32(gtx.Constraints.Width.Min) - offset
	var p1 paint.Path
	var xcenter = float32(gtx.Constraints.Width.Min/2) - float32(gtx.Config.Px(unit.Sp(1)))
	var ycenter = float32(gtx.Constraints.Height.Min/2) - float32(gtx.Config.Px(unit.Sp(1)))
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p1.Begin(gtx.Ops)
	p1.Move(f32.Point{X: offset, Y: ycenter})
	p1.Line(f32.Point{X: length, Y: 0})
	p1.Line(f32.Point{X: 0, Y: width})
	p1.Line(f32.Point{X: -length, Y: 0})
	p1.Line(f32.Point{X: 0, Y: -width})
	p1.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: length, Y: length}}}.Add(gtx.Ops)
	stack.Pop()
	stack.Push(gtx.Ops)
	var p2 paint.Path
	p2.Begin(gtx.Ops)
	p2.Move(f32.Point{X: xcenter, Y: offset})
	p2.Line(f32.Point{X: 0, Y: length})
	p2.Line(f32.Point{X: width, Y: 0})
	p2.Line(f32.Point{X: 0, Y: -length})
	p2.Line(f32.Point{X: -width, Y: 0})
	p2.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: length, Y: length}}}.Add(gtx.Ops)
	stack.Pop()
}

type placeholder struct {
	button *widget.Button
}

func newPlaceholder() placeholder {
	return placeholder{&widget.Button{}}
}

func (p placeholder) Layout(gtx *layout.Context) {
	px := gtx.Config.Px(unit.Sp(20))
	constraint := layout.Constraint{Min: px, Max: px}
	gtx.Constraints.Height = constraint
	fill(gtx, white)
	pointer.EllipseAreaOp{Rect: image.Rectangle{Max: gtx.Dimensions.Size}}.Add(gtx.Ops)
	p.button.Layout(gtx)
}

func (p placeholder) Event(gtx *layout.Context) interface{} {
	for p.button.Clicked(gtx) {
		fmt.Println("Focus Placeholder")
		return FocusPlaceholder{}
	}
	return nil
}

type FocusPlaceholder struct{}
