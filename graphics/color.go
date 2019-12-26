package graphics

import (
	"errors"
	"gioui.org/f32"
	"gioui.org/op"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"image/color"
)

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
		return nil, errors.New("RGBColor[] should have 3 floats arguments")
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
	//c.color = rgbFromFlts(0.733, 0.733, 0.733)
	//c.color = util.LightGrey
	//c.color = color.RGBA{186, 186, 186, 255}
	//c.color = color.RGBA{127, 127, 127, 255}
	return c, nil
}

func rgbFromFlts(r, g, b float32) color.RGBA {
	return color.RGBA{
		A: 255,
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255)}
}
