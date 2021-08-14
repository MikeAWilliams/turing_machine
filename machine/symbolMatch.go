package machine

type SymbolMatch func(rune) bool

func SimpleSymbolMatch(toMatch rune) func(rune) bool {
	return func(symbol rune) bool {
		return toMatch == symbol
	}
}

func AnyNonBlankSymbolMatch(r rune) bool {
	return ' ' != r
}

func SetSymbolMatch(toMatch []rune) func(rune) bool {
	return func(symbol rune) bool {
		for _, r := range toMatch {
			if r == symbol {
				return true
			}
		}
		return false
	}
}
