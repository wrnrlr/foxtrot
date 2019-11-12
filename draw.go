package foxtrot

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
)

func drawString(s *atoms.String, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, s.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawInteger(i *atoms.Integer, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawFlt(i *atoms.Flt, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.StringForm(api.ToStringParams{}))
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawRational(i *atoms.Rational, gtx *layout.Context) layout.Widget {
	return func() {

		l := theme.Label(_defaultFontSize, i.StringForm(api.ToStringParams{}))
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawComplex(i *atoms.Complex, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.StringForm(api.ToStringParams{}))
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawSymbol(i *atoms.Symbol, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func drawExpression(ex *atoms.Expression, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal}
		var children []layout.FlexChild
		for _, e := range ex.Parts {
			var w layout.Widget
			switch e := e.(type) {
			case *atoms.String:
				w = drawString(e, gtx)
			case *atoms.Integer:
				w = drawInteger(e, gtx)
			case *atoms.Flt:
				w = drawFlt(e, gtx)
			case *atoms.Rational:
				w = drawRational(e, gtx)
			case *atoms.Complex:
				w = drawComplex(e, gtx)
			case *atoms.Symbol:
				w = drawSymbol(e, gtx)
			case *atoms.Expression:
				w = drawExpression(e, gtx)
			}
			c := f.Rigid(gtx, w)
			children = append(children, c)
		}
		f.Layout(gtx, children...)
	}
}
