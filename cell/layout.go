package cell

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/colors"
	"github.com/wrnrlr/foxtrot/output"
	"github.com/wrnrlr/foxtrot/style"
)

func (c cell) Layout(selected bool, gtx *layout.Context) {
	// Layout Slot
	// Layout Label
	// layout Text
	// Layout Output
	// Layout margin
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		c.margin.Layout(gtx, selected, func() {
			c.cellLayout(gtx)
		})
	})
}

func (c *cell) cellLayout(gtx *layout.Context) {
	switch c.Type() {
	case Input:
		c.layoutInput(gtx)
	case Output:
		c.output(gtx)
	case H1:
		c.h1(gtx)
	case H2:
		c.h2(gtx)
	case H3:
		c.h3(gtx)
	case H4:
		c.h4(gtx)
	case Paragraph:
		c.text(gtx)
	case Code:
		c.code(gtx)
	}
}

func (c cell) layoutSlot(gtx *layout.Context) {

}

func (c cell) layoutContent(gtx *layout.Context) {
	c.layoutLabel(gtx)
	c.layoutMargin(gtx)
}

func (c cell) layoutLabel(gtx *layout.Context)  {}
func (c cell) layoutMargin(gtx *layout.Context) {}

func (c *cell) layoutInput(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := layout.Rigid(func() {
		c.labelLayout(gtx)
	})
	c2 := layout.Flexed(1, func() {
		c.styles.Foxtrot.Layout(gtx, c.input)
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}

func (c *cell) output(gtx *layout.Context) {
	f := layout.Flex{Alignment: layout.Middle}
	c1 := layout.Rigid(func() {
		c.labelLayout(gtx)
	})
	c2 := layout.Flexed(1, func() {
		if c.out == nil {
			return
		}
		w := output.FromEx(c.out, gtx)
		var stack op.StackOp
		stack.Push(gtx.Ops)
		//paint.ColorOp{Color: util.Black}.Add(gtx.Ops)
		// c.styles.Theme.Shaper, text.Font{Size: unit.Sp(16)}
		s := style.Style{
			Font:   text.Font{Size: unit.Sp(16)},
			Shaper: c.styles.Theme.Shaper,
			Color:  colors.Black,
		}
		w.Layout(gtx, s)
		stack.Pop()
	})
	layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func() {
		f.Layout(gtx, c1, c2)
	})
}

func (c *cell) h1(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H1.Layout(gtx, c.input)
	})
}

func (c *cell) h2(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H2.Layout(gtx, c.input)
	})
}

func (c *cell) h3(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H3.Layout(gtx, c.input)
	})
}

func (c *cell) h4(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.H4.Layout(gtx, c.input)
	})
}

func (c *cell) text(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Text.Layout(gtx, c.input)
	})
}

func (c *cell) code(gtx *layout.Context) {
	layout.Inset{Left: unit.Sp(10)}.Layout(gtx, func() {
		c.styles.Code.Layout(gtx, c.input)
	})
}

func (c *cell) labelLayout(gtx *layout.Context) {
	layout.Inset{Right: unit.Sp(10)}.Layout(gtx, func() {
		px := gtx.Px(unit.Sp(50))
		constraint := layout.Constraint{Min: px, Max: px}
		gtx.Constraints.Width = constraint
		label := c.styles.Theme.Label(unit.Sp(12), c.label)
		label.Alignment = text.End
		label.Layout(gtx)
	})
}
