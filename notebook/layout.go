package notebook

import . "gioui.org/layout"

func (nb *Notebook) Layout(gtx *Context) {
	n := len(nb.Cells)*2 + 1 // Their is one more Slot then Cells
	nb.list.Layout(gtx, n, func(i int) {
		if i%2 == 0 {
			i = i / 2
			isActive := nb.activeSlot == i // && !nb.isOutputCell(i)
			isLast := i == len(nb.slots)-1
			nb.slots[i].Layout(isActive, isLast, gtx)
		} else {
			i := (i - 1) / 2
			isSelected := nb.selection.IsSelected(i)
			nb.Cells[i].Layout(isSelected, gtx)
		}
	})
}
