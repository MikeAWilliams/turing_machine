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
	opRunes := []rune(op)
	if 2 == len(opRunes) && 'P' == opRunes[0] {
		return machine.Print(opRunes[1]), nil
	}
	return nil, errors.New("Not implmemented")
}
