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
	placeholders []Placeholder
	kernel       *expreduce.EvalState
	promptCount  int
}

func NewNotebook() *Notebook {
	cells := make([]*Cell, 0)
	kernel := expreduce.NewEvalState()
	adds := []Placeholder{*NewPlaceholder()}
	return &Notebook{cells, adds, kernel, 1}
}

func (nb *Notebook) Event(gtx *layout.Context) interface{} {
	for i, c := range nb.cells {
		e := c.Event(gtx)
		if _, ok := e.(EvalEvent); ok {
			nb.eval(i)
			nb.focusPlaceholder(i + 1)
		}
	}
	for i, _ := range nb.placeholders {
		e := nb.placeholders[i].Event(gtx)
		if ce, ok := e.(AddCellEvent); ok {
			nb.InsertCell(i, ce.Type)
			nb.focusCell(i)
		} else if _, ok := e.(FocusPreviousCellEvent); ok {
			nb.focusCell(i - 1)
		} else if _, ok := e.(FocusNextCellEvent); ok {
			nb.focusCell(i)
		}
	}
	return nil
}

func (nb *Notebook) Layout(gtx *layout.Context) {
	n := len(nb.cells)*2 + 1 // Their is one more Placeholder then cells
	list := layout.List{Axis: layout.Vertical}
	list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			nb.placeholders[i/2].Layout(i, gtx)
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
		nb.cells[i].Focus()
	}
}

func (nb *Notebook) focusPlaceholder(i int) {
	if i > 0 && i < len(nb.placeholders) {
		fmt.Printf("Focus Placeholder %d\n", i)
		nb.placeholders[i].Focus()
	}
}

func (nb *Notebook) InsertCell(index int, typ CellType) {
	nb.placeholders = append(nb.placeholders, *NewPlaceholder())
	cell := NewCell(typ, -1)
	nb.cells = append(nb.cells, cell)
	copy(nb.cells[index+1:], nb.cells[index:])
	nb.cells[index] = cell
}

func (nb *Notebook) DeleteCell(index int) {}

func ReadNotebookFile() (*Notebook, error) {
	return nil, nil
}

func WriteNotebookFile(path string, notebook *Notebook) error {
	content, _ := xml.MarshalIndent(notebook, "", "	")
	_ = ioutil.WriteFile(path, content, 0644)
	return nil
}
