package editor

import (
	"gioui.org/f32"
	"gioui.org/op/clip"
	"gioui.org/text"
	"golang.org/x/image/math/fixed"
	"image"
	"unicode/utf8"
)

type line struct {
	offset f32.Point
	clip   clip.Op
}

type lineIterator struct {
	Lines     []text.Line
	Clip      image.Rectangle
	Alignment text.Alignment
	Width     int
	Offset    image.Point

	y, prevDesc fixed.Int26_6
}

const inf = 1e6

func (l *lineIterator) Next() (text.String, f32.Point, bool) {
	for len(l.Lines) > 0 {
		line := l.Lines[0]
		l.Lines = l.Lines[1:]
		x := align(l.Alignment, line.Width, l.Width) + fixed.I(l.Offset.X)
		l.y += l.prevDesc + line.Ascent
		l.prevDesc = line.Descent
		// Align baseline and line start to the pixel grid.
		off := fixed.Point26_6{X: fixed.I(x.Floor()), Y: fixed.I(l.y.Ceil())}
		l.y = off.Y
		off.Y += fixed.I(l.Offset.Y)
		if (off.Y + line.Bounds.Min.Y).Floor() > l.Clip.Max.Y {
			break
		}
		if (off.Y + line.Bounds.Max.Y).Ceil() < l.Clip.Min.Y {
			continue
		}
		str := line.Text
		for len(str.Advances) > 0 {
			adv := str.Advances[0]
			if (off.X + adv + line.Bounds.Max.X - line.Width).Ceil() >= l.Clip.Min.X {
				break
			}
			off.X += adv
			_, s := utf8.DecodeRuneInString(str.String)
			str.String = str.String[s:]
			str.Advances = str.Advances[1:]
		}
		n := 0
		endx := off.X
		for i, adv := range str.Advances {
			if (endx + line.Bounds.Min.X).Floor() > l.Clip.Max.X {
				str.String = str.String[:n]
				str.Advances = str.Advances[:i]
				break
			}
			_, s := utf8.DecodeRuneInString(str.String[n:])
			n += s
			endx += adv
		}
		offf := f32.Point{X: float32(off.X) / 64, Y: float32(off.Y) / 64}
		return str, offf, true
	}
	return text.String{}, f32.Point{}, false
}
