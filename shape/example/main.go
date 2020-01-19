package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
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
				painting(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func painting(gtx *layout.Context) {
	defaultWidth := float32(gtx.Px(unit.Sp(2)))

	var stack op.StackOp
	stack.Push(gtx.Ops)
	line := shape.Line{{X: 10, Y: 10}, {X: 110, Y: 10}, {X: 210, Y: 10}}
	bbox := line.Stroke(unit.Sp(5), gtx)
	paint.ColorOp{red}.Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()

	stack.Push(gtx.Ops)
	circle := shape.Circle{f32.Point{70, 70}, 40}
	bbox = circle.Stroke(defaultWidth, gtx)
	paint.ColorOp{blue}.Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()

	stack.Push(gtx.Ops)
	circle = shape.Circle{f32.Point{160, 70}, 40}
	bbox = circle.Fill(gtx)
	paint.ColorOp{blue}.Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()

	//stack.Push(gtx)
	//shape.StrokeRectangle(f32.Point{40, 160}, f32.Point{100, 60}, 10, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()

	stack.Push(gtx.Ops)
	rect := shape.Rectangle{f32.Point{200, 160}, f32.Point{300, 400}}
	bbox = rect.Stroke(10, gtx)
	paint.ColorOp{green}.Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()

	//stack.Push(gtx)
	//shape.StrokeRectangle(f32.Point{40, 360}, f32.Point{100, 100}, 10, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()
	//
	//stack.Push(gtx)
	//shape.FillRectangle(f32.Point{40, 240}, f32.Point{30, 60}, gtx)
	//paint.ColorOp{green}.Add(gtx)
	//paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: 600, Y: 600}}}.Add(gtx)
	//stack.Pop()

	stack.Push(gtx.Ops)
	shape.FillTriangle(f32.Point{300, 10}, f32.Point{360, 10}, f32.Point{340, 60}, gtx)
	paint.ColorOp{red}.Add(gtx.Ops)
	paint.PaintOp{bbox}.Add(gtx.Ops)
	stack.Pop()
}
