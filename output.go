package foxtrot

import (
	"bytes"
	"fmt"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/graphics"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wcharczuk/go-chart"
	"image"
)

func NewOut(es expreduceapi.Ex) *Out {
	return &Out{es, "", nil}
}

type Out struct {
	Ex    expreduceapi.Ex
	Text  string
	Image image.Image
}

func (o *Out) Layout(num int, gtx *layout.Context) {
	flex := &layout.Flex{Alignment: layout.Middle}
	c1 := flex.Rigid(gtx, func() {
		o.promptLayout(num, gtx)
	})
	c2 := flex.Flex(gtx, 1, func() {
		if o.Image == nil {
			o.outEditor().Layout(gtx)
		} else {
			avatarOp := paint.NewImageOp(o.Image)
			imga := theme.Image(avatarOp)
			imga.Layout(gtx)
		}
	})
	layout.Inset{Bottom: _padding}.Layout(gtx, func() {
		flex.Layout(gtx, c1, c2)
	})
}

func (o *Out) promptLayout(num int, gtx *layout.Context) {
	var txt string
	if num < 0 {
		txt = fmt.Sprintf("Out[ ] ")
	} else {
		txt = fmt.Sprintf("Out[%d] ", num)
	}
	px := gtx.Config.Px(promptWidth)
	constraint := layout.Constraint{Min: px, Max: px}
	gtx.Constraints.Width = constraint
	label := promptTheme.Label(_promptFontSize, txt)
	label.Alignment = text.End
	label.Layout(gtx)
}

func (o *Out) outEditor() material.Label {
	l := theme.Label(_defaultFontSize, o.Text)
	l.Font.Variant = "Mono"
	return l
}

func (o *Out) SetState(engine *expreduce.EvalState, i int) {
	textOut := expressionToString(engine, o.Ex, i)
	o.Text = textOut
	o.Image = displayExpr(o.Ex)
}

func displayExpr(ex expreduceapi.Ex) image.Image {
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

func RenderAsPNG(expr expreduceapi.Ex) image.Image {
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
