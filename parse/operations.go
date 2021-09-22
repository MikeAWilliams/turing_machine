package parse

import (
	"errors"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func parseOperation(op string) (machine.Operation, error) {
	if "R" == op {
		return machine.Right, nil
	}
	if "L" == op {
		return machine.Left, nil
	}
	return nil, errors.New("Not implmemented")
}
