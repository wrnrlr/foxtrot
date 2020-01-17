package plot

type Fn func(x float32) float32

type Function struct {
	F Fn
}
