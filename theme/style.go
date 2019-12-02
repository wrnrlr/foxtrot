package theme

import (
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/editor"
)

type Styles struct {
	Foxtrot, Title, Section, SubSection, SubSubSection, Text, Code editor.EditorStyle
	Theme                                                          *material.Theme
	shaper                                                         *text.Shaper
}
