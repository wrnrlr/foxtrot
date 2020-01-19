package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/wrnrlr/foxtrot/plot"
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
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func plot1(gtx *layout.Context) {
	p := plot.Plot{}
	plot.Plot{}.Layout(gtx)
}
