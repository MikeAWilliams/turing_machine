package main

import (
	"fmt"

	"github.com/MikeAWilliams/turing_machine/console"
	"github.com/MikeAWilliams/turing_machine/parse"
)

func main() {
	fmt.Println("Hello Dr. Turing")

	machine, err := parse.Machine(`
b~None        ~Pə,R,Pə,R,P0,R,R,P0,L,L~o
o~1           ~R,Px,L,L,L             ~o
 ~0           ~N                      ~q
q~Any (0 or 1)~R,R                    ~q
 ~None        ~P1,L                   ~p
p~x           ~E,R                    ~q
 ~ə           ~R                      ~f
 ~None        ~L,L                    ~p
f~Any         ~R,R                    ~f
 ~None        ~P0,L,L                 ~o`, "~")

	if nil != err {
		panic(err)
	}

	console.ExecuteOperations(machine, 1000)
}
