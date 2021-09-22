package parse

import (
	"reflect"
	"runtime"

	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseOpreration_RightLeft(t *testing.T) {
	op, err := parseOperation("R")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.Right), getOperationName(op))

	op, err = parseOperation("L")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.Left), getOperationName(op))
}

func Test_ParseOpreration_Invalid(t *testing.T) {
	_, err := parseOperation("NotValidOperation")
	require.Error(t, err)
}

func getOperationName(operation machine.Operation) string {
	return runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
}
