package shape

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"image/color"
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

func (l Line) Add(ops *op.Ops) {

}

func (l Line) Stroke(width float32, style StrokeType, rgba color.RGBA, ops *op.Ops) {

}

func StrokeLine(points []f32.Point, lineWidth int, ops *op.Ops) {
	if len(points) < 2 {
		return
	}
	var path clip.Path
	path.Begin(ops)
	distance := float32(lineWidth)
	var angles []float32
	var offsetPoints, originalPoints, deltaPoints []f32.Point
	var tilt float32
	var prevDelta f32.Point
	for i, point := range points {
		if i == 0 {
			nextPoint := points[i+1]
			tilt = angle(point, nextPoint) + rad225
		} else if i == len(points)-1 {
			prevPoint := points[i-1]
			tilt = angle(prevPoint, point) + rad315
		} else {
			prevPoint := points[i-1]
			nextPoint := points[i+1]
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
	for i := len(points) - 1; i >= 0; i-- {
		point := points[i]
		if i == 0 {
			nextPoint := points[i+1]
			tilt = angle(point, nextPoint) + rad135
		} else if i == len(points)-1 {
			prevPoint := points[i-1]
			tilt = angle(prevPoint, point) + rad45
		} else {
			point := points[i]
			prevPoint := points[i-1]
			nextPoint := points[i+1]
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
	point := points[0]
	nextPoint := points[1]
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
	fmt.Printf("Delta Points:    %v\n", deltaPoints)
	path.End().Add(ops)
}
