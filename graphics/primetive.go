package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
)

type Primitive interface {
	Draw(ctx *context, ops *op.Ops)
	BoundingBox() (bbox f32.Rectangle)
}
