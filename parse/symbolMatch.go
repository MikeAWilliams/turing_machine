package parse

import (
	"errors"
	"strings"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func parseSymbolMatch(text string) (machine.SymbolMatch, error) {
	trimmed := strings.TrimSpace(text)
	if 0 == len(trimmed) {
		return machine.SimpleSymbolMatch(' '), nil
	}
	if strings.EqualFold("none", text) {
		return machine.SimpleSymbolMatch(' '), nil
	}

	runes := []rune(text)
	if 1 == len(runes) {
		return machine.SimpleSymbolMatch(runes[0]), nil
	}
	return nil, errors.New("not implmented")
}
