package slot

import (
	"gioui.org/gesture"
	"gioui.org/widget"
	"time"
)

const (
	blinksPerSecond  = 1
	maxBlinkDuration = 10 * time.Second
)

// Every cell has a slot above and below it that allow new Cells to be inserted insert.
type Slot struct {
	Active           bool
	plusButton       *widget.Button
	backgroundButton *widget.Button

	eventKey     int
	blinkStart   time.Time
	focused      bool
	caretOn      bool
	requestFocus bool

	clicker gesture.Click

	events     []interface{}
	prevEvents int
}

func NewSlot() *Slot {
	return &Slot{
		plusButton:       new(widget.Button),
		backgroundButton: new(widget.Button),
		blinkStart:       time.Now()}
}

func (s *Slot) Focus(requestFocus bool) {
	s.requestFocus = requestFocus
}
