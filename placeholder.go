package foxtrot

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
	"image"
	"time"
)

const (
	blinksPerSecond  = 1
	maxBlinkDuration = 10 * time.Second
)

type Placeholder struct {
	button  *widget.Button
	pbutton *widget.Button

	eventKey     int
	blinkStart   time.Time
	focused      bool
	caretOn      bool
	requestFocus bool

	clicker gesture.Click

	events []interface{}
}

func NewPlaceholder() *Placeholder {
	return &Placeholder{
		button:     new(widget.Button),
		pbutton:    new(widget.Button),
		blinkStart: time.Now()}
}

func (p *Placeholder) Event(gtx *layout.Context) interface{} {
	for p.button.Clicked(gtx) {
		fmt.Println("Placeholder Button Clicked")
		return AddCellEvent{}
	}
	for p.pbutton.Clicked(gtx) {
		fmt.Println("Focus Placeholder")
		p.requestFocus = true
	}
	return p.processKey(gtx)
}

func (p *Placeholder) processPointer(gtx *layout.Context) interface{} {
	for _, evt := range p.clicker.Events(gtx) {
		switch {
		case evt.Type == gesture.TypePress && evt.Source == pointer.Mouse,
			evt.Type == gesture.TypeClick && evt.Source == pointer.Touch:
			p.blinkStart = gtx.Now()
			p.requestFocus = true
		}
	}
	return nil
}

func (p *Placeholder) processKey(gtx *layout.Context) interface{} {
	for _, ke := range gtx.Events(&p.eventKey) {
		p.blinkStart = gtx.Now()
		switch ke := ke.(type) {
		case key.FocusEvent:
			p.focused = ke.Focus
			fmt.Println("Placeholder: key.FocusEvent")
		case key.Event:
			if !p.focused {
				fmt.Println("Placeholder (unfocused): key.Event")
				break
			}
			fmt.Println("Placeholder: key.Event")
			if ke.Name == key.NameReturn || ke.Name == key.NameEnter {
				return AddCellEvent{Type: FoxtrotCell}
			} else if ke.Name == key.NameUpArrow || ke.Name == key.NameLeftArrow {
				return FocusPreviousCellEvent{}
			} else if ke.Name == key.NameDownArrow || ke.Name == key.NameRightArrow {
				return FocusNextCellEvent{}
			}
		case key.EditEvent:
			fmt.Println("Placeholder: key.EditEvent")
		}
	}
	return nil
}

func (p *Placeholder) Focus() {
	p.requestFocus = true
	p.blinkStart = time.Now()
}

func (p *Placeholder) Layout(i int, gtx *layout.Context) {
	key.InputOp{Key: &p.eventKey, Focus: p.requestFocus}.Add(gtx.Ops)
	p.requestFocus = false
	p.clicker.Add(gtx.Ops)
	if p.focused {
		px := gtx.Config.Px(unit.Dp(20))
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Height = constraint
		st := layout.Stack{Alignment: layout.NW}
		c := st.Expand(gtx, func() {
			PlusButton{}.Layout(gtx, p.button)
		})
		l := st.Expand(gtx, func() {
			p.line(gtx)
			p.cursor(gtx)
		})
		st.Layout(gtx, l, c)
	} else {
		p.placeholderLayout(gtx)
	}
}

func (p *Placeholder) placeholderLayout(gtx *layout.Context) {
	px := gtx.Config.Px(unit.Sp(20))
	constraint := layout.Constraint{Min: px, Max: px}
	gtx.Constraints.Height = constraint
	fill(gtx, white)
	pointer.EllipseAreaOp{Rect: image.Rectangle{Max: gtx.Dimensions.Size}}.Add(gtx.Ops)
	p.pbutton.Layout(gtx)
}

func (p *Placeholder) line(gtx *layout.Context) {
	width := float32(gtx.Config.Px(unit.Sp(1)))
	var path paint.Path
	var lineLen = float32(gtx.Constraints.Width.Max)
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: 0, Y: merginTop})
	path.Line(f32.Point{X: lineLen, Y: 0})
	path.Line(f32.Point{X: 0, Y: width})
	path.Line(f32.Point{X: -lineLen, Y: 0})
	path.Line(f32.Point{X: 0, Y: -width})
	path.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{X: lineLen, Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
}

func (p *Placeholder) cursor(gtx *layout.Context) {
	p.caretOn = false
	now := gtx.Now()
	dt := now.Sub(p.blinkStart)
	const timePerBlink = time.Second / blinksPerSecond
	p.caretOn = dt%timePerBlink < timePerBlink/2
	if !p.caretOn {
		return
	}
	length := float32(gtx.Config.Px(unit.Sp(100)))
	width := float32(gtx.Config.Px(unit.Sp(1)))
	var path paint.Path
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var merginLeft = float32(gtx.Config.Px(unit.Sp(60)))
	var stack op.StackOp
	stack.Push(gtx.Ops)
	path.Begin(gtx.Ops)
	path.Move(f32.Point{X: merginLeft, Y: merginTop})
	path.Line(f32.Point{X: length, Y: 0})
	path.Line(f32.Point{X: 0, Y: width})
	path.Line(f32.Point{X: -length, Y: 0})
	path.Line(f32.Point{X: 0, Y: -width})
	path.End().Add(gtx.Ops)
	paint.ColorOp{black}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(length), Y: merginTop + width}}}.Add(gtx.Ops)
	stack.Pop()
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

type AddCellEvent struct {
	Type CellType
}

type FocusNextCellEvent struct{}

type FocusPreviousCellEvent struct{}
