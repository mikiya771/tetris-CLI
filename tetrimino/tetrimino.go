package tetrimino

//ActionType Tetriminoに渡す，Actionの内容の型 これに応じた更新方法でTetriminoのプロパティを更新する
type ActionType int

const (
	Rotate ActionType = iota + 1
	Left
	Right
	Down
)

//Position Stage上の座標を意味する，Y軸は重力方向
type Position struct {
	X int
	Y int
}

//Posture Tetriminoの初期状態に対しての回転角を意味する
type Posture int

const (
	DEG0 Posture = iota
	DEG90
	DEG180
	DEG270
)

//Tetriminoのもつ四つのポジションのリスト
type BlockPositions [4]Position

//Tetriminoの７種類の形それぞれのものを意味する
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

//Tetrimino構造体，その代表点と標準角からの傾き，それぞれのブロックの位置．Tetriminoが非アクティブになるべきかどうかなどを持っている
type Tetrimino struct {
	Pos           Position
	Rot           Posture
	BlockPoss     BlockPositions
	tetriminoType TetriminoType
	IsTerminate   bool
}

//Tetrimino構造体を初期化して返す関数
func NewTetrimino(tetriminoType TetriminoType) Tetrimino {
	var returnTetrimino Tetrimino
	returnTetrimino.Pos = Position{4, 0}
	returnTetrimino.Rot = DEG0
	returnTetrimino.tetriminoType = tetriminoType
	returnTetrimino.IsTerminate = false
	returnTetrimino.Update()
	return returnTetrimino
}

//tetriminoの情報から，不整合を検出して．テトリミノのブロックの位置を計算して更新
func (tetrimino *Tetrimino) Update() {
	//TODO: 本当はStick意外にもある
	switch {
	case tetrimino.tetriminoType == Stick:
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
func (tetrimino *Tetrimino) ApplyAction(action ActionType) {
	switch {
	case action == Rotate:
		tetrimino.Rot = (tetrimino.Rot + 1) % 4
		tetrimino.Update()
	case action == Left:
		tetrimino.Pos.X--
		tetrimino.Update()
	case action == Right:
		tetrimino.Pos.X++
		tetrimino.Update()
	case action == Down:
		tetrimino.Pos.Y++
		tetrimino.Update()
	default:
		panic("----Not Defined Action---")
	}
}