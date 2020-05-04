package tetrimino

import a "github.com/tetris-CLI/action"

//Position Stage上の座標
type Position struct {
	//水平方向の座標値
	X int
	//重力方向の座標値
	Y int
}

//Posture Tetriminoの初期状態に対しての回転角
type Posture int

const (
	//Deg0 Tetriminoが初期状態と同じ姿勢を持つ回転角
	Deg0 Posture = iota
	//Deg90 Tetriminoが初期状態から時計回りに90度回転した姿勢を持つ回転角
	Deg90
	//Deg180 Tetriminoが初期状態から時計回りに180度回転した姿勢を持つ回転角
	Deg180
	//Deg270 Tetriminoが初期状態から時計回りに270度回転した姿勢を持つ回転角
	Deg270
)

//BlockPositions Tetriminoの持つ4つのミノのPositionのArray型
type BlockPositions [4]Position

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
	//Tetriminoの代表点
	Pos Position
	//Tetriminoの標準角からの傾き
	Rot Posture
	//Tetriminoのそれぞれのブロックの位置
	BlockPoss BlockPositions
	//Tetriminoの形状
	Shape ShapeType
	//Tetriminoが非アクティブになるべきかどうか
	IsTerminate bool
}

//NewTetrimino Tetrimino構造体を初期化して返す
func NewTetrimino(shape ShapeType) Tetrimino {
	var returnTetrimino Tetrimino
	returnTetrimino.Pos = Position{4, 0}
	returnTetrimino.Rot = Deg0
	returnTetrimino.Shape = shape
	returnTetrimino.IsTerminate = false
	returnTetrimino.Update()
	return returnTetrimino
}

//Update tetriminoの情報から不整合を検出し，位置を再計算して更新する
func (tetrimino *Tetrimino) Update() {
	//TODO: 本当はIミノ以外にもある
	switch {
	case tetrimino.Shape == IShape:
		//TODO: 本当はRotによって違う
		if tetrimino.Pos.X >= 8 {
			tetrimino.Pos.X = 7
		} else if tetrimino.Pos.X <= 0 {
			tetrimino.Pos.X = 1
		} else {
		}
		if tetrimino.Pos.Y >= 19 {
			tetrimino.Pos.Y = 19
		}
		tetrimino.BlockPoss[0] = Position{tetrimino.Pos.X - 1, tetrimino.Pos.Y}
		tetrimino.BlockPoss[1] = Position{tetrimino.Pos.X, tetrimino.Pos.Y}
		tetrimino.BlockPoss[2] = Position{tetrimino.Pos.X + 1, tetrimino.Pos.Y}
		tetrimino.BlockPoss[3] = Position{tetrimino.Pos.X + 2, tetrimino.Pos.Y}
	default:
		panic("%s is undefined type of tetrimino")
	}
}

//ApplyAction actionに応じてtetriminoの位置や姿勢を更新する
func (tetrimino *Tetrimino) ApplyAction(action a.ActionType) {
	switch {
	case action == a.RotateLeftAction:
		tetrimino.Rot = (tetrimino.Rot + 4 - 1) % 4
		tetrimino.Update()
	case action == a.RotateRightAction:
		tetrimino.Rot = (tetrimino.Rot + 1) % 4
		tetrimino.Update()
	case action == a.MoveLeftAction:
		tetrimino.Pos.X--
		tetrimino.Update()
	case action == a.MoveRightAction:
		tetrimino.Pos.X++
		tetrimino.Update()
	case action == a.SoftDropAction:
		tetrimino.Pos.Y++
		tetrimino.Update()
	case action == a.HardDropAction:
		// TODO: implement hard drop behavior
		// tetrimino.Update()
	default:
		panic("----Not Defined Action---")
	}
}
