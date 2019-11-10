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
	lightGrey     = rgb(0xbbbbbb)
	lightPink     = rgb(0xffb6c1)
	lightBlue     = rgb(0x039be5)
	lightGreen    = rgb(0x7cb342)
	white         = rgb(0xffffff)
	black         = rgb(0x000000)
	red           = rgb(0xe53935)
	blue          = rgb(0x1e88e5)
	selectedColor = rgb(0xe1f5fe)

	inlineHeight     = layout.Constraint{Min: 50, Max: 50}
	promptWidth      = unit.Sp(50)
	cellLeftMargin   = unit.Sp(20)
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

func initThemes() {
	FoxtrotTheme = material.NewTheme()
	TitleTheme = material.NewTheme()
	TitleTheme.TextSize = unit.Sp(38)
	TitleTheme.Color.Text = red
	SectionTheme = material.NewTheme()
	SectionTheme.TextSize = unit.Sp(32)
	SectionTheme.Color.Text = red
	SubSectionTheme = material.NewTheme()
	SubSectionTheme.TextSize = unit.Sp(26)
	SubSectionTheme.Color.Text = red
	SubSubSectionTheme = material.NewTheme()
	SubSubSectionTheme.TextSize = unit.Sp(20)
	SubSubSectionTheme.Color.Text = red
	TextTheme = material.NewTheme()
	TextTheme.TextSize = unit.Sp(16)
	TextTheme.Color.Text = black
	CodeTheme = material.NewTheme()
	CodeTheme.TextSize = unit.Sp(16)
	CodeTheme.Color.Text = black
}

var (
	FoxtrotTheme, TitleTheme, SectionTheme, SubSectionTheme, SubSubSectionTheme, TextTheme, CodeTheme *material.Theme
	TitleEditor                                                                                       *material.Editor
)

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
