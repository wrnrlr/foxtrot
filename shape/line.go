package shape

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"github.com/wrnrlr/foxtrot/util"
	"math"
)

const (
	rad45  = float32(45 * math.Pi / 180)
	rad135 = float32(135 * math.Pi / 180)
	rad315 = float32(315 * math.Pi / 180)
	rad225 = float32(225 * math.Pi / 180)
	rad90  = float32(90 * math.Pi / 180)
	rad180 = float32(180 * math.Pi / 180)
)

type Line []f32.Point

func (l Line) Stroke(width unit.Value, gtx *layout.Context) (box f32.Rectangle) {
	if len(l) < 2 {
		return box
	}
	var path clip.Path
	path.Begin(gtx.Ops)
	distance := float32(gtx.Px(width))
	var angles []float32
	var offsetPoints, originalPoints, deltaPoints []f32.Point
	var tilt float32
	var prevDelta f32.Point
	for i, point := range l {
		if i == 0 {
			nextPoint := l[i+1]
			tilt = angle(point, nextPoint) + rad225
		} else if i == len(l)-1 {
			prevPoint := l[i-1]
			tilt = angle(prevPoint, point) + rad315
		} else {
			prevPoint := l[i-1]
			nextPoint := l[i+1]
			tilt = bezel(point, prevPoint, nextPoint)
		}
		angles = append(angles, tilt)
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
		if i == 0 {
			path.Move(newPoint)
		} else {
			path.Line(newPoint)
		}
	}
	for i := len(l) - 1; i >= 0; i-- {
		point := l[i]
		if i == 0 {
			nextPoint := l[i+1]
			tilt = angle(point, nextPoint) + rad135
		} else if i == len(l)-1 {
			prevPoint := l[i-1]
			tilt = angle(prevPoint, point) + rad45
		} else {
			point := l[i]
			prevPoint := l[i-1]
			nextPoint := l[i+1]
			tilt = bezel(point, nextPoint, prevPoint)
		}
		angles = append(angles, tilt)
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
		path.Line(newPoint)
	}
	point := l[0]
	nextPoint := l[1]
	tilt = angle(point, nextPoint) + rad225
	angles = append(angles, tilt)
	originalPoints = append(originalPoints, point)
	point = offsetPoint(point, distance, tilt)
	offsetPoints = append(offsetPoints, point)
	point = point.Sub(prevDelta)
	path.Line(point)
	deltaPoints = append(deltaPoints, point)
	fmt.Printf("Original Points: %v\n", originalPoints)
	printDegrees(angles)
	fmt.Printf("Offset Points:   %v\n", offsetPoints)
	for _, p := range offsetPoints {
		box.Min.X = util.Min(box.Min.X, p.X)
		box.Min.Y = util.Min(box.Min.Y, p.Y)
		box.Max.X = util.Max(box.Max.X, p.X)
		box.Max.Y = util.Max(box.Max.Y, p.Y)
	}
	fmt.Printf("Min and Max:   %v\n", box)
	fmt.Printf("Delta Points:    %v\n", deltaPoints)
	path.End().Add(gtx.Ops)
	return box
}
