package tetrimino

import (
	m "github.com/tetris-CLI/mino"
)

//ShapeType Tetriminoの形状の型
type ShapeType int

const (
	//IShape I型のTetrimino
	IShape ShapeType = iota + 1
	//LShape L型のTetrimino
	LShape
	//JShape J型のTetrimino
	JShape
	//OShape O型のTetrimino
	OShape
	//TShape T型のTetrimino
	TShape
	//SShape S型のTetrimino
	SShape
	//ZShape Z型のTetrimino
	ZShape
)

//Tetrimino Tetrimino構造体
type Tetrimino struct {
	//Tetriminoの形状
	Shape ShapeType
	//Tetriminoを構成するMinoの集合
	Minos [4]m.Mino
}

//NewTetrimino Tetrimino構造体を初期化して返す
func NewTetrimino(shape ShapeType) Tetrimino {
	switch shape {
	case IShape:
		return Tetrimino{
			Shape: IShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 5, Y: 0},
				{X: 6, Y: 0},
			},
		}
	default:
		panic("%s is undefined type of tetrimino")
	}
}

//Update tetriminoの情報から不整合を検出し，位置を再計算して更新する
func (tetrimino *Tetrimino) Update() {
	//TODO: 本当はIミノ以外にもある
	switch {
	case tetrimino.Shape == IShape:
		//TODO: 本当はRotによって違う
		if tetrimino.Minos[3].X > 9 {
			tetrimino.Minos[0].X = 6
			tetrimino.Minos[1].X = 7
			tetrimino.Minos[2].X = 8
			tetrimino.Minos[3].X = 9
		} else if tetrimino.Minos[0].X < 0 {
			tetrimino.Minos[0].X = 0
			tetrimino.Minos[1].X = 1
			tetrimino.Minos[2].X = 2
			tetrimino.Minos[3].X = 3
		}
		if tetrimino.Minos[0].Y >= 19 {
			tetrimino.Minos[0].Y = 19
			tetrimino.Minos[1].Y = 19
			tetrimino.Minos[2].Y = 19
			tetrimino.Minos[3].Y = 19
		}
	default:
		panic("%s is undefined type of tetrimino")
	}
}
