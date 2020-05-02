package tetriStage

import . "github.com/tetris-CLI/tetrimino"

const MAX_HEIGHT int = 20
const MAX_WIDTH int = 10

//Stageの一行文のブロックの存在非存在を表している
type Line [MAX_WIDTH]bool

//20行分のラインが集まっている
type TetriStage [MAX_HEIGHT]Line

//TetriStageに，他のブロックを追加して足せるようになっていく
func (tetriStage *TetriStage) AddBlocks(blockPositions BlockPositions) {
	for _, positions := range blockPositions {
		tetriStage[positions.Y][positions.X] = true
	}
}

//tetriStageに，1行分埋まっている場所を探してその部分を消去している
func (tetriStage *TetriStage) Refresh() {
	var ReturnStage TetriStage
	IndexOfReturnStage := 0
	for _, line := range tetriStage {
		if EvaluateLine(line) == true {
		} else {
			ReturnStage[IndexOfReturnStage] = line
			IndexOfReturnStage++
		}
	}
	*tetriStage = ReturnStage
}

//１行分をみて，その行がいっぱいであるかを判定している
func EvaluateLine(line Line) bool {
	for _, square := range line {
		if square == false {
			return false
		}
	}
	return true
}

//TetriStage情報からゲームが終了していないかを確認している
func (tetriStage *TetriStage) IsGameSet() bool {
	for _, tmpBlock := range tetriStage[0] {
		if tmpBlock == true {
			return true
		}
	}
	return false
}
