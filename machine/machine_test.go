package machine_test

import (
    "github.com/MikeAWilliams/turing_machine/machine"

    "testing"
    "github.com/stretchr/testify/require"
)

func Test_NewTape(t *testing.T){
    testObject := machine.NewTape()
    require.Equal(t, ' ', testObject.GetSymbol())
}