package stage

import tm "github.com/tetris-CLI/tetrimino"

//StageHeight Stageの横の長さ
const StageHeight int = 20

//StageWidth Stageの縦の長さ
const StageWidth int = 10

//Line 各ブロックにミノが存在しているかどうかを表すStageのライン
type Line [StageWidth]bool

//Stage {StageHeight}行で構成されるテトリスのステージ
type Stage [StageHeight]Line

//AddBlocks Stageに，他のブロックを追加する
func (stage *Stage) AddBlocks(blockPositions tm.BlockPositions) {
	for _, position := range blockPositions {
		stage[position.Y][position.X] = true
	}
}

//Refresh Stage内の埋まっているLineを消去する
func (stage *Stage) Refresh() {
	var refreshed Stage
	index := 0
	for _, line := range stage {
		if !isFilledLine(line) {
			refreshed[index] = line
			index++
		}
	}
	*stage = refreshed
}

//isFilledLine 与えられたLineが埋まっているかどうかを返す
func isFilledLine(line Line) bool {
	for _, square := range line {
		if square == false {
			return false
		}
	}
	return true
}

//IsGameOver Stageの情報からゲームが終了しているかどうかを返す
func (stage *Stage) IsGameOver() bool {
	for _, tmpBlock := range stage[0] {
		if tmpBlock == true {
			return true
		}
	}
	return false
}
