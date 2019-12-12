package graphics

import (
	"errors"
	"fmt"
	"gioui.org/layout"
	"github.com/corywalker/expreduce/expreduce/atoms"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"image/color"
)

type ContinuousSeries struct {
	Style Style

	XValues []float32
	YValues []float32
}

func listToPoint(expr api.Ex, gtx *layout.Context) (float32, float32, error) {
	list, isList := atoms.HeadAssertion(expr, "System`List")
	if !isList {
		return 0, 0, fmt.Errorf("expected point to be list")
	}
	if list.Len() != 2 {
		return 0, 0, fmt.Errorf("points should be of length 2")
	}
	xFlt, xIsFlt := list.GetPart(1).(*atoms.Flt)
	yFlt, yIsFlt := list.GetPart(2).(*atoms.Flt)
	if !xIsFlt || !yIsFlt {
		return 0, 0, fmt.Errorf("point coordinates should be floats")
	}
	x, _ := xFlt.Val.Float64()
	y, _ := yFlt.Val.Float64()
	return float32(x), float32(y), nil
}

func renderLine(line *atoms.Expression, style *Style, gtx *layout.Context) error {
	if line.Len() != 1 {
		return fmt.Errorf("expected Line to have one argument, but it has %v arguments", line.Len())
	}
	points, ok := atoms.HeadAssertion(line.GetPart(1), "System`List")
	if !ok {
		return errors.New("expected a nested list")
	}

	series := ContinuousSeries{}
	series.Style = *style
	for _, pointExpr := range points.GetParts()[1:] {
		x, y, err := listToPoint(pointExpr, gtx)
		if err != nil {
			return err
		}
		series.XValues = append(series.XValues, x)
		series.YValues = append(series.YValues, y)
	}
	//graph.Series = append(graph.Series, series)
	return nil
}

func setColor(colorExpr *atoms.Expression, c *color.RGBA, gtx *layout.Context) error {
	if colorExpr.Len() != 3 {
		return fmt.Errorf("expected an RGBColor with 3 arguments")
	}
	for i := 0; i < 3; i++ {
		asFlt, isFlt := colorExpr.GetPart(i + 1).(*atoms.Flt)
		if !isFlt {
			return fmt.Errorf("the RGBColor should have floating point arguments")
		}
		fltVal, _ := asFlt.Val.Float32()
		intVal := uint8(fltVal * 255)
		switch i {
		case 0:
			c.R = intVal
		case 1:
			c.G = intVal
		case 2:
			c.B = intVal
		}
	}
	c.A = 255
	return nil
}

func setOpacity(opacityExpr *atoms.Expression, c *color.RGBA) error {
	if opacityExpr.Len() != 1 {
		return fmt.Errorf("expected an Opacity with 1 argument")
	}
	asFlt, isFlt := opacityExpr.GetPart(1).(*atoms.Flt)
	if !isFlt {
		return fmt.Errorf("the Opacity should have a floating point argument")
	}
	fltVal, _ := asFlt.Val.Float32()
	intVal := uint8(fltVal * 255)
	c.A = intVal
	return nil
}

func setAbsoluteThickness(thicknessExpr *atoms.Expression, thickness *float32, gtx *layout.Context) error {
	if thicknessExpr.Len() != 1 {
		return fmt.Errorf("expected an AbsoluteThickness with 1 argument")
	}
	asFlt, isFlt := thicknessExpr.GetPart(1).(*atoms.Flt)
	if !isFlt {
		return fmt.Errorf("the AbsoluteThickness should have a floating point argument")
	}
	fltVal, _ := asFlt.Val.Float64()
	*thickness = float32(fltVal)
	return nil
}

func setDirective(dir api.Ex, style *Style, gtx *layout.Context) error {
	//rgbColor, isRgbColor := atoms.HeadAssertion(dir, "System`RGBColor")
	//if isRgbColor {
	//	err := setColor(rgbColor, style.StrokeColor, gtx)
	//	return err
	//}
	//opacity, isOpacity := atoms.HeadAssertion(dir, "System`Opacity")
	//if isOpacity {
	//	err := setOpacity(opacity, style.StrokeColor)
	//	return err
	//}
	//thickness, isThickness := atoms.HeadAssertion(dir, "System`AbsoluteThickness")
	//if isThickness {
	//	err := setAbsoluteThickness(thickness, &style.StrokeWidth, gtx)
	//	return err
	//}
	//fmt.Printf("Skipping over unknown Directive: %v\n", dir)
	return nil
}

func renderPrimitive(primitive api.Ex, style *Style, gtx *layout.Context) error {
	line, isLine := atoms.HeadAssertion(primitive, "System`Line")
	list, isList := atoms.HeadAssertion(primitive, "System`List")
	directive, isDirective := atoms.HeadAssertion(primitive, "System`Directive")
	if isList {
		newStyle := *style
		for _, subPrimitive := range list.GetParts()[1:] {
			err := renderPrimitive(subPrimitive, &newStyle, gtx)
			if err != nil {
				fmt.Printf("Skipping over Primitive: %v\n", primitive)
			}
		}
	} else if isDirective {
		for _, directivePart := range directive.GetParts()[1:] {
			err := setDirective(directivePart, style, gtx)
			if err != nil {
				fmt.Printf("Skipping over Primitive: %v\n", primitive)
			}
		}
	} else if isLine {
		err := renderLine(line, style, gtx)
		if err != nil {
			fmt.Printf("Skipping over Primitive: %v\n", primitive)
		}
	} else {
		fmt.Printf("Skipping over Primitive: %v\n", primitive)
	}
	return nil
}
