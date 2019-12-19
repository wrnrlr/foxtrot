package graphics

import (
	"errors"
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/util"
	"image"
)

type Graphics struct {
	BBox     f32.Rectangle
	ctx      *context
	elements []Primitive
	options  Options
}

func (g *Graphics) Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions {
	p := image.Point{X: 300, Y: 200}
	dims := layout.Dimensions{
		Size:     p,
		Baseline: p.Y / 2,
	}
	return dims
}

func (g *Graphics) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	dims := g.Dimensions(gtx, s, font)
	//paint.ColorOp{Color: util.Black}.Add(gtx.Ops)
	*g.ctx.style.StrokeColor = util.Black
	var stack op.StackOp
	for _, p := range g.elements {
		stack.Push(gtx.Ops)
		p.Draw(g.ctx, gtx.Ops)
		stack.Pop()
	}
	gtx.Dimensions = dims
}

func (g Graphics) drawAxis(gtx *layout.Context) {
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
		return nil, errors.New("primetive needs to be an expression")
	}
	switch expr.HeadStr() {
	case "System`RGBColor":
		p, err = toColor(expr)
	case "System`Circle":
		p, err = toCircle(expr)
	case "System`Rectangle":
		p, err = toRectangle(expr)
	default:
		return nil, errors.New("Unknown graphics primetive")
	}
	return p, err
}
