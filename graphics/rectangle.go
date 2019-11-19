package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
)

type Rectangle struct {
	min, max f32.Point
}

func (r Rectangle) Draw(ctx *context, ops *op.Ops) {

}

func (r Rectangle) BoundingBox() (bbox f32.Rectangle) {
	return bbox
}
