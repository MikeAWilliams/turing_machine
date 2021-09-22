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

	if strings.EqualFold("any", text) {
		return machine.AnyNonBlankSymbolMatch, nil
	}

	if strings.HasPrefix(strings.ToLower(text), "any (") {
		setStartIndex := strings.Index(text, "(")
		if -1 == setStartIndex {
			return nil, errors.New("invalid any statement, missing (")
		}
		setStopIndex := strings.Index(text, ")")
		if -1 == setStopIndex {
			return nil, errors.New("invalid any statement, missing )")
		}
		orIndex := strings.Index(text, "or")
		if -1 != orIndex {
			first := strings.TrimSpace(text[setStartIndex+1 : orIndex])
			firstRunes := []rune(first)
			if 1 != len(firstRunes) {
				return nil, errors.New("invalid any statement, left of or invalid")
			}
			second := strings.TrimSpace(text[orIndex+2 : setStopIndex])
			secondRunes := []rune(second)
			if 1 != len(secondRunes) {
				return nil, errors.New("invalid any statement, right of or invalid")
			}
			return machine.SetSymbolMatch(append(firstRunes, secondRunes...)), nil
		}
	}

	runes := []rune(text)
	if 1 == len(runes) {
		return machine.SimpleSymbolMatch(runes[0]), nil
	}
	return nil, errors.New("not implmented")
}
