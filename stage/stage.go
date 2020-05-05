package stage

import (
	l "github.com/tetris-CLI/line"
	m "github.com/tetris-CLI/mino"

	config "github.com/tetris-CLI/config"
)

//Stage {StageHeight}のLineで構成されるテトリスのステージ
type Stage struct {
	Lines [config.StageHeight]l.Line
}

//NewStage Stageインスタンスを初期化して返す
func NewStage() Stage {
	stage := Stage{}
	for i := 0; i < len(stage.Lines); i++ {
		stage.Lines[i] = l.NewLine()
	}
	return stage
}

//SetMino Stageに，Minoを追加する
func (stage *Stage) SetMino(mino m.Mino) {
	stage.Lines[mino.Y].Cells[mino.X].IsFilled = true
}

//RefreshLines Stage内の埋まっているLineを消去する
func (stage *Stage) RefreshLines() {
	refreshed := NewStage()
	index := config.StageHeight - 1
	for i := len(stage.Lines) - 1; i >= 0; i-- {
		line := stage.Lines[i]
		if !line.IsFilledLine() {
			refreshed.Lines[index].Cells = line.Cells
			index--
		}
	}
	*stage = refreshed
}

//IsGameOver Stageの情報からゲームが終了しているかどうかを返す
func (stage *Stage) IsGameOver() bool {
	for _, cell := range stage.Lines[0].Cells {
		if cell.IsFilled == true {
			return true
		}
	}
	return false
}
