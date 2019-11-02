package foxtrot

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"image"
)

type Add struct {
	Index  int
	button *widget.Button
	input  *widget.Editor
}

func NewAdd() *Add {
	input := &widget.Editor{SingleLine: true, Submit: true}
	return &Add{
		0,
		new(widget.Button),
		input}
}

type AddCellEvent struct{}

func (a *Add) Event(gtx *layout.Context) interface{} {
	for a.button.Clicked(gtx) {
		fmt.Println("Add Button Clicked")
		return AddCellEvent{}
	}
	for _, e := range a.input.Events(gtx) {
		if _, ok := e.(widget.SubmitEvent); ok {
			fmt.Println("Submit Cell")
			return AddCellEvent{}
		}
	}
	return nil
}

func (a *Add) Focus(i int) {
	a.Index = i
}

func (a *Add) Layout(i int, gtx *layout.Context) {
	length := gtx.Constraints.Width.Max
	gtx.Constraints.Height = inlineHeight
	st := layout.Stack{Alignment: layout.NW}
	c := st.Expand(gtx, func() {
		PlusButton{}.Layout(gtx, a.button)
	})
	l := st.Expand(gtx, func() {
		a.line(length, gtx)
		a.cursor(gtx)
	})
	st.Layout(gtx, l, c)
}

func (a *Add) line(length int, gtx *layout.Context) {
	width := float32(gtx.Config.Px(unit.Dp(2)))
	var p paint.Path
	var lineLen = float32(length)
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
			Max: f32.Point{X: float32(length), Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
}

func (a *Add) cursor(gtx *layout.Context) {
	length := float32(gtx.Config.Px(unit.Dp(100)))
	width := float32(gtx.Config.Px(unit.Dp(2)))
	var p paint.Path
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var merginLeft = float32(gtx.Config.Px(unit.Dp(60)))
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
	inset := layout.Inset{Left: unit.Dp(20)}
	inset.Layout(gtx, func() {
		px := gtx.Config.Px(unit.Dp(30))
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
	px := gtx.Config.Px(unit.Dp(30))
	size := float32(px)
	rr := float32(size) * .5
	var stack op.StackOp
	stack.Push(gtx.Ops)
	rrect(gtx.Ops, size, size, rr, rr, rr, rr)
	fill(gtx, lightGrey)
	stack.Pop()
}

func (b PlusButton) plus(gtx *layout.Context) {
	length := float32(gtx.Constraints.Width.Min)
	width := float32(gtx.Config.Px(unit.Dp(2)))
	var p1 paint.Path
	var center = float32(gtx.Constraints.Width.Min/2 - 1)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p1.Begin(gtx.Ops)
	p1.Move(f32.Point{X: 0, Y: center})
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
	p2.Move(f32.Point{X: center, Y: 0})
	p2.Line(f32.Point{X: 0, Y: length})
	p2.Line(f32.Point{X: width, Y: 0})
	p2.Line(f32.Point{X: 0, Y: -length})
	p2.Line(f32.Point{X: -width, Y: 0})
	p2.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: length, Y: length}}}.Add(gtx.Ops)
	stack.Pop()
}

type PlaceholderButton struct{}

func (b PlaceholderButton) Layout(gtx *layout.Context, button *widget.Button) {
	fill(gtx, lightPink)
	pointer.EllipseAreaOp{Rect: image.Rectangle{Max: gtx.Dimensions.Size}}.Add(gtx.Ops)
	button.Layout(gtx)
}

type placeholder struct {
	button *widget.Button
}

func newPlaceholder() placeholder {
	return placeholder{&widget.Button{}}
}

func (p placeholder) Layout(gtx *layout.Context) {
	gtx.Constraints.Height = inlineHeight
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
