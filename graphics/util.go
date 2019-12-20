package graphics

import (
	"errors"
	"gioui.org/f32"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
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
