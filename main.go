package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
	"github.com/MikeAWilliams/turing_machine/parse"
	"github.com/fatih/color"
)

func executeOperations(machine machine.Machine, n int) {
	for i := 0; i < n; i++ {
		state := machine.StateReport()
		fmt.Printf("%v ", i)
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

		tmpMachine, err := machine.Operate()
		if nil != err {
			panic(err)
		}
		machine = tmpMachine
	}
	fmt.Printf("Done\nFinal Tape\n%v", machine.TapeAsString())

}

func main() {
	fmt.Println("Hello Dr. Turing")

	machineParse := parse.TuringMachine3()
	executeOperations(machineParse, 1000)
}
