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

func drawSpecialExpression(ex *atoms.Expression, st *Style, gtx *layout.Context) layout.Widget {
	switch ex.HeadStr() {
	case "System`List":
		return List(ex, st, gtx)
	case "System`Plus":
		return drawInfix(ex, "+", st, gtx)
	case "System`Minus":
		return drawInfix(ex, "-", st, gtx)
	case "System`Times":
		return drawInfix(ex, "*", st, gtx)
	case "System`Power":
		return Power(ex, st, gtx)
	}
	return nil
}

func shortSymbolName(ex *atoms.Expression) string {
	name := ex.HeadStr()
	if strings.HasPrefix(name, "System`") {
		return name[7:]
	} else {
		return name
	}
}
