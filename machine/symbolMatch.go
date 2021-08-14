package machine

import (
	"fmt"
	"unicode/utf8"
)

type SymbolMatch func(rune) bool

func SimpleSymbolMatch(toMatch rune) func(rune) bool {
	return func(symbol rune) bool {
		return toMatch == symbol
	}
}

func GetSymbolMatch(symbol string) (SymbolMatch, error) {
	if "None" == symbol {
		return SimpleSymbolMatch(' '), nil
	}
	if 1 == utf8.RuneCountInString(symbol) {
		r, _ := utf8.DecodeLastRuneInString(symbol)
		return SimpleSymbolMatch(r), nil
	}
	return nil, fmt.Errorf("unsupported symbol %v", symbol)
}
