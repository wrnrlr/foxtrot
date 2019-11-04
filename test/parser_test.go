package test

import (
  "bytes"
  "fmt"
  "go/token"
  "modernc.org/wl"
  "testing"
)

func TestParser(t *testing.T) {
  expr := parse(`Print["Hello"]`)
  fmt.Println(expr.String())
  fmt.Printf("Position: %d", expr.Pos())
}

func TestParser2(t *testing.T) {
  expr := parse("2+2")
  fmt.Println(expr.String())
  fmt.Printf("Position: %d", expr.Pos())
}

func parse(s string) *wl.Expression {
  buf := bytes.NewBufferString(s)
  in, err := wl.NewInput(buf, true)
  if err != nil {
    fmt.Println("Lexer failed")
  }
  expr, err := in.ParseExpression(token.NewFileSet().AddFile("nofile", -1, 1e6))
  if err != nil {
    fmt.Println("Parsing failed")
  }
  return expr
}

