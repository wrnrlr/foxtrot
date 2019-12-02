package browser

import (
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/wrnrlr/foxtrot/editor"
	"github.com/wrnrlr/foxtrot/util"
	"net/url"
)

type Browser struct {
	input *editor.Editor
	add   widget.Button
	url2  *url.URL
	err   error

	st editor.EditorStyle
}

func NewBrowser() *Browser {
	input := &editor.Editor{SingleLine: true}
	shaper := font.Default()
	st := editor.EditorStyle{
		Font:       text.Font{Variant: "Mono", Size: unit.Sp(18)},
		Color:      util.LightGrey,
		CaretColor: util.Black,
		Hint:       "Temporary file",
		HintColor:  util.LightGrey,
		Shaper:     shaper}
	b := Browser{input: input, st: st}
	return &b
}

func (b *Browser) url() *url.URL {
	u, _ := url.Parse(b.input.Text())
	return u
}

func (b *Browser) domain(gtx layout.Context) string {
	return b.url().Host
}

func (b *Browser) filename(gtx layout.Context) string {
	return b.url().Host
}

func (b *Browser) Layout(gtx *layout.Context) {
	padding := layout.UniformInset(unit.Sp(20))
	padding.Layout(gtx, func() {
		b.st.Layout(gtx, b.input)
	})
}

func (b *Browser) Event(gtx *layout.Context) {

}

type page struct {
}
