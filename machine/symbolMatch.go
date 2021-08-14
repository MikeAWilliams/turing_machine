package machine

type SymbolMatch func(rune) bool

func SimpleSymbolMatch(toMatch rune) func(rune) bool {
	return func(symbol rune) bool {
		return toMatch == symbol
	}
}
