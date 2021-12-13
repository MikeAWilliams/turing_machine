package parse

import (
	"reflect"
	"runtime"

	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseOpreration_RightLeftNoOp(t *testing.T) {
	op, err := parseOperation("R")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.Right), getOperationName(op))

	op, err = parseOperation("L")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.Left), getOperationName(op))

	op, err = parseOperation("")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.NoOp), getOperationName(op))

	op, err = parseOperation("N")
	require.NoError(t, err)
	require.Equal(t, getOperationName(machine.NoOp), getOperationName(op))
}

func Test_ParseOpreration_Print(t *testing.T) {
	op, err := parseOperation("Px")
	require.NoError(t, err)
	requireTapeRuneAfterOp(t, "x", op)

	op, err = parseOperation("P ")
	require.NoError(t, err)
	requireTapeRuneAfterOp(t, " ", op)

	op, err = parseOperation("E")
	require.NoError(t, err)
	requireTapeRuneAfterOp(t, " ", op)
}

func Test_ParseOpreration_Invalid(t *testing.T) {
	_, err := parseOperation("NotValidOperation")
	require.Error(t, err)
}

func Test_ParseOperations_General(t *testing.T) {
	operations, err := parseOperations("L,R,,Px,P ,      ")

	require.NoError(t, err)
	require.Equal(t, 6, len(operations))
	require.Equal(t, getOperationName(machine.Left), getOperationName(operations[0]))
	require.Equal(t, getOperationName(machine.Right), getOperationName(operations[1]))
	require.Equal(t, getOperationName(machine.NoOp), getOperationName(operations[2]))
	requireTapeRuneAfterOp(t, "x", operations[3])
	requireTapeRuneAfterOp(t, " ", operations[4])
	require.Equal(t, getOperationName(machine.NoOp), getOperationName(operations[5]))
}

func Test_ParseOperations_Invalid(t *testing.T) {
	_, err := parseOperations("NotValidOperations")
	require.Error(t, err)
}

func getOperationName(operation machine.Operation) string {
	return runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
}

func requireTapeRuneAfterOp(t *testing.T, expected string, op machine.Operation) {
	initialTape := machine.Print(',')(machine.NewTape())
	operatedTape := op(initialTape)
	require.Equal(t, expected, string(operatedTape.GetSymbol()))
}
