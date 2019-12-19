package output

import (
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
	"github.com/wrnrlr/foxtrot/util"
)

func FromEx(ex api.Ex, gtx *layout.Context) typeset.Shape {
	st := &graphics.Style{StrokeColor: &util.LightPink}
	return Ex(ex, st, gtx)
}

func Ex(ex api.Ex, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	switch ex := ex.(type) {
	case *atoms.String:
		return &typeset.Label{Text: ex.Val, MaxWidth: typeset.FitContent}
	case *atoms.Integer:
		return &typeset.Label{Text: ex.String(), MaxWidth: typeset.FitContent}
	case *atoms.Flt:
		return &typeset.Label{Text: ex.StringForm(api.ToStringParams{}), MaxWidth: typeset.FitContent}
	case *atoms.Rational:
		return Rational(ex, st, gtx)
	case *atoms.Complex:
		return &typeset.Label{Text: ex.StringForm(api.ToStringParams{Context: atoms.NewString("Global`")}), MaxWidth: typeset.FitContent}
	case *atoms.Symbol:
		return &typeset.Label{Text: ex.StringForm(api.ToStringParams{Context: atoms.NewString("Global`")}), MaxWidth: typeset.FitContent}
	case *atoms.Expression:
		return Expression(ex, st, gtx)
	default:
		fmt.Println("unknown expression type")
		return &typeset.Label{Text: "unknown expression type", MaxWidth: typeset.FitContent}
	}
}

func Expression(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	special := drawSpecialExpression(ex, st, gtx)
	if special != nil {
		return special
	}
	head := &typeset.Label{Text: shortExpressionName(ex), MaxWidth: typeset.FitContent}
	open := &typeset.Label{Text: "[", MaxWidth: typeset.FitContent}
	var parts []typeset.Shape
	parts = append(parts, head)
	parts = append(parts, open)
	children := Parts(ex, st, gtx)
	parts = append(parts, children...)
	close := &typeset.Label{Text: "]", MaxWidth: typeset.FitContent}
	parts = append(parts, close)
	expr := &typeset.Group{Parts: parts}
	return expr
}

func drawSpecialExpression(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	shape := binaryOperation(ex, st, gtx)
	if shape != nil {
		return shape
	}
	switch ex.HeadStr() {
	case "System`List":
		return List(ex, st, gtx)
	case "System`Graphics":
		g, err := graphics.FromEx(ex, st)
		if err != nil {
			fmt.Printf("Error rendering Graphics output: %v", err)
			return nil
		}
		return g
	default:
		return nil
	}
	return nil
}

func binaryOperation(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	if ex.Len() != 2 {
		return nil
	}
	left := Ex(ex.GetPart(1), st, gtx)
	right := Ex(ex.GetPart(2), st, gtx)
	switch ex.HeadStr() {
	case "System`Plus":
		return typeset.Plus(left, right)
	case "System`Minus":
		return typeset.Minus(left, right)
	case "System`Times":
		return typeset.Multiply(left, right)
	case "System`Power":
		if isSqrt(ex) {
			return typeset.Sqrt(left)
		} else {
			return typeset.Power(left, right)
		}
	default:
		return nil
	}
}
