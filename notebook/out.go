package notebook

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
)

type Out struct {
	Ex api.Ex
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
	st := graphics.NewStyle()
	w := graphics.Ex(o.Ex, st, gtx)
	w()
}

//func (o *Out) SetState(engine *expreduce.EvalState, i int) {
//	textOut := expressionToString(engine, o.Ex, i)
//	fmt.Printf("Out: %s\n", textOut)
//	o.Text = textOut
//	//o.Image = displayExpr(o.Ex)
//}

//func displayExpr(ex api.Ex) image.Image {
//	switch e := ex.(type) {
//	case *atoms.Symbol:
//	case *atoms.Expression:
//		name := e.HeadStr()
//		if name == "System`Graphics" {
//			return RenderAsPNG(ex)
//		}
//	}
//	return nil
//}

//func RenderAsPNG(expr api.Ex) image.Image {
//	graph, err := graphics.Render(expr)
//	if err != nil {
//		return nil
//	}
//
//	buffer := bytes.NewBuffer([]byte{})
//	err = graph.Render(chart.PNG, buffer)
//	if err != nil {
//		return nil
//	}
//
//	iw := chart.ImageWriter{}
//	iw.Write(buffer.Bytes())
//
//	img, err := iw.Image()
//	if err != nil {
//		return nil
//	}
//
//	return img
//}
