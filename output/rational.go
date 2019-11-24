package output

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
)

func Rational(i *atoms.Rational, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		paint.ColorOp{black}.Add(gtx.Ops)
		num := &typeset.Label{MaxWidth: typeset.FitContent, Text: i.Num.String()}
		den := &typeset.Label{MaxWidth: typeset.FitContent, Text: i.Den.String()}
		fr := typeset.Fraction{num, den}
		fr.Layout(gtx, st.Shaper, st.Font)
	}
}
