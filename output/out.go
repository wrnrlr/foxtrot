package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
	"math/big"
)

func Power(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
	if isSqrt(ex) {
		return Sqrt(ex, st, gtx)
	}
	return drawInfix(ex, "^", st, gtx)
}

var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)

func isSqrt(ex *atoms.Expression) bool {
	if len(ex.Parts) != 3 {
		return false
	}
	r, isRational := ex.Parts[2].(*atoms.Rational)
	if !isRational {
		return false
	}
	return r.Num.Cmp(bigOne) == 0 && r.Den.Cmp(bigTwo) == 0
}

func Sqrt(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		c1 := f.Rigid(gtx, func() {
			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
			l1.Layout(gtx, st, "√")
		})
		c2 := f.Rigid(gtx, func() {
			part := ex.Parts[1]
			w := Ex(part, st, gtx)
			w()
		})
		// TODO: Draw line above body
		f.Layout(gtx, c1, c2)
	}
}

func drawInfix(ex *atoms.Expression, operator string, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		children := Parts(ex, f, operator, st, gtx)
		f.Layout(gtx, children...)
	}
}

func Parts(ex *atoms.Expression, f layout.Flex, infix string, st *graphics.Style, gtx *layout.Context) []layout.FlexChild {
	var children []layout.FlexChild
	var comma layout.FlexChild
	for _, e := range ex.Parts[1:] {
		var w layout.Widget
		switch e := e.(type) {
		case *atoms.String:
			w = String(e, st, gtx)
		case *atoms.Integer:
			w = Integer(e, st, gtx)
		case *atoms.Flt:
			w = Flt(e, st, gtx)
		case *atoms.Rational:
			w = Rational(e, st, gtx)
		case *atoms.Complex:
			w = Complex(e, st, gtx)
		case *atoms.Symbol:
			w = Symbol(e, st, gtx)
		case *atoms.Expression:
			w = Expression(e, st, gtx)
		}
		children = append(children, comma)
		comma = f.Rigid(gtx, func() {
			t := &graphics.Tag{MaxWidth: graphics.Inf}
			t.Layout(gtx, st, infix)
		})
		c := f.Rigid(gtx, w)
		children = append(children, c)
	}
	return children
}
