package main

import (
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot"
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/layout"
)

var lightGrey = color.RGBA{R: 0xbb, G: 0xbb, B: 0xbb, A: 0xff}
var theme *material.Theme
var _defaultFontSize = unit.Sp(20)
var inlineHeight = layout.Constraint{Min: 50, Max: 50}

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
	ins := &insert{
		1,
		[]foxtrot.Placeholder{
			*foxtrot.NewPlaceholder(),
			*foxtrot.NewPlaceholder(),
		}}
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
	ps    []foxtrot.Placeholder
}

func (ins *insert) Layout(gtx *layout.Context) {
	inset := layout.UniformInset(unit.Dp(8))
	list := &layout.List{Axis: layout.Vertical}
	inset.Layout(gtx, func() {
		list.Layout(gtx, len(ins.ps), func(i int) {
			ins.ps[i].Layout(false, gtx)
		})
	})
}
