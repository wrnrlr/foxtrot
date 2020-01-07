package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
	"math"
)

const (
	rad45  = float32(45 * math.Pi / 180)
	rad135 = float32(135 * math.Pi / 180)
	rad315 = float32(315 * math.Pi / 180)
	rad225 = float32(225 * math.Pi / 180)
	rad90  = float32(90 * math.Pi / 180)
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	black = color.RGBA{0, 0, 0, 255}
	lines = [][]f32.Point{
		{{X: 10, Y: 10}, {X: 110, Y: 10}},
		//{{X: 110, Y: 10}, {X: 10, Y: 10}},
		//{{X: 10, Y: 10}, {X: 110, Y: 110}},
		//{{X: 110, Y: 110}, {X: 10, Y: 10}},
		//{{X: 10, Y: 10}, {X: 110, Y: 110}, {X: 210, Y: 10}},
	}
)

func main() {
	go func() {
		//var col uint8 = 127
		w := app.NewWindow()
		gtx := layout.NewContext(w.Queue())
		//list := &layout.List{Axis: layout.Vertical}
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx.Reset(e.Config, e.Size)
				paint.ColorOp{black}.Add(gtx.Ops)
				drawLine(lines[0], gtx)
				//list.Layout(gtx, len(lines), func(i int) {
				//	drawLine(lines[i], gtx)
				//})
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

func drawLine(points []f32.Point, gtx *layout.Context) {
	var stack op.StackOp
	stack.Push(gtx.Ops)
	line(gtx.Ops, points)
	paint.ColorOp{red}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{X: float32(600), Y: 200}}}.Add(gtx.Ops)
	stack.Pop()
	gtx.Constraints = layout.RigidConstraints(image.Point{X: 200, Y: 600})
}

func line(ops *op.Ops, points []f32.Point) {
	if len(points) < 2 {
		return
	}
	var path clip.Path
	path.Begin(ops)
	distance := float32(5)
	var offsetPoints, originalPoints, deltaPoints []f32.Point
	var tilt float32
	var prevDelta f32.Point
	for i, point := range points {
		if i == 0 {
			tilt = rad225
		} else if i == len(points)-1 {
			tilt = rad315
		} else {
			prevPoint := points[i-1]
			nextPoint := points[i+1]
			tilt = across(point, nextPoint, prevPoint)
		}
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
			tilt = rad135
		} else if i == len(points)-1 {
			tilt = rad45
		} else {
			point := points[i]
			prevPoint := points[i-1]
			nextPoint := points[i+1]
			tilt = across(point, nextPoint, prevPoint)
		}
		originalPoints = append(originalPoints, point)
		point = offsetPoint(point, distance, tilt)
		offsetPoints = append(offsetPoints, point)
		newPoint := point.Sub(prevDelta)
		deltaPoints = append(deltaPoints, newPoint)
		prevDelta = point
		path.Line(newPoint)
	}
	point := points[0]
	originalPoints = append(originalPoints, point)
	point = offsetPoint(point, distance, rad225)
	offsetPoints = append(offsetPoints, point)
	point = point.Sub(prevDelta)
	path.Line(point)
	deltaPoints = append(deltaPoints, point)
	fmt.Printf("Original Points: %v\n", originalPoints)
	fmt.Printf("Offset Points:   %v\n", offsetPoints)
	fmt.Printf("Delta Points:    %v\n", deltaPoints)
	path.End().Add(ops)
}

func offsetPoint(point f32.Point, distance, angle float32) f32.Point {
	//fmt.Printf("Point X: %f, Y: %f, Angle: %f\n", point.X, point.Y, angle)
	x := point.X + distance*cos(angle)
	y := point.Y + distance*sin(angle)
	//fmt.Printf("Point X: %f, Y: %f \n", x, y)
	return f32.Point{X: x, Y: y}
}

func slope(p1, p2 f32.Point) float32 {
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

//func angle(p1, p2 f32.Point) float32 {
//	return float32(math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X)))
//}

func cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func pointDeltas(points []f32.Point) []f32.Point {
	deltas := make([]f32.Point, len(points))
	var prev f32.Point
	for i, p := range points {
		deltas[i] = p.Sub(prev)
		prev = p
	}
	return deltas
}

func pointAngles(points []f32.Point) []float32 {
	var angles []float32
	for i, point := range points {
		if i == 0 {
			angles = append(angles, rad225)
		} else if i == len(points)-1 {
			angles = append(angles, rad315)
		} else {
			prevPoint := points[i-1]
			nextPoint := points[i+1]
			a := across(point, nextPoint, prevPoint)
			angles = append(angles, a)
		}
	}
	for i := len(points) - 1; i >= 0; i-- {
		if i == 0 {
			angles = append(angles, rad135)
		} else if i == len(points)-1 {
			angles = append(angles, rad45)
		} else {
			point := points[i]
			prevPoint := points[i-1]
			nextPoint := points[i+1]
			a := across(point, nextPoint, prevPoint)
			angles = append(angles, a)
		}
	}
	return angles
}

func pointOffsets(line []f32.Point, distance float32) []f32.Point {
	fmt.Printf("==================================================================\n")
	fmt.Printf("Original Points: %v\n", line)
	angles := pointAngles(line)
	//fmt.Println(angles)
	var pps []f32.Point
	j := 0
	for i, p := range line {
		a := angles[i]
		pps = append(pps, offsetPoint(p, distance, a))
		j++
	}
	lastIndex := len(pps) - 1
	for i := lastIndex; i >= 0; i-- {
		p := line[i]
		a := angles[j]
		pps = append(pps, offsetPoint(p, distance, a))
		j++
	}
	pps = append(pps, pps[0])
	fmt.Printf("Delta Points:    %v\n", pps)
	pos := pointDeltas(pps)
	fmt.Printf("Offset Points:   %v\n", pos)
	fmt.Printf("==================================================================\n")
	return pos
}

// Return angle in radian between the vectors pq and pr
func across(p, q, r f32.Point) float32 {
	return (atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X))
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}
