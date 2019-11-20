package notebook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSelection(t *testing.T) {
	s := NewSelection()
	assert.False(t, s.IsSelected(0))
}

func TestSetFirst(t *testing.T) {
	s := NewSelection()
	s.SetFirst(1)
	assert.False(t, s.IsSelected(0))
	assert.True(t, s.IsSelected(1))
	assert.False(t, s.IsSelected(2))
}

func TestSetLast(t *testing.T) {
	s := NewSelection()
	s.SetFirst(1)
	s.SetLast(2)
	assert.False(t, s.IsSelected(0))
	assert.True(t, s.IsSelected(1))
	assert.True(t, s.IsSelected(2))
	assert.False(t, s.IsSelected(3))
}

func TestClear(t *testing.T) {
	s := NewSelection()
	s.SetFirst(1)
	s.SetLast(2)
	s.Clear()
	assert.False(t, s.IsSelected(0))
	assert.False(t, s.IsSelected(1))
	assert.False(t, s.IsSelected(2))
	assert.False(t, s.IsSelected(3))
}
