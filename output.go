package foxtrot

import (
	"gioui.org/layout"
	"strings"
)

type Output string

func (o *Output) String() string {
	return string(*o)
}

func (o *Output) Event() {}

func (o *Output) Layout(gtx *layout.Context) {
	if o == nil {
		return
	} else if strings.HasPrefix(o.String(), "Graphcs") {

	} else {

	}
}
