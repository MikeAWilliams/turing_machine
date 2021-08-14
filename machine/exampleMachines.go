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

func TuringMachine3() Machine {
	none := SimpleSymbolMatch(' ')
	zero := SimpleSymbolMatch('0')
	one := SimpleSymbolMatch('1')
	x := SimpleSymbolMatch('x')
	schwa := SimpleSymbolMatch(Schwa)
	any := AnyNonBlankSymbolMatch
	anyZeroOne := SetSymbolMatch([]rune{'0', '1'})
	return NewMachine("b",
		[]ConfigOP{
			NewConfig("b", NewRow(none, NewOperations([]Operation{Print(Schwa), Right, Print(Schwa), Right, Print('0'), Right, Right, Print('0'), Left, Left}), "o")),
			NewConfig("o", NewRow(one, NewOperations([]Operation{Right, Print('x'), Left, Left, Left}), "o")),
			NewConfig("o", NewRow(zero, NewOperations([]Operation{NoOp}), "q")),
			NewConfig("q", NewRow(anyZeroOne, NewOperations([]Operation{Right, Right}), "q")),
			NewConfig("q", NewRow(none, NewOperations([]Operation{Print('1'), Left}), "p")),
			NewConfig("p", NewRow(x, NewOperations([]Operation{Print(' '), Right}), "q")),
			NewConfig("p", NewRow(schwa, NewOperations([]Operation{Right}), "f")),
			NewConfig("p", NewRow(none, NewOperations([]Operation{Left, Left}), "p")),
			NewConfig("f", NewRow(any, NewOperations([]Operation{Right, Right}), "f")),
			NewConfig("f", NewRow(none, NewOperations([]Operation{Print('0'), Left, Left}), "o")),
		})
}
