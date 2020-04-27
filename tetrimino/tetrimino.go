package tetrimino

type ActionType int

const (
	Rotate ActionType = iota + 1
	Left
	Right
	Down
)

type Position struct {
	x int
	y int
}
type Posture int

const (
	DEG0 Posture = iota
	DEG90
	DEG180
	DEG270
)

type TetriminoType int

const (
	Stick TetriminoType = iota + 1
	LString
	ReverseLString
	Square
	TString
	SString
	ReverseSString
)

type Tetrimino struct {
	Pos           Position
	Rot           Posture
	tetriminoType TetriminoType
}

func NewTetrimino(tetriminoType TetriminoType) Tetrimino {
	var returnTetrimino Tetrimino
	returnTetrimino.Pos = Position{5, 0}
	returnTetrimino.Rot = DEG0
	returnTetrimino.tetriminoType = tetriminoType
	return returnTetrimino
}

func (tetrimino *Tetrimino) ActionToTetrimino(action ActionType) {
	switch {
	case action == Rotate:
		tetrimino.Rot = (tetrimino.Rot + 1) % 4
	case action == Left:
		tetrimino.Pos.x--
	case action == Right:
		tetrimino.Pos.x++
	case action == Down:
		tetrimino.Pos.y++
	default:
		panic("----Not Defined Action---")
	}
}
