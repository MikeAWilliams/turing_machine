package machine_test

import (
	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SimpleSymbolMatch(t *testing.T) {
	matcher := machine.SimpleSymbolMatch(' ')
	require.True(t, matcher(' '))
	require.False(t, matcher('a'))

	matcher = machine.SimpleSymbolMatch(machine.Schwa)
	require.True(t, matcher(machine.Schwa))
	require.False(t, matcher('a'))
}
