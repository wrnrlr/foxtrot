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
	cells        []*Cell
	placeholders []*Placeholder
	kernel       *expreduce.EvalState
	promptCount  int

	activePlaceholder int
	list              layout.List
	selection         *Selection
}

func NewNotebook() *Notebook {
	cells := make([]*Cell, 0)
	kernel := expreduce.NewEvalState()
	firstPlaceholder := NewPlaceholder()
	//firstPlaceholder.Focus()
	adds := []*Placeholder{firstPlaceholder}
	selection := &Selection{}
	return &Notebook{cells, adds, kernel, 1, 0, layout.List{Axis: layout.Vertical}, selection}
}

func (nb *Notebook) Event(gtx *layout.Context) interface{} {
	nb.selection.Reset()
	for i, c := range nb.cells {
		e := c.Event(gtx)
		switch ce := e.(type) {
		case EvalEvent:
			nb.eval(i)
			nb.focusPlaceholder(i+1, gtx)
		case SelectCellEvent:
			fmt.Println("Notebook: Select Cell")
			nb.unfocusPlaceholder()
			nb.selection.RequestFocus(ce.selected, gtx)
		}
	}
	for i := range nb.placeholders {
		isActive := nb.activePlaceholder == i
		es := nb.placeholders[i].Event(isActive, gtx)
		for _, e := range es {
			if _, ok := e.(SelectPlaceholderEvent); ok {
				fmt.Printf("Notebook: select placeholder: %d\n", i)
				nb.focusPlaceholder(i, gtx)
			} else if ce, ok := e.(AddCellEvent); ok {
				nb.InsertCell(i, ce.Type)
				nb.focusCell(i)
			} else if _, ok := e.(FocusPreviousCellEvent); ok {
				nb.focusCell(i - 1)
			} else if _, ok := e.(FocusNextCellEvent); ok {
				nb.focusCell(i)
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
	n := len(nb.cells)*2 + 1 // Their is one more Placeholder then cells
	nb.list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			i = i / 2
			isActive := nb.activePlaceholder == i
			nb.placeholders[i].Layout(isActive, gtx)
		} else {
			nb.cells[(i-1)/2].Layout(gtx)
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
	c.out = NewOut(expOut)
	c.out.SetState(nb.kernel, nb.promptCount)
	c.promptNum = nb.promptCount
	nb.promptCount++
}

func (nb *Notebook) focusCell(i int) {
	if i >= 0 && i < len(nb.cells) {
		fmt.Printf("Focus cell %d\n", i)
		nb.activePlaceholder = -1
		nb.cells[i].Focus()
	}
}

func (nb *Notebook) focusPlaceholder(i int, gtx *layout.Context) {
	if i >= 0 && i < len(nb.placeholders) {
		fmt.Printf("Focus Placeholder %d\n", i)
		nb.activePlaceholder = i
		nb.placeholders[i].Focus(true, gtx)
	}
}

func (nb *Notebook) unfocusPlaceholder() {
	//if nb.activePlaceholder > -1 {
	nb.activePlaceholder = -1
	//}
}

func (nb *Notebook) InsertCell(index int, typ CellType) {
	nb.placeholders = append(nb.placeholders, NewPlaceholder())
	cell := NewCell(typ, -1)
	nb.cells = append(nb.cells, cell)
	copy(nb.cells[index+1:], nb.cells[index:])
	nb.cells[index] = cell
}

func (nb *Notebook) DeleteSelected() {
	i := 0
	for _, c := range nb.cells {
		if !c.IsSelected() {
			nb.cells[i] = c
			i++
		}
	}
	nb.cells = nb.cells[:i]
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
