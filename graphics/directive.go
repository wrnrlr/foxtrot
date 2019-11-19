package graphics

import "image/color"

type Directive interface {
	Set(style *Style)
}

type RGBColor struct {
	color     color.RGBA
	thickness float32
}

func (c RGBColor) Set(style *Style) {
	style.Color = c.color
}

type Thickness struct {
	thickness float32
}

func (t Thickness) Set(style *Style) {
	style.Thickness = t.thickness
}

type CMYKColor struct{}
