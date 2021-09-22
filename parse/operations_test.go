package parse

import (
	"reflect"
	"runtime"

	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseOpreration(t *testing.T) {
	op, err := parseOperation("R")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.Right), getOperationName(op))
}

func getOperationName(operation machine.Operation) string {
	return runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
}
