package parse_test

import (
	"testing"

	"github.com/MikeAWilliams/turing_machine/parse"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Simplest(t *testing.T) {
	_, err := parse.Machine("b\tAny\tP1\tb", "\t")

	require.NoError(t, err)
}
