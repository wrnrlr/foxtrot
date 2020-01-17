package plot

import "gioui.org/layout"

type Plotter interface {
	Plot(gtx *layout.Context)
}

type Plot struct {
	Title struct {
		Text string
	}
	Plots []Plotter
}
