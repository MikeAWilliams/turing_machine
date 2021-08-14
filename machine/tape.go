package machine

type Tape struct {
	squares     []rune
	squareIndex int
}

func (t Tape) GetSymbol() rune {
	return t.squares[t.squareIndex]
}

func NewTape() Tape {
	return Tape{squares: []rune{' '}}
}
