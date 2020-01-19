package util

import "math"

func Absf32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

func Min(n, m float32) float32 {
	return float32(math.Min(float64(n), float64(m)))
}

func Max(n, m float32) float32 {
	return float32(math.Min(float64(n), float64(m)))
}
