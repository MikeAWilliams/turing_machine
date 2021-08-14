package machine_test

import (
	"github.com/MikeAWilliams/turing_machine/machine"

	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OneThirdMachine(t *testing.T) {
	machine := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Print('0'), machine.Right}), "c")),
			machine.NewConfig("c", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Right}), "e")),
			machine.NewConfig("e", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Print('1'), machine.Right}), "f")),
			machine.NewConfig("f", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Right}), "b")),
		})

	_, err := machine.Operate()

	require.NoError(t, err)
}
