package main

import (
	"github.com/nsf/termbox-go"
	ts "github.com/tetris-CLI/tetriStage"
	tm "github.com/tetris-CLI/tetrimino"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)
	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}
func draw(tetrimino tm.Tetrimino, tetriStage ts.TetriStage) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "Press ESC to exit.")
	for index, line := range dispMerge(tetrimino, tetriStage) {
		drawLine(5, index+1, line)
	}
	termbox.Flush()
}
func dispMerge(tetrimino tm.Tetrimino, tetriStage ts.TetriStage) []string {
	var returnMsgs []string
	tmpStage := tetriStage
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
	tetrimino := tm.NewTetrimino(tm.I_SHAPE)
	var tetriStage ts.TetriStage
	for {
		draw(tetrimino, tetriStage)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				tetrimino.ApplyAction(tm.Left)
			case termbox.KeyArrowRight:
				tetrimino.ApplyAction(tm.Right)
			case termbox.KeyArrowDown:
				tetrimino.ApplyAction(tm.Down)
			case termbox.KeyArrowUp:
				tetrimino.ApplyAction(tm.Rotate)
			default:
				draw(tetrimino, tetriStage)
			}
		default:
			draw(tetrimino, tetriStage)
		}
		tetriStage.Refresh()
		if tetriStage.IsGameSet() {
			break
		}
		if EvaluateTermination(tetrimino, tetriStage) {
			tetrimino.IsTerminate = true
		}
		if tetrimino.IsTerminate {
			tetriStage.AddBlocks(tetrimino.BlockPoss)
			tetrimino = tm.NewTetrimino(tm.I_SHAPE)
		}
		draw(tetrimino, tetriStage)
		draw(tetrimino, tetriStage)
	}
}
func EvaluateTermination(tetriPiece tm.Tetrimino, tetriStage ts.TetriStage) bool {
	for _, blocks := range tetriPiece.BlockPoss {
		if blocks.Y >= 19 {
			return true
		}
		if tetriStage[blocks.Y+1][blocks.X] == true {
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
