package foxtrot

import (
	"bytes"
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/graphics"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wcharczuk/go-chart"
	"image"
	"strings"
)

func NewOut(es api.Ex) *Out {
	return &Out{es, "", nil}
}

type Out struct {
	Ex    api.Ex
	Text  string
	Image image.Image
}

func (o *Out) Layout(num int, gtx *layout.Context) {
	layout.Inset{Top: unit.Sp(15)}.Layout(gtx, func() {
		flex := &layout.Flex{Alignment: layout.Middle}
		c1 := flex.Rigid(gtx, func() {
			o.promptLayout(num, gtx)
		})
		c2 := flex.Flex(gtx, 1, func() {
			f := &layout.Flex{Axis: layout.Vertical}
			c2 := f.Rigid(gtx, func() {
				o.expressionLayout(gtx)
			})
			f.Layout(gtx, c2)
		})
		layout.Inset{Bottom: _padding}.Layout(gtx, func() {
			flex.Layout(gtx, c1, c2)
		})
	})
}

func (o *Out) promptLayout(num int, gtx *layout.Context) {
	var txt string
	if num < 0 {
		txt = fmt.Sprintf("Out[ ] ")
	} else {
		txt = fmt.Sprintf("Out[%d] ", num)
	}
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		px := gtx.Config.Px(promptWidth)
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		label := promptTheme.Label(_promptFontSize, txt)
		label.Alignment = text.End
		label.Layout(gtx)
	})
}

func (o *Out) expressionLayout(gtx *layout.Context) {
	var w layout.Widget
	switch ex := o.Ex.(type) {
	case *atoms.String:
		w = o.drawString(ex, gtx)
	case *atoms.Integer:
		w = o.drawInteger(ex, gtx)
	case *atoms.Flt:
		w = o.drawFlt(ex, gtx)
	case *atoms.Rational:
		w = o.drawRational(ex, gtx)
	case *atoms.Complex:
		w = o.drawComplex(ex, gtx)
	case *atoms.Symbol:
		w = o.drawSymbol(ex, gtx)
	case *atoms.Expression:
		w = o.drawExpression(ex, gtx)
	}
	w()
}

func (o *Out) drawString(s *atoms.String, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, s.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func (o *Out) drawInteger(i *atoms.Integer, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func (o *Out) drawFlt(i *atoms.Flt, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.StringForm(api.ToStringParams{}))
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func (o *Out) drawRational(i *atoms.Rational, gtx *layout.Context) layout.Widget {
	return func() {
		Rational2(i.Num, i.Den, gtx)
	}
}

func (o *Out) drawComplex(i *atoms.Complex, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.StringForm(api.ToStringParams{}))
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func (o *Out) drawSymbol(i *atoms.Symbol, gtx *layout.Context) layout.Widget {
	return func() {
		l := theme.Label(_defaultFontSize, i.String())
		l.Font.Variant = "Mono"
		l.Layout(gtx)
	}
}

func (o *Out) drawExpression(ex *atoms.Expression, gtx *layout.Context) layout.Widget {
	special := o.drawSpecialExpression(ex, gtx)
	if special != nil {
		return special
	}
	return func() {
		o.drawSpecialExpression(ex, gtx)

		f := layout.Flex{Axis: layout.Horizontal}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: inf}
			l1.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, shortSymbolName(ex)+"[")
		})
		children = append(children, first)
		parts := o.drawParts(ex, f, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: inf}
			l1.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, "]")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}

func (o *Out) drawSpecialExpression(ex *atoms.Expression, gtx *layout.Context) layout.Widget {
	switch ex.HeadStr() {
	case "System`List":
		return o.drawList(ex, gtx)
	case "System`Plus":
	case "System`Minus":
	case "System`Times":
	case "System`Power":
	}
	return nil
}

func (o *Out) drawList(ex *atoms.Expression, gtx *layout.Context) layout.Widget {
	return func() {
		f := layout.Flex{Axis: layout.Horizontal}
		var children []layout.FlexChild
		first := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: inf}
			l1.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, "{")
		})
		children = append(children, first)
		parts := o.drawParts(ex, f, gtx)
		children = append(children, parts...)
		last := f.Rigid(gtx, func() {
			l1 := &Tag{MaxWidth: inf}
			l1.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, "}")
		})
		children = append(children, last)
		f.Layout(gtx, children...)
	}
}

func (o *Out) drawParts(ex *atoms.Expression, f layout.Flex, gtx *layout.Context) []layout.FlexChild {
	var children []layout.FlexChild
	var comma layout.FlexChild
	for _, e := range ex.Parts[1:] {
		var w layout.Widget
		switch e := e.(type) {
		case *atoms.String:
			w = o.drawString(e, gtx)
		case *atoms.Integer:
			w = o.drawInteger(e, gtx)
		case *atoms.Flt:
			w = o.drawFlt(e, gtx)
		case *atoms.Rational:
			w = o.drawRational(e, gtx)
		case *atoms.Complex:
			w = o.drawComplex(e, gtx)
		case *atoms.Symbol:
			w = o.drawSymbol(e, gtx)
		case *atoms.Expression:
			w = o.drawExpression(e, gtx)
		}
		children = append(children, comma)
		comma = f.Rigid(gtx, func() {
			t := &Tag{MaxWidth: inf}
			t.Layout(gtx, theme.Shaper, text.Font{Size: theme.TextSize}, ",")
		})
		c := f.Rigid(gtx, w)
		children = append(children, c)
	}
	return children
}

func shortSymbolName(ex *atoms.Expression) string {
	name := ex.HeadStr()
	if strings.HasPrefix(name, "System`") {
		return name[7:]
	} else {
		return name
	}
}

func (o *Out) txtLayout(txt string, gtx *layout.Context) {
	l := theme.Label(_defaultFontSize, txt)
	l.Font.Variant = "Mono"
	l.Layout(gtx)
}

func (o *Out) outEditor() material.Label {
	l := theme.Label(_defaultFontSize, o.Text)
	l.Font.Variant = "Mono"
	return l
}

func (o *Out) SetState(engine *expreduce.EvalState, i int) {
	textOut := expressionToString(engine, o.Ex, i)
	fmt.Printf("Out: %s\n", textOut)
	o.Text = textOut
	o.Image = displayExpr(o.Ex)
}

func displayExpr(ex api.Ex) image.Image {
	switch e := ex.(type) {
	case *atoms.Symbol:
	case *atoms.Expression:
		name := e.HeadStr()
		if name == "System`Graphics" {
			return RenderAsPNG(ex)
		}
	}
	return nil
}

func RenderAsPNG(expr api.Ex) image.Image {
	graph, err := graphics.Render(expr)
	if err != nil {
		return nil
	}

	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		return nil
	}

	iw := chart.ImageWriter{}
	iw.Write(buffer.Bytes())

	img, err := iw.Image()
	if err != nil {
		return nil
	}

	return img
}
