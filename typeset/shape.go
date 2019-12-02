package typeset

import (
	"gioui.org/layout"
	"gioui.org/text"
)

type Shape interface {
	Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions
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
