package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func main() {
	fmt.Println("Hello Mr. Turing")

	tape := machine.NewTape()
	fmt.Println(tape.GetSymbol())

}
