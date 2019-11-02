package foxtrot

import (
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
)

func NewKernel() *expreduce.EvalState {
	es := expreduce.NewEvalState()
	return es
}

func expressionToString(es *expreduce.EvalState, exp expreduceapi.Ex, promptCount int) string {
	var res string
	isNull := false
	asSym, isSym := exp.(*atoms.Symbol)
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
			asSpecialForm, isSpecialForm := atoms.HeadAssertion(exp, specialForm)
			if !isSpecialForm {
				continue
			}
			if len(asSpecialForm.Parts) != 2 {
				continue
			}
			res = fmt.Sprintf(
				"//%s= %s",
				specialForm[7:],
				asSpecialForm.Parts[1].StringForm(
					expreduce.ActualStringFormArgsFull(specialForm[7:], es)),
			)
			wasSpecialForm = true
		}
		if !wasSpecialForm {
			res = fmt.Sprintf("%s", exp.StringForm(expreduce.ActualStringFormArgsFull("InputForm", es)))
		}
	}
	return res
}
