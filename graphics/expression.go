package graphics

import (
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
)

func Ex(ex expreduceapi.Ex, st *Style, gtx *layout.Context) layout.Widget {
	switch ex := ex.(type) {
	case *atoms.String:
		return String(ex, st, gtx)
	case *atoms.Integer:
		return Integer(ex, st, gtx)
	case *atoms.Flt:
		return Flt(ex, st, gtx)
	case *atoms.Rational:
		return Rational(ex, st, gtx)
	case *atoms.Complex:
		return Complex(ex, st, gtx)
	case *atoms.Symbol:
		return Symbol(ex, st, gtx)
	case *atoms.Expression:
		return Expression(ex, st, gtx)
	default:
		fmt.Println("unknown expression type")
	}
	return nil
}

func Expression(ex *atoms.Expression, st *Style, gtx *layout.Context) layout.Widget {
	special := drawSpecialExpression(ex, st, gtx)
	if special != nil {
		return special
	}
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: Inf}
			l1.Layout(gtx, st, shortSymbolName(ex)+"[")
		})
		children = append(children, first)
		parts := Parts(ex, f, ",", st, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: Inf}
			l1.Layout(gtx, st, "]")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}

func drawSpecialExpression(ex *atoms.Expression, st *Style, gtx *layout.Context) layout.Widget {
	switch ex.HeadStr() {
	case "System`List":
		return List(ex, st, gtx)
	case "System`Plus":
		return drawInfix(ex, "+", st, gtx)
	case "System`Minus":
		return drawInfix(ex, "-", st, gtx)
	case "System`Times":
		return drawInfix(ex, "*", st, gtx)
	case "System`Power":
		return Power(ex, st, gtx)
	case "System`Graphics":
		return Graphics(ex, st, gtx)
	}
	return nil
}
