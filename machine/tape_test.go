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
	require.Equal(t, ' ', testObject.GetSymbol())

	newTape := testObject.Print(machine.Schwa)
	require.Equal(t, ' ', testObject.GetSymbol())
	require.Equal(t, machine.Schwa, newTape.GetSymbol())
}
