package foxtrot

import (
	"gioui.org/op"
	"gioui.org/op/paint"
	"image/color"
)

func rgb(c uint32) color.RGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func colorMaterial(ops *op.Ops, color color.RGBA) op.MacroOp {
	var mat op.MacroOp
	mat.Record(ops)
	paint.ColorOp{Color: color}.Add(ops)
	mat.Stop()
	return mat
}
