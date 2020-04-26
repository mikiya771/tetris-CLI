package tetrimino_test

import "testing"

func TestNewTetrimino(t *testing.T) {
	tetriminoType := "aaa"
	tetrimino := Tetrimino.NewTetrimino(tetriminoType)
	testify.assert(t, Position{5, 5}, tetrimino.PositionOfTetrimino)
	testify.assert(t, Rotate{0}, tetrimino.RotateOfTetrimino)
}
func TestMoveTetrimino(t *testing.T) {
	//TODO: define TetriminoType as form of blocks
	tetriminoType := "aaa"
	tetrimino := Tetrimino.NewTetrimino(tetriminoType)
	//TODO: define KeyType of UserInput
	keytype := "up"
	tetrimino.MoveTetrimino(keytype)
	var pos Position
	pos = tetrimino.PositionOfTetrimino
	expectedPos := Position{5, 5}
	testify.assert(t, expectedPos, pos)
}
