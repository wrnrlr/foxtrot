package graphics

import (
	"gioui.org/f32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformCoordinates(t *testing.T) {
	bbox := f32.Rectangle{Min: f32.Point{X: 0, Y: 0}, Max: f32.Point{X: 2, Y: 2}}
	ctx := context{BBox: bbox}
	assert.Equal(t, float32(1), ctx.x(1))
	assert.Equal(t, float32(1), ctx.y(1))

	bbox = f32.Rectangle{Min: f32.Point{X: -1, Y: -1}, Max: f32.Point{X: 1, Y: 1}}
	ctx = context{BBox: bbox}
	assert.Equal(t, float32(1), ctx.x(0))
	assert.Equal(t, float32(1), ctx.y(0))

	bbox = f32.Rectangle{Min: f32.Point{X: 1, Y: 1}, Max: f32.Point{X: 3, Y: 3}}
	ctx = context{BBox: bbox}
	assert.Equal(t, float32(1), ctx.x(2))
	assert.Equal(t, float32(1), ctx.y(2))
}

func TestSize(t *testing.T) {
	bbox := f32.Rectangle{Min: f32.Point{X: 0, Y: 0}, Max: f32.Point{X: 2, Y: 2}}
	ctx := context{BBox: bbox}
	assert.Equal(t, float32(2), ctx.width())
	assert.Equal(t, float32(2), ctx.height())

	bbox = f32.Rectangle{Min: f32.Point{X: -1, Y: -1}, Max: f32.Point{X: 1, Y: 1}}
	ctx = context{BBox: bbox}
	assert.Equal(t, float32(2), ctx.width())
	assert.Equal(t, float32(2), ctx.height())

	bbox = f32.Rectangle{Min: f32.Point{X: 1, Y: 1}, Max: f32.Point{X: 3, Y: 3}}
	ctx = context{BBox: bbox}
	assert.Equal(t, float32(2), ctx.width())
	assert.Equal(t, float32(2), ctx.height())
}
