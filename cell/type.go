package cell

type Type int

const (
	Empty Type = iota
	Input
	Output
	H1
	H2
	H3
	H4
	H5
	H6
	Paragraph
	Code
	Inline
)

var CellTypeNames = []string{"Input", "Output", "H1", "H2", "H3", "H4", "H5", "H6", "Paragraph", "Code"}

func (d Type) String() string {
	return CellTypeNames[d]
}

func ParseType(s string) Type {
	switch s {
	case "Input":
		return Input
	case "Output":
		return Output
	case "H1":
		return H2
	case "H2":
		return H2
	case "H3":
		return H3
	case "H4":
		return H4
	case "H5":
		return H5
	case "H6":
		return H6
	case "Paragraph":
		return Paragraph
	case "Code":
		return Code
	default:
		return Empty
	}
}
