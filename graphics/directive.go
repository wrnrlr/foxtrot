package graphics

type Directive interface {
	Set(style *Style)
}

type Thickness struct {
	thickness float32
}

func (t Thickness) Set(style *Style) {
	style.Thickness = t.thickness
}

type CMYKColor struct{}
