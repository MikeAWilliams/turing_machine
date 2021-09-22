package parse

import (
	"errors"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func parseOperation(op string) (machine.Operation, error) {
	switch op {
	case "":
		return machine.NoOp, nil
	case "R":
		return machine.Right, nil
	case "L":
		return machine.Left, nil
	}
	return nil, errors.New("Not implmemented")
}
