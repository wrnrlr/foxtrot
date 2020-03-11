package nbx

import (
	"encoding/xml"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/cell"
	"io"
	"os"
)

const (
	title       = `cell["\<Hello\>","H5"]`
	aPlusOneIn  = `cell[BoxData[RowBox[{"a","+","1"}]],"Input",CellLabel -> "Content[1]:= "]`
	aPlusOneOut = `cell[BoxData[RowBox[{"1","+","a"}],StandardForm],"Output",CellLabel -> "Out[1]= "]`
	inAndOut    = `CompoundExpression[
		cell[BoxData[RowBox[List["a", "+", "1"]]], "Input", Rule[CellLabel, "Content[1]:= "]],
		cell[BoxData[RowBox[List["1", "+", "a"]], StandardForm], "Output", Rule[CellLabel, "Out[1]= "]]]
`
)

const temp = `<cell type="%s">"%s"</cell>`

// Write cells to filename.nbx
func WriteFile(filename string, cells cell.Cells) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0660)
	defer file.Close()
	if err != nil {
		return err
	}
	file.Truncate(0)
	file.Seek(0, 0)
	return Write(file, cells)
}

func Write(w io.Writer, cells cell.Cells) error {
	nb := &notebookTag{}
	nb.XMLName = xml.Name{Local: "notebook"}
	for _, c := range cells {
		cell := cellTag{}
		cell.Type = c.Type().String()
		cell.Content = c.Text()
		cell.XMLName = xml.Name{Local: "cell"}
		nb.Cells = append(nb.Cells, cell)
	}
	b, err := xml.MarshalIndent(nb, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func WriteCells(w io.Writer, cells cell.Cells) expreduceapi.Ex {
	res := atoms.E(atoms.S("List"))

	//keys := make([]streamKey, 0)
	//for k := range sm.openStreams {
	//	keys = append(keys, k)
	//}
	//sort.Slice(keys, func(i, j int) bool {
	//	return keys[i].id < keys[j].id
	//})
	//
	//for _, k := range keys {
	//	res.AppendEx(atoms.E(
	//		atoms.S("OutputStream"),
	//		atoms.NewString(k.name),
	//		atoms.NewInt(k.id),
	//	))
	//}

	return res
}
