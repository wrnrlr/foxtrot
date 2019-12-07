package parser

import (
	"github.com/alecthomas/participle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserAssignment(t *testing.T) {
	expr := exampleAST("a = 1")
	assert.Nil(t, expr)
}

func TestParserLazySet(t *testing.T) {
	assert.Nil(t, exampleAST("a := 1"))
}

func TestParserSubtraction(t *testing.T) {
	assert.Nil(t, exampleAST("1.97 - 1.98"))
}

func TestParserAddition(t *testing.T) {
	assert.Nil(t, exampleAST("1.97 + 1.98"))
}

func exampleAST(s string) *Expression {
	expr := &Expression{}
	parser, _ := participle.Build(&Expression{})
	parser.ParseString(s, expr)
	return expr
}
