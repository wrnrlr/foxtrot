package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/wrnrlr/foxtrot/shape"
	"image/color"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		gofont.Register()
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				painting(gtx.Ops)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func painting(ops *op.Ops) {
	var stack op.StackOp

	stack.Push(ops)
	line1 := []f32.Point{{X: 10, Y: 10}, {X: 110, Y: 10}, {X: 210, Y: 10}}
	shape.StrokeLine(line1, ops)
	paint.ColorOp{red}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
	stack.Pop()

	stack.Push(ops)
	shape.StrokeCircle(f32.Point{70, 70}, 40, 4, ops)
	paint.ColorOp{blue}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
	stack.Pop()

	stack.Push(ops)
	shape.FillCircle(f32.Point{160, 70}, 40, ops)
	paint.ColorOp{blue}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
	stack.Pop()

	//stack.Push(ops)
	//shape.FillRectangle(f32.Point{40,70}, f32.Point{30,60}, ops)
	//paint.ColorOp{green}.Add(ops)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
	//stack.Pop()

	stack.Push(ops)
	shape.StrokeRectangle(f32.Point{40, 160}, f32.Point{100, 60}, 10, ops)
	paint.ColorOp{green}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(ops)
	stack.Pop()
}
