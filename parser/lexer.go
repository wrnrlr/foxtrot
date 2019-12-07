package parser

import "unicode"

const bnf = `
Comment = "(*"  "*)" .
Ident = (alpha | "_") { "_" | alpha | digit } .
Symbol = (alpha | "_") { "_" | alpha | digit } .
Complex = (alpha | "_") { "_" | alpha | digit } .
Expression 
Whitespace = " " | "\t" | "\n" | "\r" .

`

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isOperator(ch rune) bool {
	switch ch {
	case '!', '[', ']', '(', ')', '^', '+', '-', '*', '/', '=', '<', '>', ',', ';', '{', '}', '&':
		return true
	default:
		return false
	}
}
