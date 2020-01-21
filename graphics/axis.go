package graphics

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/util"
	"github.com/wrnrlr/shape"
)

type Canvas struct {
	Width, Height float32
}

type Axis []f32.Point

func (a Axis) Layout(bbox f32.Rectangle, gtx *layout.Context) {
	if len(a) != 2 {
		return
	}
	width := float32(gtx.Px(unit.Sp(1)))
	//p1 := a[0]
	p2 := a[1]
	w := float32(p2.X)
	h := float32(p2.Y)
	xAxis := []f32.Point{{0, h / 2}, {w, h / 2}}
	shape.Line(xAxis).Stroke(util.Black, width, gtx)
}
