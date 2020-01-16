package graphics

import (
	"gioui.org/f32"
	"github.com/wrnrlr/foxtrot/util"
)

//type Context struct {
//	*layout.Context
//	Shaper *text.Shaper
//}

type context struct {
	//
	BBox  f32.Rectangle
	style *Style
}

func (c context) width() float32 {
	return c.x(c.BBox.Max.X)
}

func (c context) height() float32 {
	return c.y(c.BBox.Max.Y)
}

func (c context) x(x float32) float32 {
	if c.BBox.Min.X <= 0 {
		return x + util.Absf32(c.BBox.Min.X)
	} else {
		return x - c.BBox.Min.X
	}
}

func (c context) y(y float32) float32 {
	if c.BBox.Min.Y <= 0 {
		return y + util.Absf32(c.BBox.Min.Y)
	} else {
		return y - c.BBox.Min.Y
	}
}

func (c context) transformPoint(p f32.Point) f32.Point {
	return f32.Point{X: c.x(p.X), Y: c.y(p.Y)}
}
