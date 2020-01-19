package plotter

import "math"

func Min(x, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

func Max(x, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}

func Inf(sign int) float32 {
	return float32(math.Inf(sign))
}

func IsNaN(sign float32) bool {
	return math.IsNaN(float64(sign))
}

func IsInf(f float32, sign int) bool {
	return math.IsInf(float64(f), sign)
}
