package foxtrot

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Add struct {
	button *widget.Button
	input  *widget.Editor
}

func newAdd() *Add {
	return &Add{
		new(widget.Button),
		new(widget.Editor)}
}

func (a *Add) Layout(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	e1 := f.Rigid(gtx, func() {
		theme.Button("+").Layout(gtx, a.button)
	})
	e2 := f.Flex(gtx, 1, func() {
		theme.Editor("").Layout(gtx, a.input)
	})
	layout.UniformInset(unit.Dp(8)).Layout(gtx, func() {
		f.Layout(gtx, e1, e2)
	})
}

type AddCellEvent struct{}

func (a *Add) Event(gtx *layout.Context) interface{} {
	for a.button.Clicked(gtx) {
		fmt.Println("Add Button Clicked")
		return AddCellEvent{}
	}
	return nil
}

func (a *Add) Focus() {
	a.input.Focus()
}
