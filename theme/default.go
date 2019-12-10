package theme

import (
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/editor"
	"github.com/wrnrlr/foxtrot/util"
)

func DefaultStyles() *Styles {
	styles := Styles{}
	shaper := font.Default()
	styles.Foxtrot = editor.EditorStyle{
		Font:       text.Font{Variant: "Mono", Size: unit.Sp(18)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.H1 = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(38)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.H2 = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(32)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.H3 = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(26)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.H4 = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(20)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.Text = editor.EditorStyle{
		Font:       text.Font{Size: unit.Sp(16)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.Code = editor.EditorStyle{
		Font:       text.Font{Variant: "Mono", Size: unit.Sp(16)},
		Color:      util.Black,
		CaretColor: util.Black,
		Shaper:     shaper}
	styles.Theme = material.NewTheme()
	styles.Theme.Shaper = shaper
	styles.Theme.Color.Text = util.LightGrey
	styles.Theme.TextSize = unit.Sp(16)
	return &styles
}
