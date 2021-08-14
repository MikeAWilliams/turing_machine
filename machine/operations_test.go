package machine_test

import (
	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PrintZeroOne(t *testing.T) {
	operations := machine.NewOperations([]machine.Operation{machine.Print('0'), machine.Right, machine.Print('1')})
	tape := machine.NewTape()

	doATest := func() {
		newOperations, newTape, err := operations.Operate(tape)
		require.NoError(t, err)
		require.Equal(t, '0', newTape.GetSymbol())

		newOperations, newTape, err = newOperations.Operate(newTape)
		require.NoError(t, err)
		require.Equal(t, ' ', newTape.GetSymbol())

		newOperations, newTape, err = newOperations.Operate(newTape)
		require.NoError(t, err)
		require.Equal(t, '1', newTape.GetSymbol())
	}
	doATest()
	doATest()
}

func Test_OperationsIsDone(t *testing.T) {
	operations := machine.NewOperations([]machine.Operation{machine.Print('0')})
	tape := machine.NewTape()

	newOperations, tape, err := operations.Operate(tape)
	require.NoError(t, err)
	require.False(t, operations.IsDone())
	require.True(t, newOperations.IsDone())
}

func Test_OperateToManyTimes(t *testing.T) {
	operations := machine.NewOperations([]machine.Operation{machine.Print('0')})
	tape := machine.NewTape()

	operations, tape, err := operations.Operate(tape)
	require.NoError(t, err)
	operations, tape, err = operations.Operate(tape)
	require.Error(t, err)
}
