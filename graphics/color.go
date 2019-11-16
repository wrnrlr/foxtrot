package graphics

import (
	"image/color"
)

var (
	lightGrey     = rgb(0xbbbbbb)
	lightPink     = rgb(0xffb6c1)
	lightBlue     = rgb(0x039be5)
	lightGreen    = rgb(0x7cb342)
	white         = rgb(0xffffff)
	black         = rgb(0x000000)
	red           = rgb(0xe53935)
	blue          = rgb(0x1e88e5)
	selectedColor = rgb(0xe1f5fe)
)

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
