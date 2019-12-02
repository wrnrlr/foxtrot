package typeset

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
)

func Plus(left, right Shape) Shape {
	return &Operator{PlusSymbol, left, right}
}

func Minus(left, right Shape) Shape {
	return &Operator{MinusSymbol, left, right}
}

func Multiply(left, right Shape) Shape {
	return &Operator{MultiplySymbol, left, right}
}

func Modulo(left, right Shape) Shape {
	return &Operator{ModuloSymbol, left, right}
}

func Factor(left, right Shape) Shape {
	return &Operator{Symbol: FactorSymbol, Left: left}
}

func Power(base, exponent Shape) Shape {
	return &Word{Content: base, Superscript: exponent}
}

type Operator struct {
	Symbol, Left, Right Shape
}

func (o *Operator) Dimensions(c *layout.Context, s *text.Shaper, font text.Font) layout.Dimensions {
	dims := o.Left.Dimensions(c, s, font)
	d := o.Symbol.Dimensions(c, s, font)
	dims.Size.X += d.Size.X
	//dims.Size.Y += d.Size.Y
	if o.Right != nil {
		d = o.Right.Dimensions(c, s, font)
		dims.Size.X += d.Size.X
		//dims.Size.Y += d.Size.Y
	}
	return dims
}

func (o *Operator) Layout(gtx *layout.Context, s *text.Shaper, font text.Font) {
	dims := o.Dimensions(gtx, s, font)
	var stack op.StackOp
	offset := f32.Point{X: 0, Y: 0}
	stack.Push(gtx.Ops)
	o.Left.Layout(gtx, s, font)
	offset.X += float32(gtx.Dimensions.Size.X)
	stack.Pop()

	stack.Push(gtx.Ops)
	od := o.Symbol.Dimensions(gtx, s, font)
	offset.Y = float32((dims.Size.Y - od.Size.Y) / 2)
	op.TransformOp{}.Offset(offset).Add(gtx.Ops)
	o.Symbol.Layout(gtx, s, font)
	offset.X += float32(gtx.Dimensions.Size.X)
	offset.Y = 0
	stack.Pop()

	if o.Right != nil {
		stack.Push(gtx.Ops)
		rd := o.Right.Dimensions(gtx, s, font)
		offset.Y = float32((dims.Size.Y - rd.Size.Y) / 2)
		op.TransformOp{}.Offset(offset).Add(gtx.Ops)
		o.Right.Layout(gtx, s, font)
		//offset.X += float32(gtx.Dimensions.Size.X)
		//offset.Y += float32(gtx.Dimensions.Size.Y)
		stack.Pop()
	}
}
