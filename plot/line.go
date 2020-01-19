package plot

import (
	"gioui.org/f32"
	"gioui.org/layout"
)

type XYs []f32.Point

type Line struct {
	Points XYs
}

func (l Line) Plot(gtx *layout.Context) {

}

func normalize(points []f32.Point) []f32.Point {
	return []f32.Point{}
}
