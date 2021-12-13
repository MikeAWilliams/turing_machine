package parse

import "github.com/MikeAWilliams/turing_machine/machine"

// The annotated Turing page 81
func TuringMachine1() machine.Machine {
	result, err := Machine(`
b None P0,R c
c None R    e
e None P1,R l
l None R    b`, " ")

	if nil != err {
		return machine.Machine{}
	}
	return result
}

// The annotated Turing page 84
func TuringMachine2() machine.Machine {
	result, err := Machine(`
b None   P0   b
  0    R,P1 b   
  1    R,R,P0 b`, " ")

	if nil != err {
		return machine.Machine{}
	}
	return result
}

// The annotated Turing page 87
func TuringMachine3() machine.Machine {
	result, err := Machine(`
b None         Pə,R,Pə,R,P0,R,R,P0,L,L o
o 1            R,Px,L,L,L              o
  0            N                       q
q Any (0 or 1) R,R                     q
  None         R1,L                    p
p x            E,R                     q
  ə            R                       f
  None         L,L                     p
f Any          R,R                     f
  None         P0,L,L                  o`, " ")

	if nil != err {
		return machine.Machine{}
	}
	return result
}
