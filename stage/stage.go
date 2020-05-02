package stage

import . "github.com/tetris-CLI/tetrimino"

const MAX_HEIGHT int = 20
const MAX_WIDTH int = 10

//Line Stageの一行文のブロックの存在非存在を表している
type Line [MAX_WIDTH]bool

//Stage 20行分のラインが集まっている
type Stage [MAX_HEIGHT]Line

//AddBlocks Stageに，他のブロックを追加して足せるようになっていく
func (stage *Stage) AddBlocks(blockPositions BlockPositions) {
	for _, positions := range blockPositions {
		stage[positions.Y][positions.X] = true
	}
}

//Refresh stageに，1行分埋まっている場所を探してその部分を消去している
func (stage *Stage) Refresh() {
	var ReturnStage Stage
	IndexOfReturnStage := 0
	for _, line := range stage {
		if EvaluateLine(line) == true {
		} else {
			ReturnStage[IndexOfReturnStage] = line
			IndexOfReturnStage++
		}
	}
	*stage = ReturnStage
}

//EvaluateLine １行分をみて，その行がいっぱいであるかを判定している
func EvaluateLine(line Line) bool {
	for _, square := range line {
		if square == false {
			return false
		}
	}
	return true
}

//IsGameSet Stage情報からゲームが終了していないかを確認している
func (stage *Stage) IsGameSet() bool {
	for _, tmpBlock := range stage[0] {
		if tmpBlock == true {
			return true
		}
	}
	return false
}
