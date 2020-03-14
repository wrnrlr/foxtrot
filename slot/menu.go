package slot

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type menu struct {
	active bool
	button *widget.Button
	list   *layout.List
	items  []item
}

type item struct {
	text   string
	button *widget.Button
}

var options = []item{
	{"Header 1", new(widget.Button)},
	{"Header 2", new(widget.Button)},
	{"Header 3", new(widget.Button)},
	{"Header 4", new(widget.Button)},
	{"Input", new(widget.Button)},
	{"Paragraph", new(widget.Button)},
	{"Code", new(widget.Button)}}

func (i item) layout(gtx *layout.Context) {

}

func (i item) event(gtx *layout.Context) SlotEvent {
	if i.button.Clicked(gtx) {
		return AddCell{Type: i.text}
	}
	return nil
}

func (m *menu) event(gtx *layout.Context) SlotEvent {
	if m.button.Clicked(gtx) {
		m.active = !m.active
	}
	for _, i := range m.items {
		return i.event(gtx)
	}
	return nil
}

func (m menu) layout(gtx *layout.Context) {
	if m.active {
		ins := layout.UniformInset(unit.Sp(20))
		ins.Layout(gtx, func() {
			length := len(m.items)
			m.list.Layout(gtx, length, nil)
		})
	}
}
