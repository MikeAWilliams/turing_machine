package parse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Line_HappyPath(t *testing.T) {
	config, runningConfiguration, err := parseLine("o\t1\tR,Px\tq", "\t", "")

	require.NoError(t, err)
	require.Equal(t, "o", runningConfiguration)
	require.Equal(t, "o", config.GetConfigurationName())
	require.Equal(t, 2, config.GetOperationsCount())
	requireMatcherResult(t, '1', ' ', config.GetSymbolMatch())
	require.Equal(t, "q", config.GetFinalConfigurationName())
}
