package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
)

func List(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
			l1.Layout(gtx, st, "{")
		})
		children = append(children, first)
		parts := Parts(ex, f, ",", st, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &graphics.Tag{MaxWidth: graphics.Inf}
			l1.Layout(gtx, st, "}")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}
