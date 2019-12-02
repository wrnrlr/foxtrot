package cell

type CellType int

const (
	InputCell CellType = iota
	OutputCell
	TitleCell
	SectionCell
	SubSectionCell
	SubSubSectionCell
	TextCell
	CodeCell
)

var CellTypeNames = []string{"Input", "Output", "Title", "Section", "SubSection", "SubSubSection", "Text", "Code"}

func (d CellType) String() string {
	return CellTypeNames[d]
}

func (d CellType) Level() int {
	switch d {
	case TitleCell:
		return 1
	case SectionCell:
		return 3
	case SubSectionCell:
		return 4
	case SubSubSectionCell:
		return 5
	default:
		return 0
	}
}
