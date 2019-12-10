package graphics

import (
	"gioui.org/f32"
)

//type Context struct {
//	*layout.Context
//	Shaper *text.Shaper
//}

type context struct {
	BBox  f32.Rectangle
	style *Style
}

func (c context) x() float32 {
	return 0
}

func (c context) y() float32 {
	return 0
}
