package foxtrot

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Cell struct {
	Type      CellType
	in        string
	out       *Out
	promptNum int
	inEditor  *widget.Editor
}

func NewCell(typ CellType, i int) *Cell {
	inEditor := &widget.Editor{Submit: true}
	return &Cell{Type: typ, promptNum: i, inEditor: inEditor}
}

func (c Cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.inEditor.Events(gtx) {
		if e, ok := e.(widget.SubmitEvent); ok {
			c.inEditor.SetText(e.Text)
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

func (c *Cell) Layout(gtx *layout.Context) {
	switch c.Type {
	case FoxtrotCell:
		c.foxtrotCell(gtx)
	case TitleCell:
		c.titleCell(gtx)
	}
}

func (c *Cell) Focus() {
	fmt.Println("Cell.Focus")
	c.inEditor.Focus()
}

func (c *Cell) promptLayout(gtx *layout.Context) {
	var txt string
	if c.promptNum < 0 {
		txt = fmt.Sprintf("In[ ] ")
	} else {
		txt = fmt.Sprintf("In[%d] ", c.promptNum)
	}
	px := gtx.Config.Px(promptWidth)
	constraint := layout.Constraint{Min: px, Max: px}
	gtx.Constraints.Width = constraint
	label := promptTheme.Label(_promptFontSize, txt)
	label.Alignment = text.End
	label.Layout(gtx)
}

func (c *Cell) inEditor2() material.Editor {
	ed := theme.Editor("")
	ed.Font.Size = _defaultFontSize
	ed.Font.Variant = "Mono"
	return ed
}

func (c *Cell) hasOut() bool {
	return c.out != nil
}

func (c *Cell) itemCount() int {
	if c.hasOut() {
		return 2
	} else {
		return 1
	}
}

func (c *Cell) foxtrotCell(gtx *layout.Context) {
	n := c.itemCount()
	list := &layout.List{Axis: layout.Vertical}
	list.Layout(gtx, n, func(i int) {
		if i == 0 {
			f := layout.Flex{Alignment: layout.Middle}
			c1 := f.Rigid(gtx, func() {
				c.promptLayout(gtx)
			})
			c2 := f.Flex(gtx, 1, func() {
				c.inEditor2().Layout(gtx, c.inEditor)
			})
			layout.Inset{Bottom: _padding}.Layout(gtx, func() {
				f.Layout(gtx, c1, c2)
			})
		} else {
			c.out.Layout(c.promptNum, gtx)
		}
	})
}

func (c *Cell) titleCell(gtx *layout.Context) {
	FoxtrotTheme.Editor("Title").Layout(gtx, c.inEditor)
}

func (c *Cell) subSectionCell(gtx *layout.Context) {

}

func (c *Cell) subSubSectionCell(gtx *layout.Context) {

}

func (c *Cell) textCell(gtx *layout.Context) {

}

func (c *Cell) codeCell(gtx *layout.Context) {

}

type cellType interface {
	Layout(gtx *layout.Context)
}

type CellType int

const (
	FoxtrotCell CellType = iota
	TitleCell
	SubSectionCell
	SubSubSectionCell
	TextCell
	CodeCell
)

var CellTypeNames = []string{"Foxtrot", "Title", "SubSection", "SubSubSection", "Text", "Code"}

func (d CellType) String() string {
	return CellTypeNames[d]
}

type SelectCellEvent struct{}

type EvalEvent struct{}

type FocusPreviousPlaceholder struct{}

type FocusNextPlaceholder struct{}
