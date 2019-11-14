package foxtrot

import (
	"encoding/xml"
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/parser"
	"io/ioutil"
)

type Notebook struct {
	cells       []*Cell
	slots       []*Slot
	kernel      *expreduce.EvalState
	promptCount int

	activeSlot int
	list       layout.List
	selection  *Selection
}

func NewNotebook() *Notebook {
	cells := make([]*Cell, 0)
	kernel := expreduce.NewEvalState()
	firstSlot := NewSlot()
	adds := []*Slot{firstSlot}
	selection := &Selection{}
	return &Notebook{cells, adds, kernel, 1, 0, layout.List{Axis: layout.Vertical}, selection}
}

func (nb *Notebook) Event(gtx *layout.Context) interface{} {
	for i, c := range nb.cells {
		e := c.Event(gtx)
		switch e.(type) {
		case EvalEvent:
			nb.eval(i)
			nb.focusSlot(i+1, gtx)
		case SelectCellEvent:
			fmt.Println("Notebook: Select Cell")
			nb.unfocusSlot()
			nb.selection.SetBegin(i)
		}
	}
	for i := range nb.slots {
		isActive := nb.activeSlot == i
		es := nb.slots[i].Event(isActive, gtx)
		for _, e := range es {
			if _, ok := e.(SelectSlotEvent); ok {
				fmt.Printf("Notebook: select placeholder: %d\n", i)
				nb.focusSlot(i, gtx)
			} else if ce, ok := e.(AddCellEvent); ok {
				nb.InsertCell(i, ce.Type)
				nb.focusCell(i)
			} else if _, ok := e.(FocusPreviousCellEvent); ok {
				nb.focusCell(i - 1)
			} else if _, ok := e.(FocusNextCellEvent); ok {
				nb.focusCell(i)
			} else if _, ok := e.(SelectPreviousCellEvent); ok {
				nb.unfocusSlot()
				nb.selection.SetBegin(i - 1)
			} else if _, ok := e.(SelectNextCellEvent); ok {
				nb.unfocusSlot()
				nb.selection.SetBegin(i)
			}
		}
	}
	es := nb.selection.Event(gtx)
	for _, e := range es {
		if _, ok := e.(DeleteSelected); ok {
			fmt.Println("Delete Selected")
			nb.DeleteSelected()
		}
	}
	return nil
}

func (nb *Notebook) Layout(gtx *layout.Context) {
	n := len(nb.cells)*2 + 1 // Their is one more Slot then cells
	nb.list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			i = i / 2
			isActive := nb.activeSlot == i
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
	textIn := c.inEditor.Text()
	if textIn == "" {
		return
	}
	expIn := parser.Interp(textIn, nb.kernel)
	expOut := nb.kernel.Eval(expIn)
	c.out = &Out{expOut}
	c.promptNum = nb.promptCount
	nb.promptCount++
}

func (nb *Notebook) focusCell(i int) {
	if i >= 0 && i < len(nb.cells) {
		fmt.Printf("Focus cell %d\n", i)
		nb.activeSlot = -1
		nb.selection.Clear()
		nb.cells[i].Focus()
	}
}

func (nb *Notebook) focusSlot(i int, gtx *layout.Context) {
	if i >= 0 && i < len(nb.slots) {
		fmt.Printf("Focus Slot %d\n", i)
		nb.activeSlot = i
		nb.selection.Clear()
		nb.slots[i].Focus(true, gtx)
	}
}

func (nb *Notebook) unfocusSlot() {
	nb.activeSlot = -1
}

func (nb *Notebook) InsertCell(index int, typ CellType) {
	nb.slots = append(nb.slots, NewSlot())
	cell := NewCell(typ, -1)
	nb.cells = append(nb.cells, cell)
	copy(nb.cells[index+1:], nb.cells[index:])
	nb.cells[index] = cell
}

func (nb *Notebook) DeleteSelected() {
	i := 0
	for _, c := range nb.cells {
		if !nb.selection.IsSelected(i) {
			nb.cells[i] = c
			i++
		}
	}
	nb.cells = nb.cells[:i]
	nb.slots = nb.slots[:i+1]
}

func (nb *Notebook) DeleteCell(i int) {
	l := len(nb.cells) - 1
	if i < l {
		copy(nb.cells[i:], nb.cells[i+1:])
	}
	nb.cells[l] = nil // or the zero value of T
	nb.cells = nb.cells[:l-1]
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
