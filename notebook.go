package foxtrot

import (
	"encoding/xml"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/parser"
	"io/ioutil"
)

type Notebook struct {
	cells        []*Cell
	add          *Add
	placeholders []placeholder
	kernel       *expreduce.EvalState
	promptCount  int
}

func NewNotebook() *Notebook {
	cells := make([]*Cell, 0)
	add := NewAdd()
	kernel := expreduce.NewEvalState()
	placeholders := []placeholder{newPlaceholder()}
	return &Notebook{cells, add, placeholders, kernel, 1}
}

func (nb *Notebook) Event(gtx *layout.Context) interface{} {
	for i, c := range nb.cells {
		e := c.Event(gtx)
		if _, ok := e.(EvalEvent); ok {
			nb.eval(i)
			nb.focusNext(i)
		}
	}
	for i, p := range nb.placeholders {
		e := p.Event(gtx)
		if _, ok := e.(FocusPlaceholder); ok {
			nb.add.Focus(i)
		}
	}
	e := nb.add.Event(gtx)
	if _, ok := e.(AddCellEvent); ok {
		nb.InsertCell(nb.add.Index)
	}
	return nil
}

func (nb *Notebook) Layout(gtx *layout.Context) {
	n := len(nb.cells)*2 + 1
	list := layout.List{Axis: layout.Vertical}
	list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			i = i / 2
			if nb.add.IsIndex(i) {
				nb.add.Layout(i, gtx)
			} else {
				nb.placeholders[i].Layout(gtx)
			}
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

func (nb *Notebook) focusNext(i int) {
	if i < len(nb.cells)-1 {
		nb.cells[i+1].Focus()
	} else {
		nb.add.Focus(len(nb.cells))
	}
}

func (nb *Notebook) InsertCell(index int) {
	nb.placeholders = append(nb.placeholders, newPlaceholder())
	cell := newCell(-1)
	nb.cells = append(nb.cells, newCell(-1))
	copy(nb.cells[index+1:], nb.cells[index:])
	nb.cells[index] = cell
	nb.add.Focus(-1)
	cell.Focus()
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
