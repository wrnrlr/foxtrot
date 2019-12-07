package parser

// Precedence maps token numbers to token precedence.
var Precedence map[int]int

func init() { Precedence = yyPrec }
