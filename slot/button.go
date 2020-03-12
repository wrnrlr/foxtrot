package slot

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	. "gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/wrnrlr/foxtrot/colors"
	"github.com/wrnrlr/foxtrot/util"
	"github.com/wrnrlr/shape"
	"image"
)

type PlusButton struct{}

func (b PlusButton) Layout(gtx *Context, button *widget.Button) {
	inset := Inset{Left: unit.Sp(20)}
	inset.Layout(gtx, func() {
		size := gtx.Px(unit.Sp(20))
		gtx.Constraints = RigidConstraints(image.Point{size, size})
		//b.drawCircle(gtx)
		b.drawPlus(gtx)
		r := image.Rectangle{Max: gtx.Dimensions.Size}
		pointer.Ellipse(r).Add(gtx.Ops)
		button.Layout(gtx)
	})
}

func (b PlusButton) drawCircle(gtx *Context) {
	width := float32(gtx.Px(unit.Sp(1)))
	px := gtx.Px(unit.Sp(20))
	size := float32(px)
	rr := float32(size) * .5
	c1 := shape.Circle{f32.Point{rr / 2, rr / 2}, rr}
	c1.Fill(util.White, gtx)
	c1.Stroke(util.LightGrey, width, gtx)
}

func (b PlusButton) drawPlus(gtx *Context) {
	s := gtx.Constraints
	w := float32(gtx.Px(unit.Sp(1)))
	yc := float32(s.Height.Min / 2)
	xc := float32(s.Width.Min / 2)
	offset := float32(s.Width.Min) / 4
	length := float32(gtx.Constraints.Width.Min) - offset
	line1 := shape.Line{{offset, yc}, {length, yc}}
	line1.Stroke(colors.Black, w, gtx)
	line2 := shape.Line{{xc, offset}, {xc, length}}
	line2.Stroke(colors.Black, w, gtx)
}
