package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/wrnrlr/foxtrot/colors"
	"github.com/wrnrlr/foxtrot/style"
	"github.com/wrnrlr/foxtrot/typeset"
	"github.com/wrnrlr/foxtrot/util"
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
)

func main() {
	go func() {
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		gofont.Register()
		theme = material.NewTheme()
		theme.TextSize = unit.Sp(12)
		theme.Color.Text = util.Black
		//list := &layout.List{Axis: layout.Vertical}
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				//layout.UniformInset(unit.Sp(10)).Layout(gtx, func() {
				//	var stack op.StackOp
				//	stack.Push(gtx.Ops)
				s := style.Style{
					Font:   text.Font{Size: unit.Sp(16), Variant: "Mono"},
					Shaper: theme.Shaper,
					Color:  colors.Black,
				}
				//paint.ColorOp{Color:s.Color}.Add(gtx.Ops)
				l := &typeset.Label{Text: "x", MaxWidth: typeset.FitContent}
				s1 := &typeset.Label{Text: "2", MaxWidth: typeset.FitContent}
				den := &typeset.Word{Content: l, Superscript: s1}
				den.Layout(gtx, s)
				//stack.Pop()
				//})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
