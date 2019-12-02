package notebook

import (
	"bytes"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
	"github.com/wrnrlr/foxtrot/cell"
	"testing"
)

func TestNewNotebook(t *testing.T) {

}

func TestDeleteCell(t *testing.T) {

}

const (
	helloTitle  = `Cell["\<Hello\>","Title"]`
	aPlusOneIn  = `Cell[BoxData[RowBox[{"a","+","1"}]],"Input",CellLabel -> "In[1]:= "]`
	aPlusOneOut = `Cell[BoxData[RowBox[{"1","+","a"}],StandardForm],"Output",CellLabel -> "Out[1]= "]`
	inAndOut    = `CompoundExpression[
		Cell[BoxData[RowBox[List["a", "+", "1"]]], "Input", Rule[CellLabel, "In[1]:= "]],
		Cell[BoxData[RowBox[List["1", "+", "a"]], StandardForm], "Output", Rule[CellLabel, "Out[1]= "]]]
`
)

func TestRead(t *testing.T) {
	src := parser.ReplaceSyms(aPlusOneIn + ";" + aPlusOneOut)
	buf := bytes.NewBufferString(src)
	kernel := expreduce.NewEvalState()
	expOut, _ := parser.InterpBuf(buf, "nofile", kernel)
	fmt.Printf("%v\n", expOut)
	// get cells
	ex, _ := expOut.(*atoms.Expression)
	if isCompoundExpression(ex.HeadStr()) {
		// data box

		// rowbox

		// cell type

		// cell label
	}
}

func isCompoundExpression(s string) bool {
	return s == "System`CompoundExpression"
}

func isCell(s string) bool {
	return s == "System`Cell"
}

func isDataBox(s string) bool {
	return s == "System`DataBox"
}

func isRule(s string) bool {
	return s == "System`Rule"
}

func cellType(s string) cell.CellType {
	switch s {
	case "Input":
		return cell.InputCell
	case "Output":
		return cell.InputCell
	case "Title":
		return cell.TitleCell
	case "Section":
		return cell.SectionCell
	case "Subsection":
		return cell.SubSectionCell
	case "Subsubsection":
		return cell.SubSubSectionCell
	case "Code":
		return cell.CodeCell
	case "Text":
		return cell.TextCell
	default:
		return cell.InputCell
	}
}
