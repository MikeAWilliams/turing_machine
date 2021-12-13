package console

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
	"github.com/fatih/color"

	"github.com/eiannone/keyboard"
)

func outputState(state machine.StateReport, row int) {
	fmt.Printf("%v ", row)
	for j := 0; j < state.SquareIndex; j++ {
		fmt.Printf("%v", string(state.Squares[j]))
	}
	redText := color.New(color.FgRed, color.Bold)
	currentSquare := state.Squares[state.SquareIndex]
	if ' ' == currentSquare {
		currentSquare = '_'
	}
	redText.Printf("%v", string(currentSquare))
	for j := state.SquareIndex + 1; j < len(state.Squares); j++ {
		fmt.Printf("%v", string(state.Squares[j]))
	}
	fmt.Printf(" %v %v %v \n", state.CurrentConfiguration, state.OperationRow, state.OperationColumn)
}

func ExecuteOperations(machine machine.Machine, n int) {
	for i := 0; i < n; i++ {
		state := machine.StateReport()
		outputState(state, i)

		tmpMachine, err := machine.Operate()
		if nil != err {
			panic(err)
		}
		machine = tmpMachine
	}
	fmt.Printf("Done\nFinal Tape\n%v", machine.TapeAsString())
}

func ExecuteInteractive(currentMachine machine.Machine) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press b to pop state. q to quit. Any other key advances")
	row := 0
	var machineStack []machine.Machine
	machineStack = append(machineStack, currentMachine)
	for {
		state := currentMachine.StateReport()
		outputState(state, row)

		key, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch key {
		case 'q':
			return
		case 'b':
			lastIndex := len(machineStack) - 1
			currentMachine = machineStack[lastIndex]
			machineStack = machineStack[:lastIndex]
			row--
		default:
			tmpMachine, err := currentMachine.Operate()
			if nil != err {
				panic(err)
			}
			machineStack = append(machineStack, currentMachine)
			currentMachine = tmpMachine
			row++
		}
	}
}
