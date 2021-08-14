package machine

import "errors"

type Operation func(Tape) Tape

type Operations struct {
	operations       []Operation
	currentOperation int
}

func (o Operations) copy() Operations {
	newOperations := make([]Operation, len(o.operations))
	copy(newOperations, o.operations)
	return Operations{operations: newOperations, currentOperation: o.currentOperation}
}

func (o Operations) Operate(t Tape) (Operations, Tape, error) {
	if len(o.operations) <= o.currentOperation {
		return Operations{}, Tape{}, errors.New("operation index out of bounds")
	}
	resultTape := o.operations[o.currentOperation](t)
	resultOperations := o.copy()
	resultOperations.currentOperation++
	return resultOperations, resultTape, nil
}

func (o Operations) IsDone() bool {
	return len(o.operations) <= o.currentOperation
}

func NewOperations(op []Operation) Operations {
	return Operations{operations: op}
}

func Right(t Tape) Tape {
	return t.Right()
}

func Left(t Tape) Tape {
	return t.Left()
}

func NoOp(t Tape) Tape {
	return t.NoOp()
}

func Print(symbol rune) func(Tape) Tape {
	return func(t Tape) Tape {
		return t.Print(symbol)
	}
}
