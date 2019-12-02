package nbx

const (
	title       = `Cell["\<Hello\>","Title"]`
	aPlusOneIn  = `Cell[BoxData[RowBox[{"a","+","1"}]],"Input",CellLabel -> "In[1]:= "]`
	aPlusOneOut = `Cell[BoxData[RowBox[{"1","+","a"}],StandardForm],"Output",CellLabel -> "Out[1]= "]`
	inAndOut    = `CompoundExpression[
		Cell[BoxData[RowBox[List["a", "+", "1"]]], "Input", Rule[CellLabel, "In[1]:= "]],
		Cell[BoxData[RowBox[List["1", "+", "a"]], StandardForm], "Output", Rule[CellLabel, "Out[1]= "]]]
`
)

const temp = `<cell type="%s">"%s"</cell>`

func Read2(data []byte) error {
	//cells := make([]*cell.Cell, 0)
	//nb := &notebookTag{}
	//return xml.Unmarshal(data, nb)
	//for _, c := range cells {
	//	nb.cells = append(nb.cells, cellTag{c.Type.String(), c.Text()})
	//}
	return nil
}
