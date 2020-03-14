package cell

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/editor"
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

	ToEx() *atoms.Expression
}

func NewCell(typ Type, label string, styles *theme.Styles) Cell {
	inEditor := &editor.Editor{}
	return &cell{typ: typ, label: label, input: inEditor, margin: &Margin{}, styles: styles}
}

type cell struct {
	typ Type

	content string
	label   string

	out expreduceapi.Ex

	input *editor.Editor

	err    error
	hide   bool
	slot   widget.Button
	margin *Margin
	styles *theme.Styles
	// Ex // Expression of cell
	//Rules   map[string]string
}

func (c *cell) evaluate() {
	textIn := c.input.Text()
	if textIn == "" {
		return
	}
}

func (c cell) Text() string {
	return c.input.Text()
}

func (c cell) Type() Type {
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
	c.err = err
}
func (c *cell) SetOut(ex expreduceapi.Ex) {
	c.out = ex
}

func (c cell) ToEx() *atoms.Expression {
	return nil
}
