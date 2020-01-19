package plot

import "gioui.org/layout"

type Plotter interface {
	Plot(gtx *layout.Context)
}

type Plot struct {
	XAxis, YXXis Axis
	Plots        []Plotter
}

func (p *Plot) Add(plotter Plotter) {
	p.Plots = append(p.Plots, plotter)
}

func (p Plot) Layout(gtx *layout.Context) {
	for _, p := range p.Plots {

	}
}
