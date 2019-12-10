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
	"image"
)

type Graphics struct {
	BBox     f32.Rectangle
	style    *Style
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
	ctx := &context{}
	var stack op.StackOp
	for _, p := range g.elements {
		stack.Push(gtx.Ops)
		p.Draw(ctx, gtx.Ops)
		stack.Pop()
	}
	gtx.Dimensions = dims
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
	for _, ex := range expr.GetParts()[1:] {
		primetive, err := toPrimetive(ex)
		if err != nil {
			continue
		}
		if primetive != nil {
			g.elements = append(g.elements, primetive)
		}
	}

	return nil, &g
}

func toPrimetive(ex expreduceapi.Ex) (p Primitive, err error) {
	expr, isExpr := ex.(*atoms.Expression)
	if !isExpr {
		return nil, errors.New("primetive needs to be an expression")
	}
	switch expr.HeadStr() {
	case "System`Circle":
		p, err = toCircle(expr)
	case "System`Rectangle":
		p, err = toRectangle(expr)
	default:
		return nil, errors.New("Unknown graphics primetive")
	}
	return p, err
}
