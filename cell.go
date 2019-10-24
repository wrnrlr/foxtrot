package foxtrot

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
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
	outLayout *layout.Flex
	inEx      *expreduceapi.Ex
	outEx     *expreduceapi.Ex
}

func newCell(in, out string, i int) Cell {
	list := &layout.List{Axis: layout.Vertical}
	inEditor := &widget.Editor{Submit: true}
	inEditor.SetText(in)
	outText := &widget.Label{}
	inLayout := &layout.Flex{}
	outLayout := &layout.Flex{}
	return Cell{in, out, i, list, inEditor, outText, inLayout, outLayout, nil, nil}
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
				c.inLabel().Layout(gtx)
			})
			c2 := c.inLayout.Flex(gtx, 1, func() {
				ed := theme.Editor("")
				ed.Font = _monoFont
				ed.Layout(gtx, c.inEditor)
			})
			layout.Inset{Bottom: padding}.Layout(gtx, func() {
				c.inLayout.Layout(gtx, c1, c2)
			})
		} else {
			c1 := c.outLayout.Rigid(gtx, func() {
				c.outLabel().Layout(gtx)
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

func (c *Cell) Focus() {
	c.inEditor.Focus()
}

func (c *Cell) outLabel() material.Label {
	l := theme.Label(_defaultFontSize, fmt.Sprintf("Out[%d] ", c.promptNum))
	l.Font = _monoFont
	return l
}

func (c *Cell) inLabel() material.Label {
	var text string
	if c.promptNum < 0 {
		text = fmt.Sprintf(" In[ ] ")
	} else {
		text = fmt.Sprintf(" In[%d] ", c.promptNum)
	}
	l := theme.Label(_defaultFontSize, text)
	l.Font = _monoFont
	return l
}
