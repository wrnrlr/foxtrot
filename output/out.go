package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
)

func Power(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
	if isSqrt(ex) {
		return Sqrt(ex, st, gtx)
	}
	return drawInfix(ex, "^", st, gtx)
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
