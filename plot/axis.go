package plot

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/shape"
	"github.com/wrnrlr/foxtrot/util"
)

// Ticker creates Ticks in a specified range
//type Ticker interface {
//	// Ticks returns Ticks in a specified range
//	Ticks(min, max float32) []Tick
//}

// Normalizer rescales values from the data coordinate system to the
// normalized coordinate system.
type Normalizer interface {
	// Normalize transforms a value x in the data coordinate system to
	// the normalized coordinate system.
	Normalize(min, max, x float32) float32
}

type Axis []f32.Point

func (a Axis) Layout(bbox f32.Rectangle, gtx *layout.Context) {
	if len(a) != 2 {
		return
	}
	//p1 := a[0]
	p2 := a[1]
	var stack op.StackOp
	stack.Push(gtx.Ops)
	paint.ColorOp{Color: util.Black}.Add(gtx.Ops)
	w := float32(p2.X)
	h := float32(p2.Y)
	xAxis := []f32.Point{{0, h / 2}, {w, h / 2}}
	shape.StrokeLine(xAxis, gtx.Px(unit.Sp(1)), gtx.Ops)
	d := f32.Point{X: w, Y: h}
	r := f32.Rectangle{Max: d}
	paint.PaintOp{Rect: r}.Add(gtx.Ops)
	stack.Pop()
}
