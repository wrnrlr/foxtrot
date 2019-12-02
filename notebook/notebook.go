package notebook

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/parser"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/theme"
	"io/ioutil"
)

type Notebook struct {
	cells       []*cell.Cell
	slots       []*Slot
	kernel      *expreduce.EvalState
	promptCount int

	activeSlot int
	list       layout.List
	selection  *Selection
	styles     *theme.Styles
}

func NewNotebook() *Notebook {
	cells := make([]*cell.Cell, 0)
	kernel := expreduce.NewEvalState()
	firstSlot := NewSlot()
	adds := []*Slot{firstSlot}
	selection := NewSelection()
	styles := theme.DefaultStyles()
	return &Notebook{cells, adds, kernel, 1, 0, layout.List{Axis: layout.Vertical}, selection, styles}
}

func (nb *Notebook) Event(gtx *layout.Context) interface{} {
	for i, c := range nb.cells {
		e := c.Event(gtx)
		switch e := e.(type) {
		case cell.EvalEvent:
			nb.eval(i)
		case cell.SelectFirstCellEvent:
			nb.unfocusSlot()
			nb.selection.SetFirst(i)
		case cell.SelectLastCellEvent:
			nb.unfocusSlot()
			nb.selection.SetLast(i)
		case cell.FocusPlaceholder:
			nb.focusSlot(i + e.Offset)
		}
	}
	for i := range nb.slots {
		isActive := nb.activeSlot == i
		es := nb.slots[i].Event(isActive, gtx)
		for _, e := range es {
			if _, ok := e.(SelectSlotEvent); ok {
				nb.focusSlot(i)
			} else if ce, ok := e.(AddCellEvent); ok {
				nb.InsertCell(i, ce.Type)
				nb.focusCell(i)
			} else if _, ok := e.(FocusPreviousCellEvent); ok {
				nb.focusCell(i - 1)
			} else if _, ok := e.(FocusNextCellEvent); ok {
				nb.focusCell(i)
			} else if _, ok := e.(SelectPreviousCellEvent); ok {
				nb.unfocusSlot()
				nb.selection.SetFirst(i - 1)
			} else if _, ok := e.(SelectNextCellEvent); ok {
				nb.unfocusSlot()
				nb.selection.SetFirst(i)
			}
		}
	}
	es := nb.selection.Event(gtx)
	for _, e := range es {
		switch e := e.(type) {
		case DeleteSelected:
			nb.DeleteSelected()
		case FocusSlotEvent:
			nb.focusSlot(e.Index)
		}
	}
	return nil
}

func (nb *Notebook) Layout(gtx *layout.Context) {
	n := len(nb.cells)*2 + 1 // Their is one more Slot then cells
	nb.list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			i = i / 2
			isActive := nb.activeSlot == i && !nb.isOutputCell(i)
			isLast := i == len(nb.slots)-1
			nb.slots[i].Layout(isActive, isLast, gtx)
		} else {
			i := (i - 1) / 2
			isSelected := nb.selection.IsSelected(i)
			nb.cells[i].Layout(isSelected, gtx)
		}
	})
}

func (nb *Notebook) eval(i int) {
	c := nb.cells[i]
	textIn := c.Text()
	if textIn == "" {
		return
	}
	c.Label = fmt.Sprintf("In[%d]:= ", nb.promptCount)
	src := parser.ReplaceSyms(textIn)
	buf := bytes.NewBufferString(src)
	expOut, err := parser.InterpBuf(buf, "nofile", nb.kernel)
	expOut = nb.kernel.Eval(expOut)
	if nb.isOutputCell(i + 1) {
		nb.DeleteCell(i + i)
	}
	nb.InsertCell(i+1, cell.OutputCell)
	nb.cells[i+1].Out = expOut
	nb.cells[i+1].Err = err
	nb.cells[i+1].Label = fmt.Sprintf("Out[%d]= ", nb.promptCount)
	nb.promptCount++
	nb.focusSlot(i + 1)
}

func (nb *Notebook) isOutputCell(i int) bool {
	if i >= 0 && i < len(nb.cells) {
		return nb.cells[i].Type == cell.OutputCell
	}
	return false
}

func (nb *Notebook) focusCell(i int) {
	if i >= 0 && i < len(nb.cells) {
		nb.activeSlot = -1
		nb.selection.Clear()
		nb.cells[i].Focus()
	}
}

func (nb *Notebook) focusSlot(i int) {
	if i >= 0 && i < len(nb.slots) {
		nb.activeSlot = i
		nb.selection.Clear()
		nb.slots[i].Focus(true)
	}
}

func (nb *Notebook) unfocusSlot() {
	nb.activeSlot = -1
}

func (nb *Notebook) InsertCell(index int, typ cell.CellType) {
	nb.slots = append(nb.slots, NewSlot())
	cell := cell.NewCell(typ, "In[ ]:=", nb.styles)
	nb.cells = append(nb.cells, cell)
	copy(nb.cells[index+1:], nb.cells[index:])
	nb.cells[index] = cell
	nb.selection.Size = len(nb.cells)
}

func (nb *Notebook) DeleteCell(i int) {
	if i < len(nb.cells)-1 {
		copy(nb.cells[i:], nb.cells[i+1:])
	}
	nb.cells[len(nb.cells)-1] = nil // or the zero value of T
	nb.cells = nb.cells[:len(nb.cells)-1]
	nb.slots = nb.slots[:len(nb.cells)+1]
	nb.selection.Size = len(nb.cells)
}

func (nb *Notebook) DeleteSelected() {
	unselectedCount := 0
	for i, c := range nb.cells {
		if !nb.selection.IsSelected(i) {
			nb.cells[unselectedCount] = c
			unselectedCount++
		}
	}
	nb.cells = nb.cells[:unselectedCount]
	nb.slots = nb.slots[:unselectedCount+1]
	nb.selection.Size = unselectedCount
}

func ReadNotebookFile() (*Notebook, error) {
	return nil, nil
}

func WriteNotebookFile(path string, notebook *Notebook) error {
	content, _ := xml.MarshalIndent(notebook, "", "	")
	_ = ioutil.WriteFile(path, content, 0644)
	return nil
}

type SaveNotebookEvent struct{}
