package stage

import (
	l "github.com/tetris-CLI/reducer/store/line"
	m "github.com/tetris-CLI/reducer/store/mino"
	tm "github.com/tetris-CLI/reducer/store/tetrimino"

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

//IsConflictedWith StageとTetriminoが競合しているかどうかの真偽値を返す
func (stage *Stage) IsConflictedWith(tetrimino tm.Tetrimino) bool {
	for _, mino := range tetrimino.Minos {
		if mino.X < 0 || config.StageWidth-1 < mino.X {
			return true
		}

		if mino.Y < 0 || config.StageHeight-1 < mino.Y {
			return true
		}

		if stage.Lines[mino.Y].Cells[mino.X].IsFilled {
			return true
		}
	}
	return false
}

//SetMino Stageに，Minoを追加する
func (stage *Stage) SetMino(mino m.Mino) {
	stage.Lines[mino.Y].Cells[mino.X].IsFilled = true
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
