package machine

type Tape struct {
	squares     []rune
	squareIndex int
}

func (t Tape) copy() Tape {
	newSquares := make([]rune, len(t.squares))
	copy(newSquares, t.squares)
	return Tape{squares: newSquares, squareIndex: t.squareIndex}
}

func (t Tape) GetSymbol() rune {
	return t.squares[t.squareIndex]
}

func (t Tape) Print(symbol rune) Tape {
	result := t.copy()
	result.squares[result.squareIndex] = symbol
	return result
}

func NewTape() Tape {
	return Tape{squares: []rune{' '}}
}
