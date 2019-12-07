package parser

import "github.com/cznic/strutil"

var (
	testFile *token.File // Testing hook

	hooks = strutil.PrettyPrintHooks{
		reflect.TypeOf(Token{}): func(f strutil.Formatter, v interface{}, prefix string, suffix string) {
			t := v.(Token)
			if t.Rune == 0 {
				return
			}

			f.Format(prefix)
			if testFile != nil {
				f.Format("%s: ", testFile.Position(t.Pos()))
			}
			f.Format("%s", yySymName(int(t.Rune)))
			if t.Val != "" {
				f.Format(", %q", t.Val)
			}
			f.Format(suffix)
		},
		reflect.TypeOf(ExpressionCase(0)): func(f strutil.Formatter, v interface{}, prefix string, suffix string) {
			t := v.(ExpressionCase)
			f.Format(prefix)
			f.Format("%s", t)
			f.Format(suffix)
		},
	}
)

func prettyString(v interface{}) string { return strutil.PrettyString(v, "", "", hooks) }
