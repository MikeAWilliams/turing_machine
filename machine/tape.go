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

func (t Tape) expandRight() Tape {
	newSquares := make([]rune, len(t.squares)+1)
	copy(newSquares, t.squares)
	newSquares[len(t.squares)] = ' '
	return Tape{squares: newSquares, squareIndex: t.squareIndex}
}

func (t Tape) expandLeft() Tape {
	newSquares := make([]rune, len(t.squares)+1)
	copy(newSquares[1:], t.squares)
	newSquares[0] = ' '
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

func (t Tape) Right() Tape {
	var result Tape
	if len(t.squares) == (t.squareIndex + 1) {
		result = t.expandRight()
	} else {
		result = t.copy()
	}
	result.squareIndex++
	return result
}

func (t Tape) Left() Tape {
	var result Tape
	if 0 == t.squareIndex {
		result = t.expandLeft()
	} else {
		result = t.copy()
		result.squareIndex--
	}
	return result
}

func (t Tape) NoOp() Tape {
	return t.copy()
}

func NewTape() Tape {
	return Tape{squares: []rune{' '}}
}
