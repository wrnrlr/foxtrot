package graphics

import (
	"gioui.org/f32"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
)

// Calculate the bounding box of a Graphics expression.
func boundingBox(expr *atoms.Expression) (bb f32.Rectangle) {
	for _, e := range expr.Parts[1:] {
		b := primitiveBox(e)
		if bb.Min.X > b.Min.X {
			bb.Min.X = b.Min.X
		}
		if bb.Min.Y > b.Min.Y {
			bb.Min.Y = b.Min.Y
		}
		if bb.Max.X < b.Max.X {
			bb.Max.X = b.Max.X
		}
		if bb.Max.Y < b.Max.Y {
			bb.Max.Y = b.Max.Y
		}
	}
	return bb
}

func primitiveBox(ex api.Ex) (bb f32.Rectangle) {
	expr, ok := ex.(*atoms.Expression)
	if !ok {
		return bb
	}

	switch expr.HeadStr() {
	case "System`Rectangle":
		bb = rectangleBox(expr)
	case "System`Circle":
		bb = rectangleBox(expr)
	case "System`Disk":
		bb = rectangleBox(expr)
	case "System`Line":
		bb = rectangleBox(expr)
	case "System`Triangle":
		bb = rectangleBox(expr)
	case "System`Polygon":
		bb = rectangleBox(expr)
	}

	return bb
}

func rectangleBox(expr *atoms.Expression) (bb f32.Rectangle) {
	if expr.Len() == 2 {
		bb.Max.X, bb.Max.Y = listToXY(expr.Parts[1])
	}
	return bb
}

func circleleBox(expr *atoms.Expression) (bb f32.Rectangle) {
	return bb
}

func diskBox(expr *atoms.Expression) (bb f32.Rectangle) {
	return bb
}

func lineBox(expr *atoms.Expression) (bb f32.Rectangle) {
	return bb
}

func triangleBox(expr *atoms.Expression) (bb f32.Rectangle) {
	return bb
}

func polygonBox(expr *atoms.Expression) (bb f32.Rectangle) {
	return bb
}

func listToXY(ex api.Ex) (x, y float32) {
	expr, ok := ex.(*atoms.Expression)
	if !ok {
		return
	}
	if expr.HeadStr() != "System`List" {
		return
	}
	if expr.Len() != 3 {
		return
	}
	x = exToNumber(expr.GetPart(1))
	y = exToNumber(expr.GetPart(1))
	return x, y
}

func exToNumber(ex api.Ex) float32 {
	switch e := ex.(type) {
	case *atoms.Integer:
		return float32(e.Val.Int64())
	case *atoms.Flt:
		f, _ := e.Val.Float32()
		return f
	case *atoms.Rational:
		num := float32(e.Num.Int64())
		den := float32(e.Num.Int64())
		return num / den
	default:
		return 0
	}
}

type Canvas struct {
	Bounds f32.Rectangle
}

func (c *Canvas) drawRect() {}

func (c *Canvas) drawCircle() {}

func (c *Canvas) drawDisk() {}
