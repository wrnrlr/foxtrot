package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
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
		shape := Ex(e, st, gtx)
		if comma != nil {
			children = append(children, comma)
		}
		comma = &typeset.Label{Text: ",", MaxWidth: typeset.FitContent}
		children = append(children, shape)
	}
	return children
}
