package tetrimino

import (
	m "github.com/tetris-CLI/store/mino"
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

//Clone Tetriminoインスタンスを複製して返す
func (tetrimino *Tetrimino) Clone() Tetrimino {
	return Tetrimino{
		Shape: tetrimino.Shape,
		Minos: [4]m.Mino{
			tetrimino.Minos[0].Clone(),
			tetrimino.Minos[1].Clone(),
			tetrimino.Minos[2].Clone(),
			tetrimino.Minos[3].Clone(),
		},
	}
}

//MoveBy Tetriminoを指定された量だけ移動する
func (tetrimino *Tetrimino) MoveBy(dx, dy int) {
	for i := 0; i < len(tetrimino.Minos); i++ {
		tetrimino.Minos[i].X += dx
		tetrimino.Minos[i].Y += dy
	}
}
