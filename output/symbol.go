package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
	"strings"
)

func Symbol(i *atoms.Symbol, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &graphics.Tag{MaxWidth: graphics.Inf}
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
