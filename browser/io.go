package browser

import (
	"bytes"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
)

func Read(s string) {
	src := parser.ReplaceSyms(s)
	buf := bytes.NewBufferString(src)
	kernel := expreduce.NewEvalState()
	expOut, _ := parser.InterpBuf(buf, "nofile", kernel)
	fmt.Printf("%v\n", expOut)
	// get cells
	ex, _ := expOut.(*atoms.Expression)
	if ex.HeadStr() == "" {
		// data box

		// rowbox

		// cell type

		// cell label
	}
}
