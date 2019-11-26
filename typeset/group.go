package typeset

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
)

type Group struct {
	Parts                  []Shape
	Subscript, SuperScript *Shape
}

func (g *Group) Dimensions(c unit.Converter, s *text.Shaper, font text.Font) layout.Dimensions {
	return layout.Dimensions{}
}

func (g *Group) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {

}
