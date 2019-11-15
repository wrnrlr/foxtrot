package graphics

import (
	"gioui.org/layout"
	"gioui.org/text"
)

type Context struct {
	*layout.Context
	Shaper *text.Shaper
}
