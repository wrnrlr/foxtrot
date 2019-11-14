package foxtrot

import (
	"fmt"
	"gioui.org/io/key"
	"gioui.org/layout"
)

// Select a range of cells from begin to end
type Selection struct {
	eventKey     int
	count        int
	events       []interface{}
	prevEvents   int
	focused      bool
	requestFocus bool

	begin, end int
}

func (s *Selection) Event(gtx *layout.Context) []interface{} {
	s.processEvents(gtx)
	return s.flushEvents()
}

func (s *Selection) flushEvents() []interface{} {
	events := s.events
	s.events = nil
	s.prevEvents = 0
	return events
}

func (s *Selection) processEvents(gtx *layout.Context) {
	key.InputOp{Key: &s.eventKey, Focus: s.requestFocus}.Add(gtx.Ops)
	s.requestFocus = false
	for _, e := range gtx.Events(&s.eventKey) {
		switch ke := e.(type) {
		case key.Event:
			fmt.Printf("Selection Key Event\n")
			if ke.Name == key.NameDeleteBackward || ke.Name == key.NameDeleteForward {
				s.events = append(s.events, DeleteSelected{})
			}
		case key.FocusEvent:
			fmt.Printf("Selection Focus Event: %b\n", ke.Focus)
			s.focused = ke.Focus
		}
	}
}

func (s *Selection) Clear() {
	s.requestFocus = false
	s.begin = -1
	s.end = -1
}

func (s *Selection) SetBegin(i int) {
	s.requestFocus = true
	s.begin = i
	s.end = i
}

func (s *Selection) SetEnd(i int) {
	s.end = i
}

func (s *Selection) IsSelected(i int) bool {
	return s.begin != -1 && s.min() >= i && s.max() <= i
}

func (s *Selection) min() int {
	if s.begin < s.end {
		return s.begin
	} else {
		return s.end
	}
}

func (s *Selection) max() int {
	if s.begin > s.end {
		return s.begin
	} else {
		return s.end
	}
}

type DeleteSelected struct{}
