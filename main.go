package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
	"github.com/fatih/color"
)

func executeOperations(machine machine.Machine, n int) {
	for i := 0; i < n; i++ {
		state := machine.StateReport()
		fmt.Printf("%v -%v", i, string(state.Squares[:state.SquareIndex]))
		redText := color.New(color.FgRed, color.Bold)
		redText.Printf("|%v|", string(state.Squares[state.SquareIndex:state.SquareIndex+1]))
		fmt.Printf("%v- %v %v %v \n", string(state.Squares[state.SquareIndex+1:]), state.CurrentConfiguration, state.OperationRow, state.OperationColumn)

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

	machine := machine.TuringMachine3()
	executeOperations(machine, 1000)
}
