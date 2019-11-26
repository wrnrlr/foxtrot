package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
)

func List(ex *atoms.Expression, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	open := &typeset.Label{Text: "{", MaxWidth: typeset.FitContent}
	close := &typeset.Label{Text: "}", MaxWidth: typeset.FitContent}
	var parts []typeset.Shape
	parts = append(parts, open)
	children := Parts(ex, st, gtx)
	parts = append(parts, children...)
	parts = append(parts, close)
	list := &typeset.Group{Parts: parts}
	return list
}
