package main

import (
	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/tex"
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
	theme.TextSize = unit.Sp(12)
	theme.Color.Text = black
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			s := font.Default()
			fnt := text.Font{Size: unit.Sp(20)}
			f := layout.Flex{Axis: layout.Vertical}
			paint.ColorOp{black}.Add(gtx.Ops)
			c1 := f.Rigid(gtx, func() {
				theme.Label(unit.Sp(38), "Examples:").Layout(gtx)
			})
			c2 := f.Rigid(gtx, func() {
				l := &tex.Label{Text: "Abc", MaxWidth: tex.FitContent}
				s1 := &tex.Label{Text: "q", MaxWidth: tex.FitContent}
				s2 := &tex.Label{Text: "nN", MaxWidth: tex.FitContent}
				w := tex.Word{l, s2, s1}
				w.Layout(gtx, s, fnt)
			})
			c3 := f.Rigid(gtx, func() {
				num := &tex.Label{Text: "12", MaxWidth: tex.FitContent}
				den := &tex.Label{Text: "100", MaxWidth: tex.FitContent}
				fr := &tex.Fraction{num, den}
				fr.Layout(gtx, s, fnt)
			})
			f.Layout(gtx, c1, c2, c3)
			e.Frame(gtx.Ops)
		}
	}
}

var (
	black = rgb(0x000000)
)

func rgb(c uint32) color.RGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
