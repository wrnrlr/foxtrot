package cell

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/editor"
	"github.com/wrnrlr/foxtrot/output"
	"github.com/wrnrlr/foxtrot/theme"
)

type Cells []Cell

type Cell interface {
	Layout(selected bool, gtx *layout.Context)
	Event(gtx *layout.Context) interface{}
	Type() Type
	Text() string
	Focus()
	SetText(s string)
	SetLabel(s string)
	SetErr(err error)
	SetOut(ex expreduceapi.Ex)
}

type cell struct {
	typ   Type
	In    string
	Out   expreduceapi.Ex
	Label string

	Editor *editor.Editor
	margin *Margin

	styles *theme.Styles

	Err  error
	Hide bool
}

func NewCell(typ Type, label string, styles *theme.Styles) Cell {
	inEditor := &editor.Editor{}
	return &cell{typ: typ, Label: label, Editor: inEditor, margin: &Margin{}, styles: styles}
}

func (c cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.Editor.Events(gtx) {
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

func (c *cell) evaluate() {
	textIn := c.Editor.Text()
	if textIn == "" {
		return
	}
}

func (c *cell) Focus() {
	fmt.Println("cell.Focus")
	c.Editor.Focus()
}

func (c *cell) Text() string {
	return c.Editor.Text()
}

func (c *cell) Type() Type {
	return c.typ
}

func (c *cell) SetText(txt string) {
	c.Editor.SetText(txt)
}

func (c *cell) SetLabel(s string) {
	c.Label = s
}

func (c *cell) SetErr(err error) {
	c.Err = err
}
func (c *cell) SetOut(ex expreduceapi.Ex) {
	c.Out = ex
}

func (c *cell) Layout(selected bool, gtx *layout.Context) {
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		c.margin.Layout(gtx, selected, func() {
			c.cellLayout(gtx)
		})
	})
}

func (c *cell) cellLayout(gtx *layout.Context) {
	switch c.Type() {
	case Input:
		c.input(gtx)
	case Output:
		c.output(gtx)
	case H1:
		c.h1(gtx)
	case H2:
		c.h2(gtx)
	case H3:
		c.h3(gtx)
	case H4:
		c.h4(gtx)
	case Text:
		c.text(gtx)
	case Code:
		c.code(gtx)
	}
}

func (c *cell) input(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := layout.Rigid(func() {
		c.labelLayout(gtx)
	})
	c2 := layout.Flexed(1, func() {
		c.styles.Foxtrot.Layout(gtx, c.Editor)
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}

func (c *cell) output(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := layout.Rigid(func() {
		c.labelLayout(gtx)
	})
	c2 := layout.Flexed(1, func() {
		var stack op.StackOp
		stack.Push(gtx.Ops)
		//paint.ColorOp{Color: util.Black}.Add(gtx.Ops)
		w := output.FromEx(c.Out, gtx)
		w.Layout(gtx, c.styles.Theme.Shaper, text.Font{Size: unit.Sp(16)})
		stack.Pop()
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}

func (c *cell) h1(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H1.Layout(gtx, c.Editor)
	})
}

func (c *cell) h2(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H2.Layout(gtx, c.Editor)
	})
}

func (c *cell) h3(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H3.Layout(gtx, c.Editor)
	})
}

func (c *cell) h4(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H4.Layout(gtx, c.Editor)
	})
}

func (c *cell) text(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Text.Layout(gtx, c.Editor)
	})
}

func (c *cell) code(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Code.Layout(gtx, c.Editor)
	})
}

func (c *cell) labelLayout(gtx *layout.Context) {
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		px := gtx.Px(unit.Sp(50))
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		label := c.styles.Theme.Label(unit.Sp(12), c.Label)
		label.Alignment = text.End
		label.Layout(gtx)
	})
}
