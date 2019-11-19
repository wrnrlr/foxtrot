package notebook

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
	"github.com/wrnrlr/foxtrot/output"
)

type Out struct {
	Ex  api.Ex
	Err error
}

func (o *Out) Layout(num int, gtx *layout.Context) {
	layout.Inset{Top: unit.Sp(15)}.Layout(gtx, func() {
		flex := &layout.Flex{Alignment: layout.Middle}
		c1 := flex.Rigid(gtx, func() {
			o.promptLayout(num, gtx)
		})
		c2 := flex.Flex(gtx, 1, func() {
			if o.Err == nil {
				o.expressionLayout(gtx)
			} else {
				o.errorLayout(gtx)
			}
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

func (o *Out) errorLayout(gtx *layout.Context) {
	paint.ColorOp{Color: red}.Add(gtx.Ops)
	l := &widget.Label{}
	ft := text.Font{Size: unit.Sp(18)}
	shaper := font.Default()
	l.Layout(gtx, shaper, ft, o.Err.Error())
}

func (o *Out) expressionLayout(gtx *layout.Context) {
	st := graphics.NewStyle()
	w := output.Ex(o.Ex, st, gtx)
	w()
}
