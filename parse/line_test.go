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

func Test_Line_RunningConfigName(t *testing.T) {
	config, runningConfiguration, err := parseLine("\t1\tR,Px\tq", "\t", "o")

	require.NoError(t, err)
	require.Equal(t, "o", runningConfiguration)
	require.Equal(t, "o", config.GetConfigurationName())
	require.Equal(t, 2, config.GetOperationsCount())
	requireMatcherResult(t, '1', ' ', config.GetSymbolMatch())
	require.Equal(t, "q", config.GetFinalConfigurationName())
}
func Test_Line_RunningConfigName_whiteSpace(t *testing.T) {
	config, runningConfiguration, err := parseLine("  \t1\tR,Px\tq", "\t", "o")

	require.NoError(t, err)
	require.Equal(t, "o", runningConfiguration)
	require.Equal(t, "o", config.GetConfigurationName())
	require.Equal(t, 2, config.GetOperationsCount())
	requireMatcherResult(t, '1', ' ', config.GetSymbolMatch())
	require.Equal(t, "q", config.GetFinalConfigurationName())
}

func Test_Line_Error_IncorrectNumber(t *testing.T) {
	_, _, err := parseLine("o\t1\tR,Px\tq\td", "\t", "")

	require.Error(t, err)
}

func Test_Line_Error_EmptyConfigAndRunning(t *testing.T) {
	_, _, err := parseLine("\t1\tR,Px\tq", "\t", "")

	require.Error(t, err)
}

func Test_Line_Error_InvalidSymbolMatcher(t *testing.T) {
	_, _, err := parseLine("o\tINVALID\tR,Px\tq", "\t", "")

	require.Error(t, err)
}

func Test_Line_Error_InvalidOperations(t *testing.T) {
	_, _, err := parseLine("o\t1\tINVALID,Px\td", "\t", "")

	require.Error(t, err)
}

func Test_Line_Error_EmptyFinalConfig(t *testing.T) {
	_, _, err := parseLine("o\t1\tR,Px\t", "\t", "")

	require.Error(t, err)
}
