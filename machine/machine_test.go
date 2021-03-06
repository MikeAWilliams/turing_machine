package machine_test

import (
	"errors"
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OneThirdMachineSingleOp(t *testing.T) {
	none := machine.SimpleSymbolMatch(' ')
	zero := machine.SimpleSymbolMatch('0')
	one := machine.SimpleSymbolMatch('1')
	m0 := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print('0')}), "b1")),
			machine.NewConfig("b1", machine.NewRow(zero, machine.NewOperations([]machine.Operation{machine.Right}), "c")),
			machine.NewConfig("c", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Right}), "e")),
			machine.NewConfig("e", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print('1')}), "e1")),
			machine.NewConfig("e1", machine.NewRow(one, machine.NewOperations([]machine.Operation{machine.Right}), "f")),
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
	s := m0.StateReport()
	require.Equal(t, "b", s.CurrentConfiguration)
	require.Equal(t, -1, s.OperationRow)
	require.Equal(t, 0, s.OperationColumn)

	m1, err := m0.Operate()
	require.NoError(t, err)
	require.Equal(t, "0", m1.TapeAsString())
	s = m1.StateReport()
	require.Equal(t, "b", s.CurrentConfiguration)
	require.Equal(t, 0, s.OperationRow)
	require.Equal(t, 1, s.OperationColumn)

	m2, err := m1.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 ", m2.TapeAsString())
	s = m2.StateReport()
	require.Equal(t, "c", s.CurrentConfiguration)
	require.Equal(t, -1, s.OperationRow)
	require.Equal(t, 0, s.OperationColumn)

	m3, err := m2.Operate()
	require.NoError(t, err)
	require.Equal(t, "0  ", m3.TapeAsString())

	m4, err := m3.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1", m4.TapeAsString())

	m5, err := m4.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 ", m5.TapeAsString())

	m6, err := m5.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1  ", m6.TapeAsString())

	m7, err := m6.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0", m7.TapeAsString())

	m8, err := m7.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 ", m8.TapeAsString())

	m9, err := m8.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0  ", m9.TapeAsString())

	m10, err := m9.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 1", m10.TapeAsString())
}

func Test_OneThirdMachineMultiRow(t *testing.T) {
	m0 := machine.TuringMachine2()

	m1, err := m0.Operate()
	require.NoError(t, err)
	require.Equal(t, "0", m1.TapeAsString())

	m2, err := m1.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 ", m2.TapeAsString())

	m3, err := m2.Operate()
	require.NoError(t, err)
	require.Equal(t, "0  ", m3.TapeAsString())

	m4, err := m3.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1", m4.TapeAsString())

	m5, err := m4.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 ", m5.TapeAsString())

	m6, err := m5.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1  ", m6.TapeAsString())

	m7, err := m6.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0", m7.TapeAsString())

	m8, err := m7.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 ", m8.TapeAsString())

	m9, err := m8.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0  ", m9.TapeAsString())

	m10, err := m9.Operate()
	require.NoError(t, err)
	require.Equal(t, "0 1 0 1", m10.TapeAsString())
}

func Test_ErrorOnInvalidConfig(t *testing.T) {
	none := machine.SimpleSymbolMatch(' ')
	m0 := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print('0')}), "b1")),
		})

	m1, err := m0.Operate()
	require.NoError(t, err)

	_, err = m1.Operate()
	require.Equal(t, errors.New("current configuration (b1) does not exist in table"), err)
}

func Test_ErrorUnableToFindRowMatchingSquare(t *testing.T) {
	none := machine.SimpleSymbolMatch(' ')
	m0 := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow(none, machine.NewOperations([]machine.Operation{machine.Print(machine.Schwa)}), "b")),
		})

	m1, err := m0.Operate()
	require.NoError(t, err)

	_, err = m1.Operate()
	require.Equal(t, fmt.Errorf("row not found matching current square (%v)", string(machine.Schwa)), err)
}
