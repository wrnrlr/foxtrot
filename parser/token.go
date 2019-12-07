package parser

// Token represents a terminal AST node.
type Token struct {
	Rune rune
	Val  string
	Pos  Position
}
