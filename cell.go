package foxtrot

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"image/color"
)

type Cells []Cell

func (cs Cells) Add(c Cell) {}

type Cell struct {
	in        string
	out       string
	promptNum int
	list      *layout.List
	inEditor  *widget.Editor
	outText   *widget.Label
	inLayout  *layout.Flex
	inLabel   *widget.Label
	outLabel  *widget.Label
	outLayout *layout.Flex
}

func newCell(in, out string, i int) Cell {
	list := &layout.List{Axis: layout.Vertical}
	inEditor := &widget.Editor{
		Submit: true}
	inEditor.SetText(in)
	outText := &widget.Label{}
	inLayout := &layout.Flex{}
	inLabel := &widget.Label{}
	outLabel := &widget.Label{}
	outLayout := &layout.Flex{}
	return Cell{in, out, i, list, inEditor, outText, inLayout, inLabel, outLabel, outLayout}
}

type EvalEvent struct{}

func (EvalEvent) ImplementsEvent() {}

func (c Cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.inEditor.Events(gtx) {
		if _, ok := e.(widget.SubmitEvent); ok {
			fmt.Println("Submit Cell")
			return EvalEvent{}
		}
	}
	return nil
}

func (c *Cell) evaluate() {
	textIn := c.inEditor.Text()
	if textIn == "" {
		return
	}
	//expIn := parser.Interp(textIn, a.engine)
	//expOut := a.engine.Eval(expIn)
	//textOut := expressionToString(a.engine, expOut, a.promptCount)
}

func (c Cell) Layout(gtx *layout.Context) {
	padding := unit.Dp(8)
	n := 1
	if c.out != "" {
		n = 2
	}
	c.list.Layout(gtx, n, func(i int) {
		if i == 0 {
			c1 := c.inLayout.Rigid(gtx, func() {
				theme.Label(_defaultFontSize, fmt.Sprintf(" In[%d] ", c.promptNum)).Layout(gtx)
			})
			c2 := c.inLayout.Flex(gtx, 1, func() {
				theme.Editor("").Layout(gtx, c.inEditor)
			})
			layout.Inset{Bottom: padding}.Layout(gtx, func() {
				c.inLayout.Layout(gtx, c1, c2)
			})
		} else {
			c1 := c.outLayout.Rigid(gtx, func() {
				theme.Label(_defaultFontSize, fmt.Sprintf("Out[%d] ", c.promptNum)).Layout(gtx)
			})
			c2 := c.outLayout.Flex(gtx, 1, func() {
				theme.Label(_defaultFontSize, c.out).Layout(gtx)
			})
			layout.Inset{Bottom: padding}.Layout(gtx, func() {
				c.outLayout.Layout(gtx, c1, c2)
			})
		}

	})
}

func rgb(c uint32) color.RGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func colorMaterial(ops *op.Ops, color color.RGBA) op.MacroOp {
	var mat op.MacroOp
	mat.Record(ops)
	paint.ColorOp{Color: color}.Add(ops)
	mat.Stop()
	return mat
}
