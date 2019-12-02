package output

import (
	"github.com/corywalker/expreduce/expreduce/atoms"
	"math/big"
)

var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)

func isSqrt(ex *atoms.Expression) bool {
	if len(ex.Parts) < 2 {
		return false
	}
	r, isRational := ex.Parts[2].(*atoms.Rational)
	if !isRational {
		return false
	}
	return r.Num.Cmp(bigOne) == 0 && r.Den.Cmp(bigTwo) == 0
}

//func Sqrt(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
//	return func() {
//		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
//		c1 := f.Rigid(gtx, func() {
//			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
//			l1.Layout(gtx, st, "√")
//		})
//		c2 := f.Rigid(gtx, func() {
//			part := ex.Parts[1]
//			w := Ex(part, st, gtx)
//			w()
//		})
//		// TODO: Draw line above body
//		f.Layout(gtx, c1, c2)
//	}
//}
