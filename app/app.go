package app

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/nbx"
	"github.com/wrnrlr/foxtrot/notebook"
	"log"
)

func RunUI(p string) {
	gofont.Register()
	go func() {
		w := app.NewWindow(app.Title("Foxtrot"))
		a := NewApp(p)
		if err := a.loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

type App struct {
	path string
	nb   *notebook.Notebook
	//br *browser.Browser
}

func NewApp(p string) *App {
	cells, err := nbx.ReadNBX(p)
	if err != nil {
		fmt.Printf("failed to open file")
	}
	nb := notebook.NewNotebook(cells)
	//br := browser.NewBrowser()
	return &App{p, nb}
}

func (a *App) loop(w *app.Window) error {
	gtx := layout.NewContext(w.Queue())
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			a.save()
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
	margin := unit.Sp(2)
	i := layout.Inset{Top: margin}
	i.Layout(gtx, func() {
		f := layout.Flex{Axis: layout.Vertical}
		//c1 := f.Rigid(gtx, func() {
		//	a.br.Layout(gtx)
		//})
		c2 := layout.Flexed(1, func() {
			a.nb.Layout(gtx)
		})
		f.Layout(gtx, c2)
	})
}

func (a *App) save() {
	fmt.Println("Quiting Foxtrot, Save notebook")
	if a.path == "" {
		return
	}
	err := nbx.WriteFile(a.path, a.nb.Cells)
	if err != nil {
		fmt.Println(err)
	}
}
