package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/typeset"
	"strings"
)

func Symbol(sym *atoms.Symbol, st *graphics.Style, gtx *layout.Context) typeset.Shape {
	txt := shortSymbolName(sym)
	return &typeset.Label{Text: txt, MaxWidth: typeset.FitContent}
}

func shortSymbolName(sym *atoms.Symbol) string {
	name := sym.Name
	if strings.HasPrefix(name, "System`") {
		return name[7:]
	} else {
		return name
	}
}

func shortExpressionName(ex *atoms.Expression) string {
	name := ex.HeadStr()
	if strings.HasPrefix(name, "System`") {
		return name[7:]
	} else {
		return name
	}
}
