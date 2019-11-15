package notebook

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/editor"
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

type Styles struct {
	Foxtrot, Title, Section, SubSection, SubSubSection, Text, Code editor.EditorStyle
	shaper                                                         *text.Shaper
}

func DefaultStyles() *Styles {
	styles := Styles{}
	shaper := font.Default()
	styles.Foxtrot = editor.EditorStyle{
		Font:       text.Font{Variant: "Mono", Size: unit.Sp(18)},
		Color:      black,
		CaretColor: black,
		Shaper:     shaper}
	styles.Title = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(38)},
		Color:      red,
		CaretColor: black,
		Shaper:     shaper}
	styles.Section = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(32)},
		Color:      red,
		CaretColor: black,
		Shaper:     shaper}
	styles.SubSection = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(26)},
		Color:      red,
		CaretColor: black,
		Shaper:     shaper}
	styles.SubSubSection = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(20)},
		Color:      red,
		CaretColor: black,
		Shaper:     shaper}
	styles.Text = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(16)},
		Color:      black,
		CaretColor: black,
		Shaper:     shaper}
	styles.Code = editor.EditorStyle{
		Font:       text.Font{Variant: "Mono", Size: unit.Sp(16)},
		Color:      black,
		CaretColor: black,
		Shaper:     shaper}
	return &styles
}

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
