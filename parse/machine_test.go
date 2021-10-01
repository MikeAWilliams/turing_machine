package parse_test

import (
	"testing"

	"github.com/MikeAWilliams/turing_machine/parse"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Valid_OneLine(t *testing.T) {
	result, err := parse.Machine("b\tAny\tP1\tb", "\t")

	require.NoError(t, err)
	require.Equal(t, 1, result.GetRowCount())
}

func Test_Parse_Valid_TwoLine(t *testing.T) {
	result, err := parse.Machine("b\t1\tR,P0\tb\n\t0\tR,P1\tb", "\t")

	require.NoError(t, err)
	require.Equal(t, 2, result.GetRowCount())
}

func Test_Parse_Valid_MultipleLinesRealTabs(t *testing.T) {
	result, err := parse.Machine(`
b   NONE   R,P0    b
b   1   R,P0    b
b   0   R,P1    b`, "   ")

	require.NoError(t, err)
	require.Equal(t, 3, result.GetRowCount())
}
