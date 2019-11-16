package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/graphics"
	"image"
	"image/color"
	"log"
)

var theme *material.Theme

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
	theme.TextSize = unit.Sp(50)
	theme.Color.Text = black
	//list := &layout.List{Axis: layout.Vertical}
	//items := []Item{newItem(), newItem()} //,newItem(),newItem(),newItem(),newItem(),newItem(),newItem(),newItem(),newItem(),newItem(),newItem(),newItem()}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			layout.UniformInset(unit.Sp(10)).Layout(gtx, func() {
				st := graphics.NewStyle()
				t := graphics.Tag{MaxWidth: graphics.Inf}
				t.Layout(gtx, st, "hello")
				//l := widget.Label{}
				//f := text.Font{Size: _defaultFontSize}
				//l.Layout(gtx, st.Shaper, f, "hello")
			})
			e.Frame(gtx.Ops)
		}
	}
}

type Item struct {
	editor *widget.Editor
}

func newItem() Item {
	return Item{&widget.Editor{}}
}

func (ins *Item) Layout(gtx *layout.Context) {
	theme.Editor("Hellooooooooooo").Layout(gtx, ins.editor)
}

var (
	lightPink        = rgb(0xffb6c1)
	lightBlue        = rgb(0x89cff0)
	black            = rgb(0x000000)
	_defaultFontSize = unit.Sp(20)
)

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func toPointF(p image.Point) f32.Point {
	return f32.Point{X: float32(p.X), Y: float32(p.Y)}
}
