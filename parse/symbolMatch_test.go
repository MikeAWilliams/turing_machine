package parse

import (
	"testing"

	"github.com/MikeAWilliams/turing_machine/machine"
	"github.com/stretchr/testify/require"
)

func Test_SymbolMatch_Simple(t *testing.T) {
	matcher, err := parseSymbolMatch("1")
	require.NoError(t, err)
	requireMatcherResult(t, '1', ' ', matcher)
}

func Test_SymbolMatch_None(t *testing.T) {
	matcher, err := parseSymbolMatch(" ")
	require.NoError(t, err)
	requireMatcherResult(t, ' ', '1', matcher)

	matcher, err = parseSymbolMatch("")
	require.NoError(t, err)
	requireMatcherResult(t, ' ', '1', matcher)

	matcher, err = parseSymbolMatch("None")
	require.NoError(t, err)
	requireMatcherResult(t, ' ', '1', matcher)

	matcher, err = parseSymbolMatch("NONE")
	require.NoError(t, err)
	requireMatcherResult(t, ' ', '1', matcher)
}

func Test_SymbolMatch_Any(t *testing.T) {
	matcher, err := parseSymbolMatch("Any")
	require.NoError(t, err)
	requireMatcherResult(t, '1', ' ', matcher)
	requireMatcherResult(t, 't', ' ', matcher)

	matcher, err = parseSymbolMatch("ANY")
	require.NoError(t, err)
	requireMatcherResult(t, '1', ' ', matcher)
	requireMatcherResult(t, 't', ' ', matcher)
}

func Test_SymbolMatch_AnyOr(t *testing.T) {
	matcher, err := parseSymbolMatch("Any (0 or 1)")
	require.NoError(t, err)
	requireMatcherResult(t, '0', ' ', matcher)
	requireMatcherResult(t, '1', ' ', matcher)

	matcher, err = parseSymbolMatch("Any (a or b)")
	require.NoError(t, err)
	requireMatcherResult(t, 'b', ' ', matcher)
	requireMatcherResult(t, 'a', ' ', matcher)
}

func Test_SymbolMatch_AnySet(t *testing.T) {
	matcher, err := parseSymbolMatch("Any (1,2,3)")
	require.NoError(t, err)
	requireMatcherResult(t, '1', ' ', matcher)
	requireMatcherResult(t, '2', ' ', matcher)
	requireMatcherResult(t, '3', ' ', matcher)
}

func Test_SymbolMatch_Error(t *testing.T) {
	_, err := parseSymbolMatch("INVALID")
	require.Error(t, err)
}

func requireMatcherResult(t *testing.T, expected rune, invalid rune, matcher machine.SymbolMatch) {
	require.True(t, matcher(expected))
	require.False(t, matcher(invalid))
}
