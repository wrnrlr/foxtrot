package foxtrot

import (
	"fmt"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/editor"
)

type Cell struct {
	Type      CellType
	in        string
	out       *Out
	promptNum int
	inEditor  *editor.Editor
	margin    *Margin
	styles    *Styles
}

func NewCell(typ CellType, i int, styles *Styles) *Cell {
	inEditor := &editor.Editor{}
	return &Cell{Type: typ, promptNum: i, inEditor: inEditor, margin: &Margin{}, styles: styles}
}

func (c Cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.inEditor.Events(gtx) {
		if ce, ok := e.(editor.CommandEvent); ok {
			lineCount, lineWidth, caretY, caretX := c.inEditor.CaretLine()
			if (ce.Event.Name == key.NameEnter || ce.Event.Name == key.NameReturn) && ce.Event.Modifiers.Contain(key.ModShift) {
				return EvalEvent{}
			} else if ce.Event.Name == key.NameUpArrow && caretY == 0 {
				return FocusPlaceholder{Offset: 0}
			} else if ce.Event.Name == key.NameLeftArrow && caretY == 0 && caretX == 0 {
				return FocusPlaceholder{Offset: 0}
			} else if ce.Event.Name == key.NameDownArrow && caretY == lineCount {
				return FocusPlaceholder{Offset: 1}
			} else if ce.Event.Name == key.NameRightArrow && caretY == lineCount && caretX == lineWidth {
				return FocusPlaceholder{Offset: 1}
			}
			fmt.Printf("Caret: lineCount: %v, lineWidth: %v, caretY: %v, caretX: %v\n", lineCount, lineWidth, caretY, caretX)
		}
	}
	return c.margin.Event(gtx)
}

func (c *Cell) evaluate() {
	textIn := c.inEditor.Text()
	if textIn == "" {
		return
	}
}

func (c *Cell) Layout(selected bool, gtx *layout.Context) {
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		c.margin.Layout(gtx, selected, func() {
			c.cellLayout(gtx)
		})
	})
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
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		px := gtx.Config.Px(promptWidth)
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		label := promptTheme.Label(_promptFontSize, txt)
		label.Alignment = text.End
		label.Layout(gtx)
	})
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

func (c *Cell) cellLayout(gtx *layout.Context) {
	switch c.Type {
	case FoxtrotCell:
		c.foxtrotCell(gtx)
	case TitleCell:
		c.titleCell(gtx)
	case SectionCell:
		c.sectionCell(gtx)
	case SubSectionCell:
		c.subSectionCell(gtx)
	case SubSubSectionCell:
		c.subSubSectionCell(gtx)
	case TextCell:
		c.textCell(gtx)
	case CodeCell:
		c.codeCell(gtx)
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
				c.styles.Foxtrot.Layout(gtx, c.inEditor)
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
	c.styles.Title.Layout(gtx, c.inEditor)
}

func (c *Cell) sectionCell(gtx *layout.Context) {
	c.styles.Section.Layout(gtx, c.inEditor)
}

func (c *Cell) subSectionCell(gtx *layout.Context) {
	c.styles.SubSection.Layout(gtx, c.inEditor)
}

func (c *Cell) subSubSectionCell(gtx *layout.Context) {
	c.styles.SubSubSection.Layout(gtx, c.inEditor)
}

func (c *Cell) textCell(gtx *layout.Context) {
	c.styles.Text.Layout(gtx, c.inEditor)
}

func (c *Cell) codeCell(gtx *layout.Context) {
	c.styles.Code.Layout(gtx, c.inEditor)
}

type CellType int

const (
	FoxtrotCell CellType = iota
	TitleCell
	SectionCell
	SubSectionCell
	SubSubSectionCell
	TextCell
	CodeCell
)

var CellTypeNames = []string{"Foxtrot", "Title", "Section", "SubSection", "SubSubSection", "Text", "Code"}

func (d CellType) String() string {
	return CellTypeNames[d]
}

func (d CellType) Level() int {
	switch d {
	case TitleCell:
		return 1
	case SectionCell:
		return 3
	case SubSectionCell:
		return 4
	case SubSubSectionCell:
		return 5
	default:
		return 0
	}
}

type EvalEvent struct{}

type FocusPlaceholder struct {
	Offset int
}

type FocusPreviousPlaceholder struct{}

type FocusNextPlaceholder struct{}
