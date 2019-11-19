package output

import (
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
)

func Ex(ex api.Ex, st *graphics.Style, gtx *layout.Context) layout.Widget {
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

func Expression(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
	special := drawSpecialExpression(ex, st, gtx)
	if special != nil {
		return special
	}
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
			l1.Layout(gtx, st, shortSymbolName(ex)+"[")
		})
		children = append(children, first)
		parts := Parts(ex, f, ",", st, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
			l1.Layout(gtx, st, "]")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}

func drawSpecialExpression(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
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
		err, g := graphics.FromEx(ex)
		if err != nil {
			fmt.Printf("Error rendering Graphics output: %v", err)
			return nil
		}
		return func() {
			g.Layout(gtx)
		}
	}
	return nil
}
