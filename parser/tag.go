package parser

import "fmt"

// TagCase represents case numbers of production Tag
type TagCase int

// Values of type TagCase
const (
	TagIdent TagCase = iota
	TagString
)

// String implements fmt.Stringer
func (n TagCase) String() string {
	switch n {
	case TagIdent:
		return "TagIdent"
	case TagString:
		return "TagString"
	default:
		return fmt.Sprintf("TagCase(%v)", int(n))
	}
}

type Tag struct {
	Case  TagCase
	Token Token
}

func (n *Tag) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tag) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tag) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}
