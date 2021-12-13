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

	if strings.EqualFold("none", trimmed) {
		return machine.SimpleSymbolMatch(' '), nil
	}

	if strings.EqualFold("any", trimmed) {
		return machine.AnyNonBlankSymbolMatch, nil
	}

	if strings.HasPrefix(strings.ToLower(trimmed), "any (") {
		setStartIndex := strings.Index(trimmed, "(")
		if -1 == setStartIndex {
			return nil, errors.New("invalid any statement, missing (")
		}
		setStopIndex := strings.Index(trimmed, ")")
		if -1 == setStopIndex {
			return nil, errors.New("invalid any statement, missing )")
		}
		orIndex := strings.Index(trimmed, "or")
		if -1 != orIndex {
			first := strings.TrimSpace(trimmed[setStartIndex+1 : orIndex])
			firstRunes := []rune(first)
			if 1 != len(firstRunes) {
				return nil, errors.New("invalid any statement, left of or invalid")
			}
			second := strings.TrimSpace(trimmed[orIndex+2 : setStopIndex])
			secondRunes := []rune(second)
			if 1 != len(secondRunes) {
				return nil, errors.New("invalid any statement, right of or invalid")
			}
			return machine.SetSymbolMatch(append(firstRunes, secondRunes...)), nil
		} else {
			betweenParens := trimmed[setStartIndex+1 : setStopIndex]
			split := strings.Split(betweenParens, ",")
			toMatch := make([]rune, len(split))
			for splitIndex, symbol := range split {
				symbolRunes := []rune(strings.TrimSpace(symbol))
				if 1 != len(symbolRunes) {
					return nil, errors.New("invalid any statement, more than one character between ,")
				}
				toMatch[splitIndex] = symbolRunes[0]
			}
			return machine.SetSymbolMatch(toMatch), nil
		}

	}

	runes := []rune(trimmed)
	if 1 == len(runes) {
		return machine.SimpleSymbolMatch(runes[0]), nil
	}
	return nil, errors.New("not implmented")
}
