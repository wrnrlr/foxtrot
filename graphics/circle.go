package graphics

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
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

func (c Circle) Draw(ctx *context, ops *op.Ops) {
	w, h := float32(100), float32(100)
	rr := float32(100) * .5
	rrect(ops, w, h, rr, rr, rr, rr)
	paint.ColorOp{*ctx.style.StrokeColor}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(100), Y: 100}}}.Add(ops)
}

func (c Circle) BoundingBox() (bbox f32.Rectangle) {
	min := f32.Point{X: 0, Y: 0}
	max := f32.Point{X: 1, Y: 1}
	return f32.Rectangle{Min: min, Max: max}
}

type Ellipse struct{}

type Arc struct{}

func rrect(ops *op.Ops, width, height, se, sw, nw, ne float32) {
	w, h := float32(width), float32(height)
	const c = 0.55228475 // 4*(sqrt(2)-1)/3
	var b clip.Path
	b.Begin(ops)
	b.Move(f32.Point{X: w, Y: h - se})
	b.Cube(f32.Point{X: 0, Y: se * c}, f32.Point{X: -se + se*c, Y: se}, f32.Point{X: -se, Y: se})    // SE
	b.Cube(f32.Point{X: -sw * c, Y: 0}, f32.Point{X: -sw, Y: -sw + sw*c}, f32.Point{X: -sw, Y: -sw}) // SW
	b.Cube(f32.Point{X: 0, Y: -nw * c}, f32.Point{X: nw - nw*c, Y: -nw}, f32.Point{X: nw, Y: -nw})   // NW
	b.Cube(f32.Point{X: ne * c, Y: 0}, f32.Point{X: ne, Y: ne - ne*c}, f32.Point{X: ne, Y: ne})      // NE
	// Return to origin
	b.Move(f32.Point{X: -w, Y: -ne})
	const scale = 0.95
	b.Move(f32.Point{X: w * (1 - scale) * .5, Y: h * (1 - scale) * .5})
	w *= scale
	h *= scale
	se *= scale
	sw *= scale
	nw *= scale
	ne *= scale
	b.Move(f32.Point{X: 0, Y: h - se})
	b.Cube(f32.Point{X: 0, Y: se * c}, f32.Point{X: +se - se*c, Y: se}, f32.Point{X: +se, Y: se})      // SW
	b.Cube(f32.Point{X: +sw * c, Y: 0}, f32.Point{X: +sw, Y: -sw + sw*c}, f32.Point{X: +sw, Y: -sw})   // SE
	b.Cube(f32.Point{X: 0, Y: -nw * c}, f32.Point{X: -(nw - nw*c), Y: -nw}, f32.Point{X: -nw, Y: -nw}) // NE
	b.Cube(f32.Point{X: -ne * c, Y: 0}, f32.Point{X: -ne, Y: ne - ne*c}, f32.Point{X: -ne, Y: ne})     // NW
	b.End().Add(ops)
}
