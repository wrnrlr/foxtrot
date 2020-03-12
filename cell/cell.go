package cell

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/editor"
	. "github.com/wrnrlr/foxtrot/slot"
	"github.com/wrnrlr/foxtrot/theme"
)

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
	typ Type

	content string
	out     expreduceapi.Ex
	label   string

	input *editor.Editor

	Err    error
	Hide   bool
	slot   Slot
	margin *Margin
	styles *theme.Styles
	// Ex // Expression of cell
	//Rules   map[string]string
}

func NewCell(typ Type, label string, styles *theme.Styles) Cell {
	inEditor := &editor.Editor{}
	return &cell{typ: typ, label: label, input: inEditor, margin: &Margin{}, styles: styles}
}

func (c *cell) evaluate() {
	textIn := c.input.Text()
	if textIn == "" {
		return
	}
}

func (c *cell) Text() string {
	return c.input.Text()
}

func (c *cell) Type() Type {
	return c.typ
}

func (c *cell) SetText(txt string) {
	c.input.SetText(txt)
}

func (c *cell) SetType(typ Type) {
	c.typ = typ
}

func (c *cell) SetLabel(s string) {
	c.label = s
}

func (c *cell) SetErr(err error) {
	c.Err = err
}
func (c *cell) SetOut(ex expreduceapi.Ex) {
	c.out = ex
}
