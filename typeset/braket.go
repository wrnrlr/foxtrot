package typeset

var (
	LeftRoundBraket  = &Label{Text: "(", MaxWidth: FitContent}
	RightRoundBraket = &Label{Text: ")", MaxWidth: FitContent}

	LeftSquareBraket  = &Label{Text: "[", MaxWidth: FitContent}
	RightSquareBraket = &Label{Text: "]", MaxWidth: FitContent}

	LeftCurlyBraket  = &Label{Text: "{", MaxWidth: FitContent}
	RightCurlyBraket = &Label{Text: "}", MaxWidth: FitContent}
)

var (
	PlusSymbol       = &Label{Text: "+", MaxWidth: FitContent}
	MinusSymbol      = &Label{Text: "-", MaxWidth: FitContent}
	MultiplySymbol   = &Label{Text: "*", MaxWidth: FitContent}
	FactorSymbol     = &Label{Text: "!", MaxWidth: FitContent}
	InterpunctSymbol = &Label{Text: "·", MaxWidth: FitContent}
	ModuloSymbol     = &Label{Text: "%", MaxWidth: FitContent}

	SqrtSymbol = &Label{Text: "√", MaxWidth: FitContent}
)
