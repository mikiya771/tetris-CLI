package main

import (
	"github.com/nsf/termbox-go"
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
	tmpStage := stage
	for _, blockPos := range tetrimino.BlockPoss {
		tmpStage[blockPos.Y][blockPos.X] = true
	}
	for _, line := range tmpStage {
		tmpString := "|"
		for _, block := range line {
			if block == true {
				tmpString += "x"
			} else {
				tmpString += "_"
			}
		}
		tmpString += "|"
		returnMsgs = append(returnMsgs, tmpString)
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
				tetrimino.ApplyAction(tm.MoveLeftAction)
			case termbox.KeyArrowRight:
				tetrimino.ApplyAction(tm.MoveRightAction)
			case termbox.KeyArrowDown:
				tetrimino.ApplyAction(tm.SoftDropAction)
			case termbox.KeyArrowUp:
				tetrimino.ApplyAction(tm.RotateRightAction)
			default:
				draw(tetrimino, stage)
			}
		default:
			draw(tetrimino, stage)
		}
		stage.Refresh()
		if stage.IsGameOver() {
			break
		}
		if EvaluateTermination(tetrimino, stage) {
			tetrimino.IsTerminate = true
		}
		if tetrimino.IsTerminate {
			stage.AddBlocks(tetrimino.BlockPoss)
			tetrimino = tm.NewTetrimino(tm.IShape)
		}
		draw(tetrimino, stage)
		draw(tetrimino, stage)
	}
}

// EvaluateTermination ゲームオーバーか否かを判定する
func EvaluateTermination(tetrimino tm.Tetrimino, stage st.Stage) bool {
	for _, blocks := range tetrimino.BlockPoss {
		if blocks.Y >= 19 {
			return true
		}
		if stage[blocks.Y+1][blocks.X] == true {
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
