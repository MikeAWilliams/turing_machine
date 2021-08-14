package machine_test

import (
	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OneThirdMachineSingleOp(t *testing.T) {
	none := machine.SimpleSymbolMatch(' ')
	m0 := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print('0')}), "b1")),
			machine.NewConfig("b1", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Right}), "c")),
			machine.NewConfig("c", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Right}), "e")),
			machine.NewConfig("e", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print('1')}), "e1")),
			machine.NewConfig("e1", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Right}), "f")),
			machine.NewConfig("f", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Right}), "b")),
		})

	m1, err := m0.Operate()
	require.NoError(t, err)
	require.Equal(t, " ", m0.TapeAsString())
	require.Equal(t, "0", m1.TapeAsString())

	m2, err := m1.Operate()
	require.NoError(t, err)
	require.Equal(t, "0", m1.TapeAsString())
	require.Equal(t, "0 ", m2.TapeAsString())

	m3, err := m2.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 ", m2.TapeAsString())
	require.Equal(t, "0  ", m3.TapeAsString())

	m4, err := m3.Operate()
	require.NoError(t, err)
	require.Equal(t, "0  ", m3.TapeAsString())
	require.Equal(t, "0 1", m4.TapeAsString())

	m5, err := m4.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1", m4.TapeAsString())
	require.Equal(t, "0 1 ", m5.TapeAsString())

	m6, err := m5.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 ", m5.TapeAsString())
	require.Equal(t, "0 1  ", m6.TapeAsString())

	m7, err := m6.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1  ", m6.TapeAsString())
	require.Equal(t, "0 1 0", m7.TapeAsString())

	m8, err := m7.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0", m7.TapeAsString())
	require.Equal(t, "0 1 0 ", m8.TapeAsString())

	m9, err := m8.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 ", m8.TapeAsString())
	require.Equal(t, "0 1 0  ", m9.TapeAsString())

	m10, err := m9.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0  ", m9.TapeAsString())
	require.Equal(t, "0 1 0 1", m10.TapeAsString())
}

func Test_OneThirdMachineMultiOp(t *testing.T) {
	m0 := machine.TuringMachine1()

	m1, err := m0.Operate()
	require.NoError(t, err)
	require.Equal(t, " ", m0.TapeAsString())
	require.Equal(t, "0", m1.TapeAsString())

	m2, err := m1.Operate()
	require.NoError(t, err)
	require.Equal(t, "0", m1.TapeAsString())
	require.Equal(t, "0 ", m2.TapeAsString())

	m3, err := m2.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 ", m2.TapeAsString())
	require.Equal(t, "0  ", m3.TapeAsString())

	m4, err := m3.Operate()
	require.NoError(t, err)
	require.Equal(t, "0  ", m3.TapeAsString())
	require.Equal(t, "0 1", m4.TapeAsString())

	m5, err := m4.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1", m4.TapeAsString())
	require.Equal(t, "0 1 ", m5.TapeAsString())

	m6, err := m5.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 ", m5.TapeAsString())
	require.Equal(t, "0 1  ", m6.TapeAsString())

	m7, err := m6.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1  ", m6.TapeAsString())
	require.Equal(t, "0 1 0", m7.TapeAsString())

	m8, err := m7.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0", m7.TapeAsString())
	require.Equal(t, "0 1 0 ", m8.TapeAsString())

	m9, err := m8.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 ", m8.TapeAsString())
	require.Equal(t, "0 1 0  ", m9.TapeAsString())

	m10, err := m9.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0  ", m9.TapeAsString())
	require.Equal(t, "0 1 0 1", m10.TapeAsString())
}
