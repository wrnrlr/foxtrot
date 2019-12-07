package parser

const (
	ExpressionPreInc ExpressionCase = iota // ++
	ExpressionPreDec                       // --
	ExpressionParenExpr
	ExpressionUnaryPlus  // +
	ExpressionUnaryMinus // -
	ExpressionNe         // !=
	ExpressionLAnd       // &&
	ExpressionMulAssign  // *=
	ExpressionPostInc    // ++
	ExpressionAddAssign  // +=
	ExpressionPostDec    // --
	ExpressionSubAssign  // -=
	ExpressionEq         // ==
	ExpressionGe         // <=
	ExpressionRsh
	ExpressionLOr
	ExpressionFactorial
	ExpressionMul // *
	ExpressionAdd // +
	ExpressionSub // -
	ExpressionDiv // /
	ExpressionCompound
	ExpressionLt     // <
	ExpressionAssign // =
	ExpressionGt     // >
	ExpressionPatternTest
	ExpressionOr
	ExpressionInfoShort
	ExpressionInfo
	ExpressionFloat
	ExpressionIdent
	ExpressionMessageName
	ExpressionMessageName2
	ExpressionInteger
	ExpressionPattern
	ExpressionSlot
	ExpressionString
)
