package line

import config "github.com/tetris-CLI/config"

//Line 各ブロックにミノが存在しているかどうかを表すStageのライン
type Line [config.StageWidth]bool

//IsFilledLine 与えられたLineが埋まっているかどうかを返す
func (line *Line) IsFilledLine() bool {
	for _, square := range line {
		if square == false {
			return false
		}
	}
	return true
}
