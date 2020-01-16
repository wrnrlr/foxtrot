package graphics

import (
	"gioui.org/f32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAngle(t *testing.T) {
	p1 := f32.Point{X: 0, Y: 0}
	p2 := f32.Point{X: 1, Y: 1}
	assert.Equal(t, float32(45), angle(p1, p2))
	p1 = f32.Point{X: 0, Y: 0}
	p2 = f32.Point{X: 0, Y: 1}
	assert.Equal(t, float32(90), angle(p1, p2))
	p1 = f32.Point{X: 0, Y: 0}
	p2 = f32.Point{X: 1, Y: 0}
	assert.Equal(t, float32(0), angle(p1, p2))
	p1 = f32.Point{X: 0, Y: 0}
	p2 = f32.Point{X: 1, Y: -1}
	assert.Equal(t, float32(-45), angle(p1, p2))
}

func TestOffset(t *testing.T) {
	p1 := f32.Point{X: 0, Y: 0}
	p2 := offset(p1, -5, 135)
	assert.Equal(t, p1, p2)
}

func TestLineAngles(t *testing.T) {
	line := []f32.Point{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 0}}
	expected := []float32{90}
	assert.Equal(t, expected, toAngles(line))

	line = []f32.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}
	expected = []float32{-180}
	assert.Equal(t, expected, toAngles(line))

	//angles := toAngles(line)
	//assert.Equal(t, 2, len(angles))
	//assert.Equal(t, float32(90), angles[0])
	//assert.Equal(t, float32(270), angles[1])
	//assert.Equal(t, expected, angles)
}

func TestAcross(t *testing.T) {
	p := f32.Point{X: 0, Y: 0}
	q := f32.Point{X: 1, Y: 0}
	r := f32.Point{X: 0, Y: 1}

	assert.Equal(t, float32(90), across(p, q, r))

	p = f32.Point{X: 0, Y: 0}
	q = f32.Point{X: 0, Y: 0}
	r = f32.Point{X: 0, Y: 0}

	assert.Equal(t, float32(0), across(p, q, r))

	//angles := toAngles(line)
	//assert.Equal(t, 2, len(angles))
	//assert.Equal(t, float32(90), angles[0])
	//assert.Equal(t, float32(270), angles[1])
	//assert.Equal(t, expected, angles)
}

func TestRelativePoints(t *testing.T) {
	line := []f32.Point{{X: 10, Y: 10}, {X: 100, Y: 100}, {X: 0, Y: 200}}
	deltas := relativePoints(line)
	assert.Equal(t, []f32.Point{{10, 10}, {90, 90}, {-90, 110}}, deltas)
}
