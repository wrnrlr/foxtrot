package style

import (
	"gioui.org/text"
	"image/color"
)

type Style struct {
	// Color is the text color.
	Color color.RGBA
	Font  text.Font
	// Hint contains the text displayed when the editor is empty.
	Hint string
	// HintColor is the color of hint text.
	HintColor  color.RGBA
	CaretColor color.RGBA

	Shaper *text.Shaper
}
