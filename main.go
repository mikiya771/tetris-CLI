package main

import (
	"github.com/nsf/termbox-go"
	. "github.com/tetris-CLI/stage"
	. "github.com/tetris-CLI/tetrimino"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)
	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func draw(tetrimino Tetrimino, stage Stage) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ESC to exit.")
	for index, line := range dispMerge(tetrimino, stage) {
		drawLine(5, index+1, line)
	}
	termbox.Flush()
}

func dispMerge(tetrimino Tetrimino, stage Stage) []string {
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
	tetrimino := NewTetrimino(I_SHAPE)
	var stage Stage
	for {
		draw(tetrimino, stage)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				tetrimino.ApplyAction(MOVE_LEFT)
			case termbox.KeyArrowRight:
				tetrimino.ApplyAction(MOVE_RIGHT)
			case termbox.KeyArrowDown:
				tetrimino.ApplyAction(SOFT_DROP)
			case termbox.KeyArrowUp:
				tetrimino.ApplyAction(ROTATE_RIGHT)
			default:
				draw(tetrimino, stage)
			}
		default:
			draw(tetrimino, stage)
		}
		stage.Refresh()
		if stage.IsGameSet() {
			break
		}
		if EvaluateTermination(tetrimino, stage) {
			tetrimino.IsTerminate = true
		}
		if tetrimino.IsTerminate {
			stage.AddBlocks(tetrimino.BlockPoss)
			tetrimino = NewTetrimino(I_SHAPE)
		}
		draw(tetrimino, stage)
		draw(tetrimino, stage)
	}
}

// EvaluateTermination ゲームオーバーか否かを判定する
func EvaluateTermination(tetrimino Tetrimino, stage Stage) bool {
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
