package graphics

import (
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/util"
	"image/color"
)

type Style struct {
	Shaper    *text.Shaper
	Font      text.Font
	TextColor color.RGBA
	TextSize  unit.Value

	StrokeWidth float32
	StrokeColor *color.RGBA

	Thickness float32
}

func NewStyle() *Style {
	st := &Style{
		Shaper: font.Default(),
	}
	st.Font = text.Font{Size: unit.Sp(20)}
	st.TextColor = util.Black
	st.StrokeColor = &util.Black
	st.TextSize = unit.Sp(20)
	return st
}

//StrokeWidth float32
//StrokeColor color.RGBA
//
//FontSize  float32
//FontColor color.RGBA
//Font      *text.Shaper
//
//FillColor color.RGBA
