package notebook

import (
	. "gioui.org/layout"
	. "github.com/wrnrlr/foxtrot/cell"
)

func (nb *Notebook) Event(gtx *Context) interface{} {
	nb.cellEvents(gtx)
	nb.slotEvents(gtx)
	nb.selectionEvent(gtx)
	return nil
}

func (nb *Notebook) cellEvents(gtx *Context) {
	for i, c := range nb.Cells {
		e := c.Event(gtx)
		switch e := e.(type) {
		case EvalEvent:
			nb.eval(i)
		case SelectFirstCellEvent:
			nb.unfocusSlot()
			nb.selection.SetFirst(i)
		case SelectLastCellEvent:
			nb.unfocusSlot()
			nb.selection.SetLast(i)
		case FocusPlaceholder:
			nb.focusSlot(i + e.Offset)
		}
	}
}

func (nb *Notebook) slotEvents(gtx *Context) {
	for i := range nb.slots {
		isActive := nb.activeSlot == i
		es := nb.slots[i].Event(isActive, gtx)
		for _, e := range es {
			switch ev := e.(type) {
			case SelectSlot:
				nb.focusSlot(i)
			case AddCell:
				nb.InsertCell(i, ev.Type)
				nb.focusCell(i)
			case FocusPreviousCell:
				if nb.isOutputCell(i - 1) {
					i -= 1
				}
				nb.focusCell(i - 1)
			case FocusNextCell:
				if nb.isOutputCell(i) {
					i += 1
				}
				nb.focusCell(i)
			case SelectPreviousCell:
				nb.unfocusSlot()
				nb.selection.SetFirst(i - 1)
			case SelectNextCell:
				nb.unfocusSlot()
				nb.selection.SetFirst(i)
			}
		}
	}
}

func (nb *Notebook) selectionEvent(gtx *Context) {
	es := nb.selection.Event(gtx)
	for _, e := range es {
		switch e := e.(type) {
		case DeleteSelected:
			nb.DeleteSelected()
		case FocusSlotEvent:
			nb.focusSlot(e.Index)
		}
	}
}
