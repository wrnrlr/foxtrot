package graphics

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
)

func List(ex *atoms.Expression, st *Style, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: Inf}
			l1.Layout(gtx, st, "{")
		})
		children = append(children, first)
		parts := Parts(ex, f, ",", st, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: Inf}
			l1.Layout(gtx, st, "}")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}
