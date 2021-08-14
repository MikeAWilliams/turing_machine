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

func Test_Left(t *testing.T) {
	testObject := machine.NewTape()
	testObject = testObject.Print(machine.Schwa)

	newTape := testObject.Left()

	require.Equal(t, machine.Schwa, testObject.GetSymbol())
	require.Equal(t, ' ', newTape.GetSymbol())
}

func Test_Practical(t *testing.T) {
	tape := machine.NewTape()
	tape = tape.Print(machine.Schwa)
	tape = tape.Left()
	tape = tape.Print('z')
	tape = tape.Left()
	tape = tape.Print('y')
	tape = tape.Left()
	tape = tape.Print('x')

	require.Equal(t, 'x', tape.GetSymbol())
	tape = tape.Right()
	require.Equal(t, 'y', tape.GetSymbol())
	tape = tape.Right()
	require.Equal(t, 'z', tape.GetSymbol())
	tape = tape.Right()
	require.Equal(t, machine.Schwa, tape.GetSymbol())

	tape2 := tape.Print('0')
	tape2 = tape2.Left()
	tape2 = tape2.Print('1')
	tape2 = tape2.Left()
	tape2 = tape2.Print('2')
	tape2 = tape2.Left()
	tape2 = tape2.Print('3')

	require.Equal(t, '3', tape2.GetSymbol())
	tape2 = tape2.Right()
	require.Equal(t, '2', tape2.GetSymbol())
	tape2 = tape2.Right()
	require.Equal(t, '1', tape2.GetSymbol())
	tape2 = tape2.Right()
	require.Equal(t, '0', tape2.GetSymbol())

	require.Equal(t, machine.Schwa, tape.GetSymbol())
	tape = tape.Left()
	require.Equal(t, 'z', tape.GetSymbol())
	tape = tape.Left()
	require.Equal(t, 'y', tape.GetSymbol())
	tape = tape.Left()
	require.Equal(t, 'x', tape.GetSymbol())
}
