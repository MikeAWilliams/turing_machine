package machine_test

import (
	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewTape(t *testing.T) {
	testObject := machine.NewTape()
	require.Equal(t, ' ', testObject.GetSymbol())
}

func Test_Print(t *testing.T) {
	testObject := machine.NewTape()

	newTape := testObject.Print(machine.Schwa)

	require.Equal(t, ' ', testObject.GetSymbol())
	require.Equal(t, machine.Schwa, newTape.GetSymbol())
}

func Test_Right(t *testing.T) {
	testObject := machine.NewTape()
	testObject = testObject.Print(machine.Schwa)

	newTape := testObject.Right()

	require.Equal(t, machine.Schwa, testObject.GetSymbol())
	require.Equal(t, ' ', newTape.GetSymbol())
}
