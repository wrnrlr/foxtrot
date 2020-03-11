package main

import (
	"bytes"
	"fmt"
	"gioui.org/app/headless"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
	api "github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/app"
	"github.com/wrnrlr/foxtrot/colors"
	"github.com/wrnrlr/foxtrot/output"
	"github.com/wrnrlr/foxtrot/style"
	"image"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"time"
)

//var screenshot = flag.String("screenshot", "", "save a screenshot to a file and exit")
//var width = flag.String("screenshot", "", "save a screenshot to a file and exit")
//var height = flag.String("screenshot", "", "save a screenshot to a file and exit")

func main() {
	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	//if path == "version" {
	//	fmt.Printf("Foxtrot %s\n", foxtrot.Version)
	//} else if path == "save" {
	//	if err := saveOutput(*screenshot); err != nil {
	//		fmt.Fprintf(os.Stderr, "failed to save screenshot: %v\n", err)
	//		os.Exit(1)
	//	}
	//	os.Exit(0)
	//}
	app.RunUI(path)
}

var (
	fnt = text.Font{Size: unit.Sp(20)}
)

func saveOutput(f string) error {
	kernel := expreduce.NewEvalState()
	const scale = 1.5
	gtx := new(layout.Context)
	//gtx.Reset(&scaledConfig{scale}, sz)
	th := material.NewTheme()
	s := style.Style{
		Font:   text.Font{Size: unit.Sp(16)},
		Shaper: th.Shaper,
		Color:  colors.Black,
	}
	exOut := parser.Interp(f, kernel)
	exOut = kernel.Eval(exOut)
	//outTxt := formattedOutput(kernel, exOut, 0)
	o := output.FromEx(exOut, gtx)
	o.Layout(gtx, s)
	// Set window size based on o.Dimensions()
	dims := o.Dimensions(gtx, s)
	sz := image.Point{X: int(float32(dims.Size.X) * scale), Y: int(float32(dims.Size.Y) * scale)}
	w, err := headless.NewWindow(sz.X, sz.Y)
	if err != nil {
		return err
	}
	w.Frame(gtx.Ops)
	img, err := w.Screenshot()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	return ioutil.WriteFile(f, buf.Bytes(), 0666)
}

type scaledConfig struct {
	Scale float32
}

func (s *scaledConfig) Now() time.Time {
	return time.Now()
}

func (s scaledConfig) Px(v unit.Value) int {
	scale := s.Scale
	if v.U == unit.UnitPx {
		scale = 1
	}
	return int(math.Round(float64(scale * v.V)))
}

func formattedOutput(es *expreduce.EvalState, res api.Ex, promptNum int) (s string) {
	isNull := false
	asSym, isSym := res.(*atoms.Symbol)
	if isSym {
		if asSym.Name == "System`Null" {
			isNull = true
		}
	}

	if !isNull {
		// Print formatted result
		specialForms := []string{
			"System`FullForm",
			"System`OutputForm",
		}
		wasSpecialForm := false
		for _, specialForm := range specialForms {
			asSpecialForm, isSpecialForm := atoms.HeadAssertion(
				res, specialForm)
			if !isSpecialForm {
				continue
			}
			if len(asSpecialForm.Parts) != 2 {
				continue
			}
			s = fmt.Sprintf(
				"Out[%d]//%s= %s\n\n",
				promptNum,
				specialForm[7:],
				asSpecialForm.Parts[1].StringForm(
					expreduce.ActualStringFormArgsFull(specialForm[7:], es)))
			wasSpecialForm = true
		}
		if !wasSpecialForm {
			s = fmt.Sprintf("Out[%d]= %s\n\n", promptNum, res.StringForm(
				expreduce.ActualStringFormArgsFull("InputForm", es)))
		}
	}
	return s
}
