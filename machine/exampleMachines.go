package machine

func TuringMachine1() Machine {
	none := SimpleSymbolMatch(' ')
	return NewMachine("b",
		[]ConfigOP{
			NewConfig("b", NewRow(none, NewOperations([]Operation{Print('0'), Right}), "c")),
			NewConfig("c", NewRow(none, NewOperations([]Operation{Right}), "e")),
			NewConfig("e", NewRow(none, NewOperations([]Operation{Print('1'), Right}), "f")),
			NewConfig("f", NewRow(none, NewOperations([]Operation{Right}), "b")),
		})
}

func TuringMachine2() Machine {
	none := SimpleSymbolMatch(' ')
	zero := SimpleSymbolMatch('0')
	one := SimpleSymbolMatch('1')
	return NewMachine("b",
		[]ConfigOP{
			NewConfig("b", NewRow(none, NewOperations([]Operation{Print('0')}), "b")),
			NewConfig("b", NewRow(zero, NewOperations([]Operation{Right, Right, Print('1')}), "b")),
			NewConfig("b", NewRow(one, NewOperations([]Operation{Right, Right, Print('0')}), "b")),
		})
}
