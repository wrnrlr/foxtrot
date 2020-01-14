package graphics

import (
	"gioui.org/f32"
	"github.com/corywalker/expreduce/expreduce/atoms"
)

func toTriangle(e *atoms.Expression) (*Triangle, error) {
	return nil, nil
}

type Triangle struct {
	p1, p2, p3 f32.Point
}
