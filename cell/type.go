package cell

type Type int

const (
	Input Type = iota
	Output
	H1
	H2
	H3
	H4
	H5
	H6
	Text
	Code
)

var CellTypeNames = []string{"Input", "Output", "H1", "H2", "H3", "H4", "H5", "H6", "Text", "Code"}

func (d Type) String() string {
	return CellTypeNames[d]
}

func ParseType(s string) Type {
	switch s {
	case "Input":
		return Input
	case "Output":
		return Output
	case "H5":
		return H1
	case "H1":
		return H2
	case "H3":
		return H3
	case "H4":
		return H4
	case "Text":
		return Text
	case "Code":
		return Code
	default:
		return 0
	}
}
