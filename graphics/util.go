package graphics

import (
	"errors"
	"gioui.org/f32"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"math"
)

func toFloat(e expreduceapi.Ex) (float32, error) {
	i, isInt := e.(*atoms.Integer)
	if isInt {
		f := float32(i.Val.Int64())
		return f, nil
	}
	f, isFlt := e.(*atoms.Flt)
	if isFlt {
		i, _ := f.Val.Float32()
		return i, nil
	}
	return 0, errors.New("Connot be converted to a float")
}

func toPoint(e expreduceapi.Ex) (p f32.Point, err error) {
	expr, isExpr := e.(*atoms.Expression)
	if !isExpr {
		return p, errors.New("point should be a list {x,y}")
	}
	if expr.HeadStr() != "System`List" {
		return p, errors.New("point should be a list {x,y}")
	}
	if expr.Len() != 2 {
		return p, errors.New("point should be a list {x,y}")
	}
	p.X, err = toFloat(expr.GetPart(1))
	if err != nil {
		return p, err
	}
	p.Y, err = toFloat(expr.GetPart(2))
	if err != nil {
		return p, err
	}
	return p, nil
}

func toPoints(e expreduceapi.Ex) (points []f32.Point, err error) {
	expr, isExpr := e.(*atoms.Expression)
	if !isExpr {
		return nil, errors.New("point should be a list {x,y}")
	}
	if expr.HeadStr() != "System`List" {
		return nil, errors.New("point should be a list {x,y}")
	}
	for _, ex := range expr.Parts[1:] {
		p, err := toPoint(ex)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

func toAngles(points []f32.Point) []float32 {
	var angles []float32
	for i, point := range points {
		if i != 0 && i != len(points)-1 {
			prevPoint := points[i-1]
			nextPoint := points[i+1]
			a := across(point, nextPoint, prevPoint)
			angles = append(angles, a)
		}
	}
	//for _, a := range angles {
	//	angles = append(angles, mod(a+180, 360))
	//}
	return angles
}

func angle(p1, p2 f32.Point) float32 {
	return float32(math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X))) * 180 / math.Pi
}

func cos(v float32) float32 {
	return float32(math.Cos(float64(v)))
}

func sin(v float32) float32 {
	return float32(math.Sin(float64(v)))
}

func offset(point f32.Point, distance, angle float32) f32.Point {
	x := point.X + distance*cos(angle)
	y := point.Y + distance*sin(angle)
	return f32.Point{X: x, Y: y}
}

func mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

// Return angle in degrees between the vectors pq and pr
func across(p, q, r f32.Point) float32 {
	return (atan2(q.Y-p.Y, q.X-p.X) - atan2(r.Y-p.Y, r.X-p.X)) * 180 / math.Pi
}

func atan2(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

// Return points with relative coordinates based on delta between two consecutive points
func relativePoints(points []f32.Point) []f32.Point {
	deltas := make([]f32.Point, len(points))
	var prev f32.Point
	for i, p := range points {
		prev = p.Sub(prev)
		deltas[i] = prev
	}
	return deltas
}

func min(n, m float32) float32 {
	return float32(math.Min(float64(n), float64(m)))
}

func max(n, m float32) float32 {
	return float32(math.Min(float64(n), float64(m)))
}
