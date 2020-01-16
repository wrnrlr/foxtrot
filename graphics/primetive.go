package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
)

type Primitive interface {
	Draw(ctx *context, gtx *layout.Context)
	BoundingBox() (bbox f32.Rectangle)
}
