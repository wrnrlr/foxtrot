package nbx

import (
	"bytes"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
	"testing"
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
