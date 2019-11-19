package output

import (
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/graphics"
	"image/color"
)

func String(s *atoms.String, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &graphics.Tag{MaxWidth: graphics.Inf}
		l.Layout(gtx, st, s.String())
	}
}

func Integer(i *atoms.Integer, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &graphics.Tag{MaxWidth: graphics.Inf}
		l.Layout(gtx, st, i.String())
	}
}

func Flt(f *atoms.Flt, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &graphics.Tag{MaxWidth: graphics.Inf}
		l.Layout(gtx, st, f.StringForm(api.ToStringParams{}))
	}
}

func Complex(i *atoms.Complex, st *graphics.Style, gtx *layout.Context) layout.Widget {
	return func() {
		l := &graphics.Tag{MaxWidth: graphics.Inf}
		l.Layout(gtx, st, i.StringForm(api.ToStringParams{}))
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
