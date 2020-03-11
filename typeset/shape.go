package typeset

import (
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/wrnrlr/foxtrot/style"
)

type Shape interface {
	Dimensions(gtx *layout.Context, s style.Style) layout.Dimensions
	Layout(gtx *layout.Context, s style.Style)
}

func scaleDownFont(font text.Font) text.Font {
	return text.Font{
		Typeface: font.Typeface,
		Variant:  font.Variant,
		Size:     font.Size.Scale(0.8),
		Style:    font.Style,
		Weight:   font.Weight}
}
