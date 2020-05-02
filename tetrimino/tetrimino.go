package tetrimino

//ActionType テトリミノのプロパティを更新するためのAction型
type ActionType int

const (
	//RotateLeftAction テトリミノを左方向 (反時計回り) に回転するAction
	RotateLeftAction ActionType = iota + 1
	//RotateRightAction テトリミノを右方向 (時計回り) に回転するAction
	RotateRightAction
	//MoveLeftAction テトリミノを左方向に1ミノ分移動するAction
	MoveLeftAction
	//MoveRightAction テトリミノを右方向に1ミノ分移動するAction
	MoveRightAction
	//SoftDropAction テトリミノを下方向に1ミノ分移動するAction
	SoftDropAction
	//HardDropAction テトリミノを可能な限り下方向に移動し、固定するAction
	HardDropAction
)

//Position Xを水平方向，Y軸は重力方向を表すStage上の座標
type Position struct {
	X int
	Y int
}

//Posture Tetriminoの初期状態に対しての回転角
type Posture int

const (
	//Deg0 テトリミノが初期状態と同じ姿勢を持つ回転角
	Deg0 Posture = iota
	//Deg90 テトリミノが初期状態から時計回りに90度回転した姿勢を持つ回転角
	Deg90
	//Deg180 テトリミノが初期状態から時計回りに180度回転した姿勢を持つ回転角
	Deg180
	//Deg270 テトリミノが初期状態から時計回りに270度回転した姿勢を持つ回転角
	Deg270
)

//BlockPositions テトリミノの持つ4つのポジションのlist型
type BlockPositions [4]Position

//ShapeType テトリミノの７種類の形の型
type ShapeType int

const (
	//IShape I型のテトリミノ
	IShape ShapeType = iota + 1
	//LShape L型のテトリミノ
	LShape
	//JShape J型のテトリミノ
	JShape
	//OShape O型のテトリミノ
	OShape
	//TShape T型のテトリミノ
	TShape
	//SShape S型のテトリミノ
	SShape
	//ZShape Z型のテトリミノ
	ZShape
)

//Tetrimino Tetrimino構造体，その代表点と標準角からの傾き，それぞれのブロックの位置．Tetriminoが非アクティブになるべきかどうかなどを持っている
type Tetrimino struct {
	Pos         Position
	Rot         Posture
	BlockPoss   BlockPositions
	Shape       ShapeType
	IsTerminate bool
}

//NewTetrimino Tetrimino構造体を初期化して返す関数
func NewTetrimino(shape ShapeType) Tetrimino {
	var returnTetrimino Tetrimino
	returnTetrimino.Pos = Position{4, 0}
	returnTetrimino.Rot = Deg0
	returnTetrimino.Shape = shape
	returnTetrimino.IsTerminate = false
	returnTetrimino.Update()
	return returnTetrimino
}

//Update tetriminoの情報から，不整合を検出して．テトリミノのブロックの位置を計算して更新
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
func (tetrimino *Tetrimino) ApplyAction(action ActionType) {
	switch {
	case action == RotateLeftAction:
		tetrimino.Rot = (tetrimino.Rot + 4 - 1) % 4
		tetrimino.Update()
	case action == RotateRightAction:
		tetrimino.Rot = (tetrimino.Rot + 1) % 4
		tetrimino.Update()
	case action == MoveLeftAction:
		tetrimino.Pos.X--
		tetrimino.Update()
	case action == MoveRightAction:
		tetrimino.Pos.X++
		tetrimino.Update()
	case action == SoftDropAction:
		tetrimino.Pos.Y++
		tetrimino.Update()
	case action == HardDropAction:
		// TODO: implement hard drop behavior
		// tetrimino.Update()
	default:
		panic("----Not Defined Action---")
	}
}
