package graphics

import (
	"errors"
	"gioui.org/f32"
	"gioui.org/op"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"image/color"
)

var (
	black = rgb(0x000000)
)

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

type RGBColor struct {
	color     color.RGBA
	thickness float32
}

func (c RGBColor) Draw(ctx *context, ops *op.Ops) {
	*ctx.style.StrokeColor = c.color
}

func (c RGBColor) BoundingBox() (bbox f32.Rectangle) {
	return bbox
}

func toColor(e *atoms.Expression) (*RGBColor, error) {
	if len(e.Parts) != 4 {
		return nil, errors.New("Color[] should have 3 floats arguments")
	}
	r, err := toFloat(e.GetPart(1))
	if err != nil {
		return nil, err
	}
	g, err := toFloat(e.GetPart(2))
	if err != nil {
		return nil, err
	}
	b, err := toFloat(e.GetPart(3))
	if err != nil {
		return nil, err
	}
	c := &RGBColor{color: rgbFromFlts(r, g, b)}
	return c, nil
}

func rgbFromFlts(r, g, b float32) color.RGBA {
	return color.RGBA{
		A: 255,
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255)}
}
