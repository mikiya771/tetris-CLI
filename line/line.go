package line

import (
	c "github.com/tetris-CLI/cell"
	config "github.com/tetris-CLI/config"
)

//Line 複数のCellによって構成される横一列のライン
type Line struct {
	Cells [config.StageWidth]c.Cell
}

//NewLine Lineインスタンスを初期化して返す
func NewLine() Line {
	line := Line{}
	for i := 0; i < len(line.Cells); i++ {
		line.Cells[i] = c.NewCell(false)
	}
	return line
}

//IsFilledLine 与えられたLineが埋まっているかどうかを返す
func (line *Line) IsFilledLine() bool {
	for _, cell := range line.Cells {
		if cell.IsFilled == false {
			return false
		}
	}
	return true
}
