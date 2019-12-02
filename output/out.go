package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
)

func Power(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	if ex.Len() != 3 {
		return nil
	}
	left := Ex(ex.GetPart(1), st, gtx)
	right := Ex(ex.GetPart(2), st, gtx)
	if isSqrt(ex) {
		return typeset.Sqrt(left)
	}
	return typeset.Power(left, right)
}

func Parts(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) []typeset.Shape {
	var children []typeset.Shape
	var comma typeset.Shape
	for _, e := range ex.Parts[1:] {
		var shape typeset.Shape
		switch ex := e.(type) {
		case *atoms.String:
			shape = &typeset.Label{Text: ex.Val, MaxWidth: typeset.FitContent}
		case *atoms.Integer:
			shape = &typeset.Label{Text: ex.String(), MaxWidth: typeset.FitContent}
		case *atoms.Flt:
			shape = &typeset.Label{Text: ex.StringForm(api.ToStringParams{}), MaxWidth: typeset.FitContent}
		case *atoms.Rational:
			Rational(ex, st, gtx)
		case *atoms.Complex:
			shape = &typeset.Label{Text: ex.StringForm(api.ToStringParams{}), MaxWidth: typeset.FitContent}
		case *atoms.Symbol:
			shape = &typeset.Label{Text: ex.StringForm(api.ToStringParams{Context: atoms.NewString("Global`")}), MaxWidth: typeset.FitContent}
		case *atoms.Expression:
			shape = &typeset.Label{Text: ex.StringForm(api.ToStringParams{Context: atoms.NewString("Global`")}), MaxWidth: typeset.FitContent}
		}
		if comma != nil {
			children = append(children, comma)
		}
		comma = &typeset.Label{Text: ",", MaxWidth: typeset.FitContent}
		children = append(children, shape)
	}
	return children
}
