package shape

import (
	"fmt"
	"gioui.org/f32"
	"github.com/wrnrlr/foxtrot/util"
	"math"
)

func slope(p1, p2 f32.Point) float32 {
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func angle(p1, p2 f32.Point) float32 {
	return float32(math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X)))
}

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

//func pointAngles(points []f32.Point) []float32 {
//	var angles []float32
//	for i, point := range points {
//		if i == 0 {
//			angles = append(angles, rad225)
//		} else if i == len(points)-1 {
//			angles = append(angles, rad315)
//		} else {
//			prevPoint := points[i-1]
//			nextPoint := points[i+1]
//			a := across(point, nextPoint, prevPoint) - rad90
//			angles = append(angles, a)
//		}
//	}
//	for i := len(points) - 1; i >= 0; i-- {
//		if i == 0 {
//			angles = append(angles, rad135)
//		} else if i == len(points)-1 {
//			angles = append(angles, rad45)
//		} else {
//			point := points[i]
//			prevPoint := points[i-1]
//			nextPoint := points[i+1]
//			a := across(point, nextPoint, prevPoint) + rad90
//			angles = append(angles, a)
//		}
//	}
//	return angles
//}

func offsetPoint(point f32.Point, distance, angle float32) f32.Point {
	//fmt.Printf("Point X: %f, Y: %f, Angle: %f\n", point.X, point.Y, angle)
	x := point.X + distance*cos(angle)
	y := point.Y + distance*sin(angle)
	//fmt.Printf("Point X: %f, Y: %f \n", x, y)
	return f32.Point{X: x, Y: y}
}

//func pointOffsets(line []f32.Point, distance float32) []f32.Point {
//	fmt.Printf("==================================================================\n")
//	fmt.Printf("Original Points: %v\n", line)
//	angles := pointAngles(line)
//	//fmt.Println(angles)
//	var pps []f32.Point
//	j := 0
//	for i, p := range line {
//		a := angles[i]
//		pps = append(pps, offsetPoint(p, distance, a))
//		j++
//	}
//	lastIndex := len(pps) - 1
//	for i := lastIndex; i >= 0; i-- {
//		p := line[i]
//		a := angles[j]
//		pps = append(pps, offsetPoint(p, distance, a))
//		j++
//	}
//	pps = append(pps, pps[0])
//	fmt.Printf("Delta Points:    %v\n", pps)
//	pos := pointDeltas(pps)
//	fmt.Printf("Offset Points:   %v\n", pos)
//	fmt.Printf("==================================================================\n")
//	return pos
//}

// Return angle in radian between the vectors pq and pr
func across(p, q, r f32.Point) float32 {
	return atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
}

func bezel(p, q, r f32.Point) float32 {
	angle := atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)
	if angle < -rad180 || angle > rad180 { // concave

	} else { // convex

	}
	if angle < 0 {
		angle = angle
	} else {
		angle = angle
	}
	return angle
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

func printDegrees(radials []float32) {
	var degrees []float32
	for _, a := range radials {
		degrees = append(degrees, mod(a*180/math.Pi, 360))
	}
	fmt.Printf("Angles: %v\n", degrees)
}

func Overlay(rs ...f32.Rectangle) f32.Rectangle {
	rr := rs[0]
	for _, r := range rs[1:] {
		rr = f32.Rectangle{
			Min: f32.Point{util.Min(rr.Min.Y, r.Min.Y), util.Min(rr.Min.Y, r.Min.Y)},
			Max: f32.Point{util.Max(rr.Max.Y, r.Max.Y), util.Max(rr.Max.Y, r.Max.Y)},
		}
	}
	return rr
}
