package slot

import (
	"fmt"
	"gioui.org/gesture"
	. "gioui.org/io/key"
	"gioui.org/io/pointer"
	. "gioui.org/layout"
	. "github.com/wrnrlr/foxtrot/cell"
)

func (s *Slot) Event(isActive bool, gtx *Context) []interface{} {
	s.processEvents(isActive, gtx)
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *Slot) processEvents(isActive bool, gtx *Context) {
	for s.plusButton.Clicked(gtx) {
		s.events = append(s.events, AddCell{Type: Input})
	}
	for s.backgroundButton.Clicked(gtx) {
		s.Focus(true)
		s.events = append(s.events, SelectSlot{})
	}
	s.processKey(isActive, gtx)
}

func (s *Slot) processPointer(gtx *Context) interface{} {
	if !s.focused {
		return nil
	}
	for _, evt := range s.clicker.Events(gtx) {
		switch {
		case evt.Type == gesture.TypePress && evt.Source == pointer.Mouse,
			evt.Type == gesture.TypeClick && evt.Source == pointer.Touch:
			s.blinkStart = gtx.Now()
			s.requestFocus = true
		}
	}
	return nil
}

func (s *Slot) processKey(isActive bool, gtx *Context) {
	if !isActive {
		return
	}
	for _, ke := range gtx.Events(&s.eventKey) {
		s.blinkStart = gtx.Now()
		switch ke := ke.(type) {
		case FocusEvent:
			s.focused = ke.Focus
		case Event:
			if !s.focused {
				break
			}
			e := s.switchKey(ke)
			if e != nil {
				s.events = append(s.events, e)
			}
		case EditEvent:
			fmt.Println("Slot: key.EditEvent")
		}
	}
}

func (Slot) switchKey(ke Event) interface{} {
	switch {
	case isKey(ke, NameUpArrow, ModShift):
		return SelectPreviousCell{}
	case isKey(ke, NameDownArrow, ModShift):
		return SelectNextCell{}
	case isKey(ke, NameReturn), isKey(ke, NameEnter):
		return AddCell{Type: Input}
	case isKey(ke, NameUpArrow), isKey(ke, NameLeftArrow):
		return FocusPreviousCell{}
	case isKey(ke, NameDownArrow), isKey(ke, NameRightArrow):
		return FocusNextCell{}
	case isKey(ke, "1", ModCommand):
		return AddCell{Type: H1}
	case isKey(ke, "2", ModCommand):
		return AddCell{Type: H2}
	case isKey(ke, "3", ModCommand):
		return AddCell{Type: H3}
	case isKey(ke, "4", ModCommand):
		return AddCell{Type: H4}
	case isKey(ke, "5", ModCommand):
		return AddCell{Type: Paragraph}
	case isKey(ke, "6", ModCommand):
		return AddCell{Type: Code}
	default:
		return nil
	}
}

type AddCell struct{ Type Type }
type SelectSlot struct{}
type FocusNextCell struct{}
type FocusPreviousCell struct{}
type SelectNextCell struct{}
type SelectPreviousCell struct{}

func isKey(e Event, k string, ms ...Modifiers) bool {
	return e.Name == k && hasMod(e, ms)
}

func hasMod(e Event, ms []Modifiers) bool {
	for _, m := range ms {
		if !e.Modifiers.Contain(m) {
			return false
		}
	}
	return true
}
