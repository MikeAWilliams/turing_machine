package parse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Line_HappyPath(t *testing.T) {
	_, runningConfiguration, err := parseLine("o\t1\tR,Px\tq", "\t", "")

	require.NoError(t, err)
	require.Equal(t, "o", runningConfiguration)
}
