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
	//Tetriminoの回転軸のMinoのindex
	CenterMinoIndex int
}

//NewTetrimino Tetrimino構造体を初期化して返す
func NewTetrimino(shape ShapeType) Tetrimino {
	switch shape {
	case IShape:
		return Tetrimino{
			Shape: IShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 1},
				{X: 4, Y: 1},
				{X: 5, Y: 1},
				{X: 6, Y: 1},
			},
			CenterMinoIndex: 2,
		}
	case LShape:
		return Tetrimino{
			Shape: LShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 1},
				{X: 4, Y: 1},
				{X: 5, Y: 1},
				{X: 5, Y: 0},
			},
			CenterMinoIndex: 1,
		}
	case JShape:
		return Tetrimino{
			Shape: JShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 0},
				{X: 3, Y: 1},
				{X: 4, Y: 1},
				{X: 5, Y: 1},
			},
			CenterMinoIndex: 2,
		}
	case OShape:
		return Tetrimino{
			Shape: OShape,
			Minos: [4]m.Mino{
				{X: 4, Y: 0},
				{X: 4, Y: 1},
				{X: 5, Y: 0},
				{X: 5, Y: 1},
			},
			CenterMinoIndex: 1,
		}
	case TShape:
		return Tetrimino{
			Shape: TShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 1},
				{X: 4, Y: 1},
				{X: 4, Y: 0},
				{X: 5, Y: 1},
			},
			CenterMinoIndex: 1,
		}
	case SShape:
		return Tetrimino{
			Shape: SShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 1},
				{X: 4, Y: 0},
				{X: 4, Y: 1},
				{X: 5, Y: 0},
			},
			CenterMinoIndex: 1,
		}
	case ZShape:
		return Tetrimino{
			Shape: ZShape,
			Minos: [4]m.Mino{
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 1},
				{X: 5, Y: 1},
			},
			CenterMinoIndex: 2,
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
		CenterMinoIndex: tetrimino.CenterMinoIndex,
	}
}

//MoveBy Tetriminoを指定された量だけ移動する
func (tetrimino *Tetrimino) MoveBy(dx, dy int) {
	for i := 0; i < len(tetrimino.Minos); i++ {
		tetrimino.Minos[i].X += dx
		tetrimino.Minos[i].Y += dy
	}
}

const sin90, sinMinus90 = 1, -1
const cos90, cosMinus90 = 0, 0

//RotateToLeft Tetriminoを90度左回転する
func (tetrimino *Tetrimino) RotateToLeft() {
	centerMino := tetrimino.Minos[tetrimino.CenterMinoIndex].Clone()

	for i := 0; i < len(tetrimino.Minos); i++ {
		dx := tetrimino.Minos[i].X - centerMino.X
		dy := tetrimino.Minos[i].Y - centerMino.Y

		tetrimino.Minos[i].X += dx*cos90 - dy*sin90
		tetrimino.Minos[i].Y += dx*sin90 + dy*cos90
	}
}

//RotateToRight Tetriminoを90度右回転する
func (tetrimino *Tetrimino) RotateToRight() {
	centerMino := tetrimino.Minos[tetrimino.CenterMinoIndex].Clone()

	for i := 0; i < len(tetrimino.Minos); i++ {
		dx := tetrimino.Minos[i].X - centerMino.X
		dy := tetrimino.Minos[i].Y - centerMino.Y

		tetrimino.Minos[i].X = centerMino.X + dx*cosMinus90 - dy*sinMinus90
		tetrimino.Minos[i].Y = centerMino.Y + dx*sinMinus90 + dy*cosMinus90
	}
}
