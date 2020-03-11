package cell

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/editor"
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
	SetType(s Type)
	SetLabel(s string)
	SetErr(err error)
	SetOut(ex expreduceapi.Ex)
}

type cell struct {
	typ     Type
	Content string
	Out     expreduceapi.Ex
	Label   string

	Input  *editor.Editor
	margin *Margin

	styles *theme.Styles

	Err  error
	Hide bool
	//Rules   map[string]string
}

func NewCell(typ Type, label string, styles *theme.Styles) Cell {
	inEditor := &editor.Editor{}
	return &cell{typ: typ, Label: label, Input: inEditor, margin: &Margin{}, styles: styles}
}

func (c cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.Input.Events(gtx) {
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
	textIn := c.Input.Text()
	if textIn == "" {
		return
	}
}

func (c *cell) Text() string {
	return c.Input.Text()
}

func (c *cell) Type() Type {
	return c.typ
}

func (c *cell) SetText(txt string) {
	c.Input.SetText(txt)
}

func (c *cell) SetType(typ Type) {
	c.typ = typ
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
