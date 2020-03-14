package slot

import (
	"gioui.org/gesture"
	"gioui.org/layout"
	"gioui.org/widget"
	"time"
)

const (
	blinksPerSecond  = 1
	maxBlinkDuration = 10 * time.Second
)

type Slot interface {
	Index() int
	Layout(isLast bool, gtx *layout.Context)
}

// Every cell has a slot above and below it that allow new cells to be inserted insert.
type slot struct {
	plusButton *widget.Button
	blinkStart time.Time

	eventKey     int
	focused      bool
	caretOn      bool
	requestFocus bool

	clicker gesture.Click

	events     []interface{}
	prevEvents int
	index      int
	menu       menu
}

func NewSlot() Slot {
	return &slot{
		plusButton: new(widget.Button),
		blinkStart: time.Now()}
}

func (s slot) Index() int {
	return s.index
}

func (s slot) Focus(i int, requestFocus bool) {
	s.index = i
	if i > 0 {
		s.requestFocus = requestFocus
	}
}
