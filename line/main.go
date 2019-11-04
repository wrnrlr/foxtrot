package main

import (
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/layout"
)

var lightGrey = color.RGBA{R: 0xbb, G: 0xbb, B: 0xbb, A: 0xff}
var lightPink = rgb(0xffb6c1)
var white = rgb(0xffffff)
var theme     *material.Theme
var _defaultFontSize = unit.Sp(20)
var inlineHeight = layout.Constraint{Min:50, Max:50}

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	gtx := &layout.Context{
		Queue: w.Queue(),
	}
	gofont.Register()
	theme = material.NewTheme()
	ins := &insert{1}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			ins.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

type insert struct {
	Index int
}

func (ins *insert) Layout(gtx *layout.Context) {
	inset := layout.UniformInset(unit.Dp(8))
	list := &layout.List{Axis: layout.Vertical}
	inset.Layout(gtx, func() {
		list.Layout(gtx, 3, func(i int) {
			if i == ins.Index {
				ins.inline(gtx)
			} else {
				ins.placeholder(gtx)
			}
		})
	})
}

func (ins *insert) inline(gtx *layout.Context) {
	length := gtx.Constraints.Width.Max
	gtx.Constraints.Height = inlineHeight
	st := layout.Stack{Alignment: layout.NW}
	c := st.Expand(gtx, func() {
		ins.circle(gtx)
	})
	l := st.Expand(gtx, func() {
		ins.line(length, 1, gtx)
	})
	st.Layout(gtx, l, c)
}

func (ins *insert) line(length int, width float32, gtx *layout.Context) {
	var p paint.Path
	var lineLen = float32(length)
	var merginTop = float32(gtx.Constraints.Height.Min / 2)
	var stack op.StackOp
	stack.Push(gtx.Ops)
	p.Begin(gtx.Ops)
	p.Move(f32.Point{X: 0, Y: merginTop})
	p.Line(f32.Point{X: lineLen, Y: 0})
	p.Line(f32.Point{X: 0, Y: width})
	p.Line(f32.Point{X: -lineLen, Y: 0})
	p.Line(f32.Point{X: 0, Y: -width})
	p.End().Add(gtx.Ops)
	paint.ColorOp{lightGrey}.Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{
			Max: f32.Point{X: float32(length), Y: merginTop + width},
		},
	}.Add(gtx.Ops)
	stack.Pop()
}

func (ins *insert) circle(gtx *layout.Context) {

	gtx.Constraints.Width = inlineHeight
	gtx.Constraints.Height = inlineHeight
	size := float32(50)
	rr := float32(size) * .5
	var stack op.StackOp
	stack.Push(gtx.Ops)
	rrect(gtx.Ops, size, size, rr, rr, rr, rr)
	fill(gtx, lightGrey)
	stack.Pop()
}

func (ins *insert) placeholder(gtx *layout.Context) {
	gtx.Constraints.Height = inlineHeight
	fill(gtx, white)
}

func rrect(ops *op.Ops, width, height, se, sw, nw, ne float32) {
	w, h := float32(width), float32(height)
	const c = 0.55228475 // 4*(sqrt(2)-1)/3
	var b paint.Path
	b.Begin(ops)
	b.Move(f32.Point{X: w, Y: h - se})
	b.Cube(f32.Point{X: 0, Y: se * c}, f32.Point{X: -se + se*c, Y: se}, f32.Point{X: -se, Y: se}) // SE
	b.Line(f32.Point{X: sw - w + se, Y: 0})
	b.Cube(f32.Point{X: -sw * c, Y: 0}, f32.Point{X: -sw, Y: -sw + sw*c}, f32.Point{X: -sw, Y: -sw}) // SW
	b.Line(f32.Point{X: 0, Y: nw - h + sw})
	b.Cube(f32.Point{X: 0, Y: -nw * c}, f32.Point{X: nw - nw*c, Y: -nw}, f32.Point{X: nw, Y: -nw}) // NW
	b.Line(f32.Point{X: w - ne - nw, Y: 0})
	b.Cube(f32.Point{X: ne * c, Y: 0}, f32.Point{X: ne, Y: ne - ne*c}, f32.Point{X: ne, Y: ne}) // NE
	// Return to origin
	b.Move(f32.Point{X: -w, Y: -ne})
	const scale = 0.85
	b.Move(f32.Point{X: w * (1 - scale) * .5, Y: h * (1 - scale) * .5})
	w *= scale
	h *= scale
	se *= scale
	sw *= scale
	nw *= scale
	ne *= scale
	b.Move(f32.Point{X: w, Y: h - se})
	b.Cube(f32.Point{X: 0, Y: se * c}, f32.Point{X: -se + se*c, Y: se}, f32.Point{X: -se, Y: se}) // SE
	b.Line(f32.Point{X: sw - w + se, Y: 0})
	b.Cube(f32.Point{X: -sw * c, Y: 0}, f32.Point{X: -sw, Y: -sw + sw*c}, f32.Point{X: -sw, Y: -sw}) // SW
	b.Line(f32.Point{X: 0, Y: nw - h + sw})
	b.Cube(f32.Point{X: 0, Y: -nw * c}, f32.Point{X: nw - nw*c, Y: -nw}, f32.Point{X: nw, Y: -nw}) // NW
	b.Line(f32.Point{X: w - ne - nw, Y: 0})
	b.Cube(f32.Point{X: ne * c, Y: 0}, f32.Point{X: ne, Y: ne - ne*c}, f32.Point{X: ne, Y: ne}) // NE
	b.End().Add(ops)
}

func fill(gtx *layout.Context, col color.RGBA) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: d}
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
