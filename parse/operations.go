package parse

import (
	"errors"
	"strings"

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
	case "E":
		return machine.Print(' '), nil
	}
	opRunes := []rune(op)
	if 2 == len(opRunes) && 'P' == opRunes[0] {
		return machine.Print(opRunes[1]), nil
	}
	return nil, errors.New("Not implmemented")
}

func parseOperations(operations string) ([]machine.Operation, error) {
	opStrings := strings.Split(operations, ",")
	result := make([]machine.Operation, len(opStrings))
	for index, op := range opStrings {
		thisOp, err := parseOperation(op)
		if nil != err {
			return []machine.Operation{}, err
		}
		result[index] = thisOp
	}
	return result, nil
}
