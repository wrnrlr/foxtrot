package graphics

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"strings"
)

func Symbol(i *atoms.Symbol, st *Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &Tag{MaxWidth: Inf}
		l.Layout(gtx, st, i.StringForm(api.ToStringParams{}))
	}
}

func shortSymbolName(ex *atoms.Expression) string {
	name := ex.HeadStr()
	if strings.HasPrefix(name, "System`") {
		return name[7:]
	} else {
		return name
	}
}
