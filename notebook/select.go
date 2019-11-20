package notebook

import (
	"gioui.org/io/key"
	"gioui.org/layout"
)

// Select a range of cells from first to last
type Selection struct {
	eventKey     int
	count        int
	events       []interface{}
	prevEvents   int
	focused      bool
	requestFocus bool

	first, last int
}

func NewSelection() *Selection {
	return &Selection{first: -1, last: -1}
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
			if ke.Name == key.NameDeleteBackward || ke.Name == key.NameDeleteForward {
				s.events = append(s.events, DeleteSelected{})
			}
		case key.FocusEvent:
			s.focused = ke.Focus
		}
	}
}

func (s *Selection) Clear() {
	s.requestFocus = false
	s.first = -1
	s.last = -1
}

func (s *Selection) SetFirst(i int) {
	s.requestFocus = true
	s.first = i
	s.last = i
}

func (s *Selection) SetLast(i int) {
	s.requestFocus = true
	if s.first == -1 {
		s.first = i
	}
	s.last = i
}

func (s *Selection) IsSelected(i int) bool {
	return s.first != -1 && i >= s.min() && i <= s.max()
}

func (s *Selection) min() int {
	if s.first < s.last {
		return s.first
	} else {
		return s.last
	}
}

func (s *Selection) max() int {
	if s.first > s.last {
		return s.first
	} else {
		return s.last
	}
}

type DeleteSelected struct{}
