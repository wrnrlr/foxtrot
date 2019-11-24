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
	"github.com/wrnrlr/foxtrot/typeset"
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
			c0 := f.Rigid(gtx, func() {
				theme.Label(unit.Sp(38), "Examples:").Layout(gtx)
			})
			c1 := f.Rigid(gtx, func() {
				l := &typeset.Label{Text: "x", MaxWidth: typeset.FitContent}
				s1 := &typeset.Label{Text: "2", MaxWidth: typeset.FitContent}
				den := &typeset.Word{Content: l, Superscript: s1}
				den.Layout(gtx, s, fnt)
			})
			c2 := f.Rigid(gtx, func() {
				l := &typeset.Label{Text: "q", MaxWidth: typeset.FitContent}
				s1 := &typeset.Label{Text: "2i", MaxWidth: typeset.FitContent}
				den := &typeset.Word{Content: l, Subscript: s1}
				den.Layout(gtx, s, fnt)
			})
			c3 := f.Rigid(gtx, func() {
				l := &typeset.Label{Text: "Abg", MaxWidth: typeset.FitContent}
				s1 := &typeset.Label{Text: "q", MaxWidth: typeset.FitContent}
				s2 := &typeset.Label{Text: "nN", MaxWidth: typeset.FitContent}
				w := typeset.Word{l, s2, s1}
				w.Layout(gtx, s, fnt)
			})
			c4 := f.Rigid(gtx, func() {
				num := &typeset.Label{Text: "12", MaxWidth: typeset.FitContent}
				den := &typeset.Label{Text: "100", MaxWidth: typeset.FitContent}
				fr := &typeset.Fraction{num, den}
				fr.Layout(gtx, s, fnt)
			})
			c5 := f.Rigid(gtx, func() {
				num := &typeset.Label{Text: "1", MaxWidth: typeset.FitContent}
				l := &typeset.Label{Text: "n", MaxWidth: typeset.FitContent}
				s1 := &typeset.Label{Text: "2", MaxWidth: typeset.FitContent}
				den := &typeset.Word{Content: l, Superscript: s1}
				fr := &typeset.Fraction{num, den}
				fr.Layout(gtx, s, fnt)
			})
			c6 := f.Rigid(gtx, func() {
				x := &typeset.Label{Text: "x", MaxWidth: typeset.FitContent}
				num := &typeset.Label{Text: "1", MaxWidth: typeset.FitContent}
				den := &typeset.Label{Text: "2", MaxWidth: typeset.FitContent}
				fr := &typeset.Fraction{num, den}
				power := &typeset.Word{Content: x, Superscript: fr}
				power.Layout(gtx, s, fnt)
			})
			f.Layout(gtx, c0, c1, c2, c3, c4, c5, c6)
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
