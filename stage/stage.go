package stage

import (
	l "github.com/tetris-CLI/line"
	tm "github.com/tetris-CLI/tetrimino"

	c "github.com/tetris-CLI/config"
)

//Stage {StageHeight}のLineで構成されるテトリスのステージ
type Stage [c.StageHeight]l.Line

//AddBlocks Stageに，他のブロックを追加する
func (stage *Stage) AddBlocks(blockPositions tm.BlockPositions) {
	for _, position := range blockPositions {
		stage[position.Y][position.X] = true
	}
}

//RefreshLines Stage内の埋まっているLineを消去する
func (stage *Stage) RefreshLines() {
	refreshed := Stage{}
	index := c.StageHeight - 1
	for i := len(stage) - 1; i >= 0; i-- {
		line := stage[i]
		if !line.IsFilledLine() {
			refreshed[index] = line
			index--
		}
	}
	*stage = refreshed
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
