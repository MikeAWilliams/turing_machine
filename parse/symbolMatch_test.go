package parse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SymbolMatch_Simple(t *testing.T) {
	matcher, err := parseSymbolMatch("1")
	require.NoError(t, err)
	require.True(t, matcher('1'))
	require.False(t, matcher(' '))
}
