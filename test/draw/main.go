package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/shape"
	"image"
	"image/color"
	"math"
)

const (
	rad45  = float32(45 * math.Pi / 180)
	rad135 = float32(135 * math.Pi / 180)
	rad315 = float32(315 * math.Pi / 180)
	rad225 = float32(225 * math.Pi / 180)
	rad90  = float32(90 * math.Pi / 180)
	rad180 = float32(180 * math.Pi / 180)
)

var theme *material.Theme

var (
	red   = color.RGBA{255, 0, 0, 255}
	black = color.RGBA{0, 0, 0, 255}
	lines = [][]f32.Point{
		//{{X: 10, Y: 10}, {X: 110, Y: 10}},
		//{{X: 110, Y: 10}, {X: 10, Y: 10}},
		//{{X: 10, Y: 10}, {X: 110, Y: 110}},
		//{{X: 110, Y: 110}, {X: 10, Y: 10}},

		{{X: 10, Y: 10}, {X: 210, Y: 10}},
		{{X: 10, Y: 10}, {X: 110, Y: 10}, {X: 210, Y: 10}},
		{{X: 210, Y: 10}, {X: 110, Y: 10}, {X: 10, Y: 10}},

		{{X: 10, Y: 10}, {X: 110, Y: 110}, {X: 210, Y: 10}},
		{{X: 210, Y: 10}, {X: 110, Y: 110}, {X: 10, Y: 10}},

		{{X: 10, Y: 110}, {X: 110, Y: 10}, {X: 210, Y: 110}},
		{{X: 210, Y: 110}, {X: 110, Y: 10}, {X: 10, Y: 110}},

		{{X: 10, Y: 10}, {X: 110, Y: 110}, {X: 10, Y: 210}},
		{{X: 10, Y: 210}, {X: 110, Y: 110}, {X: 10, Y: 10}},

		{{X: 110, Y: 10}, {X: 10, Y: 110}, {X: 110, Y: 210}},
		{{X: 110, Y: 210}, {X: 10, Y: 110}, {X: 110, Y: 10}},
	}
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		gofont.Register()
		list := &layout.List{Axis: layout.Vertical}
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				list.Layout(gtx, len(lines), func(i int) {
					layout.UniformInset(unit.Sp(10)).Layout(gtx, func() {
						fmt.Printf("%d ============================================================================================================\n", i)
						drawLine(lines[i], gtx)
					})
				})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawLine(points []f32.Point, gtx *layout.Context) {
	var stack op.StackOp
	stack.Push(gtx.Ops)
	shape.DrawLine(points, gtx.Ops)
	paint.ColorOp{red}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx.Ops)
	stack.Pop()
	gtx.Dimensions = layout.Dimensions{image.Point{X: 600, Y: 210}, 100}
}
