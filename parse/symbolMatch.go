package parse

import (
	"errors"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func parseSymbolMatch(text string) (machine.SymbolMatch, error) {
	runes := []rune(text)
	if 1 == len(runes) {
		return machine.SimpleSymbolMatch(runes[0]), nil
	}
	return nil, errors.New("not implmented")
}
