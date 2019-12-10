package notebook

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"github.com/stretchr/testify/assert"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/theme"
	"testing"
)

func TestNewNotebook(t *testing.T) {
	nb := NewNotebook(nil)
	assert.Equal(t, 0, nb.Size())
}

func TestNewNotebookWithCells(t *testing.T) {
	style := theme.DefaultStyles()
	var cells cell.Cells
	c := cell.NewCell(cell.Input, "In[0]:=", style)
	cells = append(cells, c)
	nb := NewNotebook(nil)
	assert.Equal(t, 1, nb.Size())
}

func TestDeleteCell(t *testing.T) {
	style := theme.DefaultStyles()
	var cells cell.Cells
	c := cell.NewCell(cell.Input, "In[0]:=", style)
	cells = append(cells, c)
	nb := NewNotebook(nil)
	nb.DeleteCell(0)
	assert.Equal(t, 0, nb.Size())
}

type evalCell struct {
	cell.Cell
	event interface{}
}

func (c *evalCell) Event(gtx *layout.Context) interface{} {
	return c.event
}

type queue struct{}

func (queue) Events(k event.Key) []event.Event {
	return nil
}

var q queue

func TestEvalCell(t *testing.T) {
	gtx := &layout.Context{Queue: q}
	style := theme.DefaultStyles()
	var cells cell.Cells
	c := cell.NewCell(cell.Input, "In[0]:=", style)
	c.SetText("1+1")
	ec := &evalCell{Cell: c, event: cell.EvalEvent{}}
	cells = append(cells, ec)
	nb := NewNotebook(cells)
	nb.Event(gtx)
	assert.Equal(t, 2, nb.Size())
}
