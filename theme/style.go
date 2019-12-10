package theme

import (
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/editor"
)

type Styles struct {
	Foxtrot, H1, H2, H3, H4, H5, H6, Text, Code editor.EditorStyle
	Theme                                       *material.Theme
	shaper                                      *text.Shaper
}
