package slot

import (
	"fmt"
	"gioui.org/gesture"
	. "gioui.org/io/key"
	"gioui.org/io/pointer"
	. "gioui.org/layout"
	//. "github.com/wrnrlr/foxtrot/cell"
)

func (s *slot) Event(isActive bool, gtx *Context) []interface{} {
	s.processEvents(isActive, gtx)
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *slot) processEvents(isActive bool, gtx *Context) {
	//for s.plusButton.Clicked(gtx) {
	//s.events = append(s.events,
	//s.events = append(s.events, AddCell{Type: Input})
	//}
	//for s.backgroundButton.Clicked(gtx) {
	//	s.Focus(true)
	//	s.events = append(s.events, SelectSlot{})
	//}
	s.processKey(isActive, gtx)
}

func (s *slot) processPointer(gtx *Context) interface{} {
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

func (s *slot) processKey(isActive bool, gtx *Context) {
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
			fmt.Println("slot: key.EditEvent")
		}
	}
}

func (slot) switchKey(ke Event) interface{} {
	switch {
	case isKey(ke, NameUpArrow, ModShift):
		return SelectPrevCell
	case isKey(ke, NameDownArrow, ModShift):
		return SelectNextCell
	//case isKey(ke, NameReturn), isKey(ke, NameEnter):
	//	return ke
	case isKey(ke, NameUpArrow), isKey(ke, NameLeftArrow):
		return FocusPrevCell
	case isKey(ke, NameDownArrow), isKey(ke, NameRightArrow):
		return FocusNextCell
	case isKey(ke, "1", ModCommand):
		return AddCell{"Header1", ""}
	case isKey(ke, "2", ModCommand):
		return AddCell{"Header2", ""}
	case isKey(ke, "3", ModCommand):
		return AddCell{"Header3", ""}
	case isKey(ke, "4", ModCommand):
		return AddCell{"Header4", ""}
	case isKey(ke, "5", ModCommand):
		return AddCell{"Paragraph", ""}
	case isKey(ke, "6", ModCommand):
		return AddCell{"Code", ""}
	default:
		return nil
	}
}

type SlotEvent interface{ SlotEvent() }

type FocusCell struct{ Delta int }

func (SelectCell) SlotEvent() {}

var FocusPrevCell = FocusCell{Delta: -1}
var FocusNextCell = FocusCell{Delta: 1}

type SelectCell struct{ Delta int }

func (FocusCell) SlotEvent() {}

var SelectPrevCell = FocusCell{Delta: -1}
var SelectNextCell = FocusCell{Delta: 1}

type AddCell struct{ Type, Content string }

func (AddCell) SlotEvent() {}

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
