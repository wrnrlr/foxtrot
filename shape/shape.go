package shape

import (
	"gioui.org/op"
	"image/color"
)

type Shape interface {
	Add(ops *op.Ops)
}

type Stroke interface {
	Shape
	Stroke(width float32, style StrokeType, rgba color.RGBA, ops *op.Ops)
}

type Fill interface {
	Stroke
	Fill(rgba color.RGBA, ops *op.Ops)
}

type StrokeType int

const (
	Solid StrokeType = iota
	Dotted
	Dashed
)
