package app

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/wrnrlr/foxtrot/browser"
	"github.com/wrnrlr/foxtrot/notebook"
	"log"
)

func RunUI() {
	gofont.Register()
	go func() {
		w := app.NewWindow(app.Title("Foxtrot"))
		a := NewApp()
		if err := a.loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

type App struct {
	nb *notebook.Notebook
	br *browser.Browser
}

func NewApp() *App {
	nb := notebook.NewNotebook()
	br := browser.NewBrowser()
	return &App{nb, br}
}

func (a *App) loop(w *app.Window) error {
	gtx := &layout.Context{
		Queue: w.Queue(),
	}
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)
			a.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func (a *App) Event(gtx *layout.Context) interface{} {
	a.nb.Event(gtx)
	return nil
}

func (a *App) Layout(gtx *layout.Context) {
	a.Event(gtx)
	f := layout.Flex{Axis: layout.Vertical}
	c1 := f.Rigid(gtx, func() {
		a.br.Layout(gtx)
	})
	c2 := f.Flex(gtx, 1, func() {
		a.nb.Layout(gtx)
	})
	f.Layout(gtx, c1, c2)
}
