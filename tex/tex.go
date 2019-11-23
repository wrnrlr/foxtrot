package tex

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
)

type Shape interface {
	Dimensions(c unit.Converter, s *text.Shaper, font text.Font) layout.Dimensions
	Layout(gtx *layout.Context, s *text.Shaper, font text.Font)
}

func scaleDownFont(font text.Font) text.Font {
	return text.Font{
		Typeface: font.Typeface,
		Variant:  font.Variant,
		Size:     font.Size.Scale(0.8),
		Style:    font.Style,
		Weight:   font.Weight}
}

type Sqrt struct{}

type Group struct {
	Parts                  []Shape
	Subscript, SuperScript *Shape
}
