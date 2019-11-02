package foxtrot

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
)

var (
	lightGrey  = rgb(0xbbbbbb)
	lightPink  = rgb(0xffb6c1)
	lightBlue  = rgb(0x039be5)
	lightGreen = rgb(0x7cb342)
	white      = rgb(0xffffff)
	black      = rgb(0x000000)
	red        = rgb(0xe53935)
	blue       = rgb(0x1e88e5)

	inlineHeight     = layout.Constraint{Min: 50, Max: 50}
	promptWidth      = layout.Constraint{Min: 80, Max: 80}
	_defaultFontSize = unit.Sp(20)
	_promptFontSize  = unit.Sp(12)
	theme            *material.Theme
	_padding         = unit.Dp(8)
	promptTheme      *material.Theme
)

type Theme struct {
	Shaper *text.Shaper
	Color  struct {
		Primary color.RGBA
		Text    color.RGBA
		Hint    color.RGBA
	}
	TextSize unit.Value
}

func NewTheme() *Theme {
	t := &Theme{
		Shaper: font.Default(),
	}
	t.Color.Primary = rgb(0x3f51b5)
	t.Color.Text = rgb(0x000000)
	t.Color.Hint = rgb(0xbbbbbb)
	t.TextSize = unit.Sp(16)
	return t
}
func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

type Label struct {
	// Face defines the text style.
	Font text.Font
	// Color is the text color.
	Color color.RGBA
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	Text     string

	shaper *text.Shaper
}
