package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func main() {
	fmt.Println("Hello Mr. Turing")

	machine := machine.TuringMachine1()

	for i := 0; i < 100; i++ {
		fmt.Printf("%v -%v-\n", i, machine.TapeAsString())
		tmpMachine, err := machine.Operate()
		if nil != err {
			panic(err)
		}
		machine = tmpMachine
	}

}
