package stage

import . "github.com/tetris-CLI/tetrimino"

//StageHeight Stageの横の長さ
const StageHeight int = 20

//StageWidth Stageの縦の長さ
const StageWidth int = 10

//Line 各ブロックにミノが存在しているかどうかを表すStageのライン
type Line [StageWidth]bool

//Stage {StageHeight}行で構成されるテトリスのステージ
type Stage [StageHeight]Line

//AddBlocks Stageに，他のブロックを追加する
func (stage *Stage) AddBlocks(blockPositions BlockPositions) {
	for _, positions := range blockPositions {
		stage[positions.Y][positions.X] = true
	}
}

//Refresh Stage内の埋まっているLineを消去する
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

//EvaluateLine 与えられたLineが埋まっているかどうかを返す
func EvaluateLine(line Line) bool {
	for _, square := range line {
		if square == false {
			return false
		}
	}
	return true
}

//IsGameSet Stageの情報からゲームが終了しているかどうかを返す
func (stage *Stage) IsGameSet() bool {
	for _, tmpBlock := range stage[0] {
		if tmpBlock == true {
			return true
		}
	}
	return false
}
