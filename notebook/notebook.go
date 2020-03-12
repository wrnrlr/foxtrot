package notebook

import (
	"encoding/xml"
	. "gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/theme"
	"io/ioutil"
)

type Notebook struct {
	Cells       cell.Cells
	slots       []*Slot
	kernel      *expreduce.EvalState
	promptCount int

	activeSlot int
	list       List
	selection  *Selection
	styles     *theme.Styles
}

func NewNotebook() *Notebook {
	kernel := expreduce.NewEvalState()
	firstSlot := NewSlot()
	adds := []*Slot{firstSlot}
	selection := NewSelection()
	styles := theme.DefaultStyles()
	return &Notebook{nil, adds, kernel, 1, 0, List{Axis: Vertical}, selection, styles}
}

func (nb *Notebook) isOutputCell(i int) bool {
	if i >= 0 && i < len(nb.Cells) {
		return nb.Cells[i].Type() == cell.Output
	}
	return false
}

func (nb *Notebook) focusCell(i int) {
	if i >= 0 && i < len(nb.Cells) {
		nb.activeSlot = -1
		nb.selection.Clear()
		nb.Cells[i].Focus()
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

func (nb *Notebook) AddCells(cells cell.Cells) {
	for _, c := range cells {
		nb.slots = append(nb.slots, NewSlot())
		nb.Cells = append(nb.Cells, c)
	}
	nb.selection.Size = len(nb.Cells)
}

func (nb *Notebook) InsertCell(index int, typ cell.Type) {
	nb.slots = append(nb.slots, NewSlot())
	c := cell.NewCell(typ, "Content[ ]:=", nb.styles)
	nb.Cells = append(nb.Cells, c)
	copy(nb.Cells[index+1:], nb.Cells[index:])
	nb.Cells[index] = c
	nb.selection.Size = len(nb.Cells)
}

func (nb *Notebook) DeleteCell(i int) {
	if i < len(nb.Cells)-1 {
		copy(nb.Cells[i:], nb.Cells[i+1:])
	}
	nb.Cells[len(nb.Cells)-1] = nil // or the zero value of T
	nb.Cells = nb.Cells[:len(nb.Cells)-1]
	nb.slots = nb.slots[:len(nb.Cells)+1]
	nb.selection.Size = len(nb.Cells)
}

func (nb *Notebook) DeleteSelected() {
	unselectedCount := 0
	for i, c := range nb.Cells {
		if !nb.selection.IsSelected(i) {
			nb.Cells[unselectedCount] = c
			unselectedCount++
		}
	}
	nb.Cells = nb.Cells[:unselectedCount]
	nb.slots = nb.slots[:unselectedCount+1]
	nb.selection.Size = unselectedCount
}

func (nb *Notebook) Size() int {
	return len(nb.Cells)
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
