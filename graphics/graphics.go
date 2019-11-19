package graphics

import (
	"errors"
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"image"
)

func drawCanvas(ex *atoms.Expression, gtx *layout.Context) {
	w := float32(400)
	h := float32(300)
	dr := f32.Rectangle{
		Max: f32.Point{X: w, Y: h},
	}
	paint.ColorOp{Color: lightPink}.Add(gtx.Ops)
	paintRect(w, h, gtx)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{
		Size: image.Point{X: int(w), Y: int(h)}}
}

type Primitive interface {
	Draw(ctx *context, ops *op.Ops)
	BoundingBox() (bbox f32.Rectangle)
}

type context struct {
	BBox  f32.Rectangle
	style *Style
}

func (c context) x() float32 {
	return 0
}

func (c context) y() float32 {
	return 0
}

type Graphics struct {
	BBox     f32.Rectangle
	style    *Style
	ctx      *context
	elements []Primitive
	options  Options
}

func newGraphics(ex api.Ex) Graphics {
	return Graphics{}
}

func (g Graphics) dimentions() image.Rectangle {
	return image.Rectangle{Max: image.Point{X: 300, Y: 200}}
}

func (g Graphics) Layout(gtx *layout.Context) {
	for _, e := range g.elements {
		e.Draw(g.ctx, gtx.Ops)
	}
	g.drawAxis(gtx)
}

func (g Graphics) drawAxis(gtx *layout.Context) {
}

func FromEx(expr *atoms.Expression) (error, *Graphics) {
	graphics, ok := atoms.HeadAssertion(expr, "System`Graphics")
	if !ok {
		return errors.New("trying to render a non-Graphics[] expression"), nil
	}

	if graphics.Len() < 1 {
		fmt.Printf("the Graphics[] expression must have at least one argument\n")
		return errors.New("the Graphics[] expression must have at least one argument"), nil
	}

	g := Graphics{}

	return nil, &g
}
