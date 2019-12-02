package util

import (
	"gioui.org/f32"
	"image"
	"image/color"
)

var (
	LightGrey  = Rgb(0xbbbbbb)
	LightPink  = Rgb(0xffb6c1)
	LightBlue  = Rgb(0x039be5)
	LightGreen = Rgb(0x7cb342)

	White         = Rgb(0xffffff)
	Black         = Rgb(0x000000)
	Red           = Rgb(0xe53935)
	Blue          = Rgb(0x1e88e5)
	SelectedColor = Rgb(0xe1f5fe)
)

func Rgb(c uint32) color.RGBA {
	return Argb(0xff000000 | c)
}

func Argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func ToPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
