package main

import (
	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	st "github.com/tetris-CLI/stage"
	tm "github.com/tetris-CLI/tetrimino"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func draw(tetrimino tm.Tetrimino, stage st.Stage) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ESC to exit.")
	for index, line := range dispMerge(tetrimino, stage) {
		drawLine(5, index+1, line)
	}
	termbox.Flush()
}

func dispMerge(tetrimino tm.Tetrimino, stage st.Stage) []string {
	var returnMsgs []string
	mergedStage := &*&stage

	for _, mino := range tetrimino.Minos {
		mergedStage.SetMino(mino)
	}

	for _, line := range mergedStage.Lines {
		lineString := ""
		for _, cell := range line.Cells {
			if cell.IsFilled == true {
				lineString += "x"
			} else {
				lineString += "_"
			}
		}
		returnMsgs = append(returnMsgs, "|"+lineString+"|")
	}
	return returnMsgs
}

func tetris() {
	tetrimino := tm.NewTetrimino(tm.IShape)
	var stage st.Stage
	for {
		draw(tetrimino, stage)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				tetrimino.ApplyAction(a.MoveLeftAction)
			case termbox.KeyArrowRight:
				tetrimino.ApplyAction(a.MoveRightAction)
			case termbox.KeyArrowDown:
				tetrimino.ApplyAction(a.SoftDropAction)
			case termbox.KeyArrowUp:
				tetrimino.ApplyAction(a.RotateRightAction)
			default:
				draw(tetrimino, stage)
			}
		default:
			draw(tetrimino, stage)
		}
		stage.RefreshLines()
		if stage.IsGameOver() {
			break
		}
		if EvaluateTermination(tetrimino, stage) {
			for _, mino := range tetrimino.Minos {
				stage.SetMino(mino)
			}
			tetrimino = tm.NewTetrimino(tm.IShape)
		}
		draw(tetrimino, stage)
		draw(tetrimino, stage)
	}
}

// EvaluateTermination ゲームオーバーか否かを判定する
func EvaluateTermination(tetrimino tm.Tetrimino, stage st.Stage) bool {
	for _, mino := range tetrimino.Minos {
		if mino.Y >= 19 {
			return true
		}
		if stage.Lines[mino.Y+1].Cells[mino.X].IsFilled == true {
			return true
		}
	}
	return false
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	tetris()
}
