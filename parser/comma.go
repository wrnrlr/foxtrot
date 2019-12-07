package parser

// CommaOpt represents data reduced by productions:
//
//	CommaOpt:
//	        /* empty */  // Case 0
//	|       ','          // Case 1
type CommaOpt struct {
	Token Token
}

func (n *CommaOpt) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *CommaOpt) String() string {
	return prettyString(n)
}

func (n *CommaOpt) Pos() Position {
	if n == nil {
		return Position{}
	}

	return n.Token.Pos
}
