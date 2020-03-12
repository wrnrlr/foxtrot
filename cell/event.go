package cell

import (
	"gioui.org/layout"
	"github.com/wrnrlr/foxtrot/editor"
)

func (c cell) Event(gtx *layout.Context) interface{} {
	for _, e := range c.input.Events(gtx) {
		switch e.(type) {
		case editor.SubmitEvent:
			return EvalEvent{}
		case editor.UpEvent:
			return FocusPlaceholder{Offset: 0}
		case editor.DownEvent:
			return FocusPlaceholder{Offset: 1}
		}
	}
	return c.margin.Event(gtx)
}

type FocusPlaceholder struct{ Offset int }
type EvalEvent struct{}
type SelectFirstCellEvent struct{}
type SelectLastCellEvent struct{}
