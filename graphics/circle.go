package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/shape"
)

func toCircle(e *atoms.Expression) (*Circle, error) {
	var err error
	center := f32.Point{}
	radius := float32(1)
	length := e.Len()
	if length == 1 {
		center, err = toPoint(e.GetPart(1))
		if err != nil {
			return nil, err
		}
	} else if length == 2 {
		center, err = toPoint(e.GetPart(1))
		if err != nil {
			return nil, err
		}
		radius, err = toFloat(e.GetPart(2))
		if err != nil {
			return nil, err
		}
	}
	c := &Circle{center, radius}
	return c, nil
}

type Circle struct {
	center f32.Point
	radius float32
}

func (c Circle) Draw(ctx *context, gtx *layout.Context) {
	center := ctx.transformPoint(c.center)
	radius := c.radius * float32(gtx.Px(unit.Sp(50)))
	rgba := *ctx.style.StrokeColor
	shape.Circle{center, radius}.Stroke(rgba, ctx.style.Thickness, gtx)
}

func (c Circle) BoundingBox() (bbox f32.Rectangle) {
	min := f32.Point{X: c.center.X - c.radius, Y: c.center.Y - c.radius}
	max := f32.Point{X: c.center.X + c.radius, Y: c.center.Y + c.radius}
	return f32.Rectangle{Min: min, Max: max}
}

type Ellipse struct{}

type Arc struct{}

//func scaleRectangle(r f32.Rectangle) f32.Rectangle {
//	min := r.Min.Mul(100)
//	min := r.Min.Mul(100)
//	rs := f32.Rectangle{}
//}
