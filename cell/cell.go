package cell

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/editor"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/output"
	"github.com/wrnrlr/foxtrot/theme"
	"github.com/wrnrlr/foxtrot/util"
)

var promptWidth = unit.Sp(50)

type Cell struct {
	Type  CellType `xml:"type"`
	In    string
	Out   expreduceapi.Ex
	Label string

	inEditor *editor.Editor
	margin   *Margin
	styles   *theme.Styles

	Err  error
	Hide bool
}

func NewCell(typ CellType, i string, styles *theme.Styles) *Cell {
	inEditor := &editor.Editor{}
	return &Cell{Type: typ, Label: i, inEditor: inEditor, margin: &Margin{}, styles: styles}
}

func (c Cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.inEditor.Events(gtx) {
		switch e.(type) {
		case editor.SubmitEvent:
			return EvalEvent{}
		case editor.UpEvent:
			return FocusPlaceholder{Offset: 0}
		case editor.DownEvent:
			return FocusPlaceholder{Offset: 1}
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

func (c *Cell) labelLayout(gtx *layout.Context) {
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		px := gtx.Config.Px(promptWidth)
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		label := c.styles.Theme.Label(unit.Sp(12), c.Label)
		label.Alignment = text.End
		label.Layout(gtx)
	})
}

func (c *Cell) Text() string {
	return c.inEditor.Text()
}

func (c *Cell) cellLayout(gtx *layout.Context) {
	switch c.Type {
	case InputCell:
		c.inputCell(gtx)
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
	case OutputCell:
		c.outputCell(gtx)
	}
}

func (c *Cell) inputCell(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := f.Rigid(gtx, func() {
		c.labelLayout(gtx)
	})
	c2 := f.Flex(gtx, 1, func() {
		c.styles.Foxtrot.Layout(gtx, c.inEditor)
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}

func (c *Cell) titleCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Title.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) sectionCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Section.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) subSectionCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.SubSection.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) subSubSectionCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.SubSubSection.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) textCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Text.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) codeCell(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Code.Layout(gtx, c.inEditor)
	})
}

func (c *Cell) outputCell(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := f.Rigid(gtx, func() {
		c.labelLayout(gtx)
	})
	c2 := f.Flex(gtx, 1, func() {
		st := graphics.NewStyle()
		var stack op.StackOp
		stack.Push(gtx.Ops)
		paint.ColorOp{Color: util.Black}.Add(gtx.Ops)
		w := output.Ex(c.Out, st, gtx)
		w.Layout(gtx, c.styles.Theme.Shaper, text.Font{Size: unit.Sp(16)})
		stack.Pop()
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}
