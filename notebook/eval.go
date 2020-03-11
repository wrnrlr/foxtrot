package notebook

import (
	"bytes"
	"fmt"
	"github.com/corywalker/expreduce/expreduce/parser"
	"github.com/wrnrlr/foxtrot/cell"
)

func (nb *Notebook) eval(i int) {
	c := nb.Cells[i]
	textIn := c.Text()
	if textIn == "" {
		return
	}
	c.SetLabel(fmt.Sprintf("Content[%d]:= ", nb.promptCount))
	src := parser.ReplaceSyms(textIn)
	buf := bytes.NewBufferString(src)
	expOut, err := parser.InterpBuf(buf, "nofile", nb.kernel)
	expOut = nb.kernel.Eval(expOut)
	if nb.isOutputCell(i + 1) {
		nb.DeleteCell(i + i)
	}
	nb.InsertCell(i+1, cell.Output)
	nb.Cells[i+1].SetOut(expOut)
	nb.Cells[i+1].SetErr(err)
	nb.Cells[i+1].SetLabel(fmt.Sprintf("Out[%d]= ", nb.promptCount))
	nb.promptCount++
	nb.focusSlot(i + 1)
}
