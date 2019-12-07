package parser

type Expression struct {
	Operator ExpressionCase

	CommaOpt *CommaOpt

	Expression  *Expression
	Expression2 *Expression
	Expression3 *Expression

	Tag  *Tag
	Tag2 *Tag

	Token  Value
	Token2 Value
	Token3 Value

	Pos Position
}
