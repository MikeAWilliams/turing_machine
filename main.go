package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func main() {
	fmt.Println("Hello Mr. Turing")

	machine := machine.NewMachine("b",
		[]machine.ConfigOP{
			machine.NewConfig("b", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Print('0'), machine.Right}), "c")),
			machine.NewConfig("c", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Right}), "e")),
			machine.NewConfig("e", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Print('1'), machine.Right}), "f")),
			machine.NewConfig("f", machine.NewRow("None", machine.NewOperations([]machine.Operation{machine.Right}), "b")),
		})

	for i := 0; i < 100; i++ {
		fmt.Printf("%v -%v-\n", i, machine.TapeAsString())
		tmpMachine, err := machine.Operate()
		if nil != err {
			panic(err)
		}
		machine = tmpMachine
	}

}
