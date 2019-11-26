package output

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
)

func Rational(i *atoms.Rational, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	paint.ColorOp{black}.Add(gtx.Ops)
	num := &typeset.Label{MaxWidth: typeset.FitContent, Text: i.Num.String()}
	den := &typeset.Label{MaxWidth: typeset.FitContent, Text: i.Den.String()}
	fraction := typeset.Fraction{num, den}
	return &fraction
}
