package graphics

import (
	"errors"
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/util"
	"github.com/wrnrlr/shape"
	"image"
)

type Box f32.Rectangle

func (b Box) Width() (w float32) {
	if b.Min.X < 0 {
		w += util.Absf32(b.Min.X)
	} else {
		w -= b.Min.X
	}
	return w + b.Max.X
}

func (b Box) Height() (h float32) {
	if b.Min.Y < 0 {
		h += util.Absf32(b.Min.Y)
	} else {
		h -= b.Min.Y
	}
	return h + b.Max.Y
}

type Graphics struct {
	BBox     f32.Rectangle
	ctx      *context
	elements primetives
	options  Options
}

type primetives []Primitive

func (ps primetives) bbox() (bbox f32.Rectangle) {
	for _, e := range ps {
		b := e.BoundingBox()
		bbox.Min.X = min(b.Min.X, bbox.Min.X)
		bbox.Min.Y = min(b.Min.Y, bbox.Min.Y)
		bbox.Max.X = max(b.Max.X, bbox.Max.X)
		bbox.Max.Y = max(b.Max.Y, bbox.Max.Y)
	}
	return bbox
}

func (g *Graphics) Dimensions(gtx *layout.Context, _ *text.Shaper, font text.Font) layout.Dimensions {
	//width := g.BBox.Min.X*-1 + g.BBox.Max.X
	//height := g.BBox.Min.Y*-1 + g.BBox.Max.Y
	g.BBox = g.elements.bbox()
	r := g.size()
	f := float32(gtx.Px(unit.Sp(100)))
	r = r.Mul(f)
	p := image.Point{X: int(r.X), Y: int(r.Y)}
	dims := layout.Dimensions{
		Size:     p,
		Baseline: p.Y / 2,
	}
	return dims
}

func (g *Graphics) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	// X Axis
	dims := g.Dimensions(gtx, s, font)
	g.drawAxis(gtx, s, font)
	g.drawYAxis(gtx, s, font)
	*g.ctx.style.StrokeColor = util.Black
	g.ctx.style.Thickness = float32(0.95)
	var stack op.StackOp
	for _, p := range g.elements {
		stack.Push(gtx.Ops)
		p.Draw(g.ctx, gtx)
		stack.Pop()
	}
	gtx.Dimensions = dims
}

func (g Graphics) drawAxis(gtx *layout.Context, s *text.Shaper, font text.Font) {
	width := float32(gtx.Px(unit.Sp(1)))
	dims := g.Dimensions(gtx, s, font)
	w := float32(gtx.Constraints.Width.Max)
	h := float32(dims.Size.Y)
	xAxis := shape.Line{{0, h / 2}, {w, h / 2}}
	xAxis.Stroke(util.Black, width, gtx)
}

func (g Graphics) drawYAxis(gtx *layout.Context, s *text.Shaper, font text.Font) {
	width := float32(gtx.Px(unit.Sp(1)))
	dims := g.Dimensions(gtx, s, font)
	w := float32(dims.Size.X)
	h := float32(dims.Size.Y)
	yAxis := []f32.Point{{w / 2, 0}, {w / 2, h}}
	shape.Line(yAxis).Stroke(util.Black, width, gtx)
}

func (g *Graphics) size() (bbox f32.Point) {
	width := g.BBox.Min.X*-1 + g.BBox.Max.X
	height := g.BBox.Min.Y*-1 + g.BBox.Max.Y
	return f32.Point{width, height}
}

func (g *Graphics) calculateBoundingBox() (bbox f32.Rectangle) {
	for _, e := range g.elements {
		b := e.BoundingBox()
		bbox.Min.X = min(b.Min.X, bbox.Min.X)
		bbox.Min.Y = min(b.Min.Y, bbox.Min.Y)
		bbox.Max.X = max(b.Max.X, bbox.Max.X)
		bbox.Max.Y = max(b.Max.Y, bbox.Max.Y)
	}
	return bbox
}

func FromEx(expr *atoms.Expression, st *Style) (*Graphics, error) {
	graphics, ok := atoms.HeadAssertion(expr, "System`Graphics")
	if !ok {
		return nil, errors.New("trying to render a non-Graphics[] expression")
	}

	if graphics.Len() < 1 {
		fmt.Printf("the Graphics[] expression must have at least one argument\n")
		return nil, errors.New("the Graphics[] expression must have at least one argument")
	}

	ctx := &context{style: st}
	g := Graphics{ctx: ctx}
	primitives, err := toPrimetives(expr.GetParts()[1])
	if err != nil {
		return nil, err
	}
	g.elements = primitives
	g.calculateBoundingBox()
	return &g, err
}

func toPrimetives(ex expreduceapi.Ex) (ps []Primitive, err error) {
	expr, isExpr := ex.(*atoms.Expression)
	if !isExpr {
		return nil, errors.New("Graphics[] first argument should be a primitive or list of primitives")
	}
	isList := expr.HeadStr() == "System`List"
	if isList {
		for _, ex := range expr.GetParts()[1:] {
			p, err := toPrimetive(ex)
			if err != nil {
				continue
			}
			if p != nil {
				ps = append(ps, p)
			}
		}
	} else {
		p, err := toPrimetive(ex)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func toPrimetive(ex expreduceapi.Ex) (p Primitive, err error) {
	expr, isExpr := ex.(*atoms.Expression)
	if !isExpr {
		return nil, errors.New("primitive needs to be an expression")
	}
	switch expr.HeadStr() {
	case "System`RGBColor":
		p, err = toColor(expr)
	case "System`Circle":
		p, err = toCircle(expr)
	case "System`Rectangle":
		p, err = toRectangle(expr)
	case "System`Line":
		p, err = toLine(expr)
	case "System`Triangle":
		p, err = toTriangle(expr)
	default:
		return nil, errors.New("unknown graphics primitive")
	}
	return p, err
}
