package editor

import (
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"golang.org/x/image/math/fixed"
	"image"
	"time"
)

// Layout lays out the editor.
func (e *Editor) Layout(gtx *layout.Context, sh *text.Shaper, font text.Font) {
	// Flush events from before the previous frame.
	copy(e.events, e.events[e.prevEvents:])
	e.events = e.events[:len(e.events)-e.prevEvents]
	e.prevEvents = len(e.events)
	if e.font != font {
		e.invalidate()
		e.font = font
	}
	e.processEvents(gtx)
	e.layout(gtx, sh)
}

func (e *Editor) layout(gtx *layout.Context, sh *text.Shaper) {
	// Crude configuration change detection.
	if scale := gtx.Px(unit.Sp(100)); scale != e.scale {
		e.invalidate()
		e.scale = scale
	}
	cs := gtx.Constraints
	e.carWidth = fixed.I(gtx.Px(unit.Dp(1)))

	maxWidth := cs.Width.Max
	if e.SingleLine {
		maxWidth = inf
	}
	if maxWidth != e.maxWidth {
		e.maxWidth = maxWidth
		e.invalidate()
	}

	if !e.valid {
		e.lines, e.dims = e.layoutText(gtx, sh, e.font)
		e.valid = true
		e.lineCount = len(e.lines)
	}

	e.viewSize = cs.Constrain(e.dims.Size)
	// Adjust scrolling for new viewport and layout.
	e.scrollRel(0, 0)

	if e.caretScroll {
		e.caretScroll = false
		e.scrollToCaret()
	}

	off := image.Point{
		X: -e.scrollOff.X,
		Y: -e.scrollOff.Y,
	}
	clip := textPadding(e.lines)
	clip.Max = clip.Max.Add(e.viewSize)
	it := lineIterator{
		Lines:     e.lines,
		Clip:      clip,
		Alignment: e.Alignment,
		Width:     e.viewSize.X,
		Offset:    off,
	}
	e.shapes = e.shapes[:0]
	for {
		str, off, ok := it.Next()
		if !ok {
			break
		}
		path := sh.Shape(gtx, e.font, str)
		e.shapes = append(e.shapes, line{off, path})
	}

	key.InputOp{Key: &e.eventKey, Focus: e.requestFocus}.Add(gtx.Ops)
	e.requestFocus = false
	pointerPadding := gtx.Px(unit.Dp(4))
	r := image.Rectangle{Max: e.viewSize}
	r.Min.X -= pointerPadding
	r.Min.Y -= pointerPadding
	r.Max.X += pointerPadding
	r.Max.X += pointerPadding
	pointer.Rect(r).Add(gtx.Ops)
	e.scroller.Add(gtx.Ops)
	e.clicker.Add(gtx.Ops)
	e.caretOn = false
	if e.focused {
		now := gtx.Now()
		dt := now.Sub(e.blinkStart)
		blinking := dt < maxBlinkDuration
		const timePerBlink = time.Second / blinksPerSecond
		nextBlink := now.Add(timePerBlink/2 - dt%(timePerBlink/2))
		if blinking {
			redraw := op.InvalidateOp{At: nextBlink}
			redraw.Add(gtx.Ops)
		}
		e.caretOn = e.focused && (!blinking || dt%timePerBlink < timePerBlink/2)
	}

	gtx.Dimensions = layout.Dimensions{Size: e.viewSize, Baseline: e.dims.Baseline}
}
