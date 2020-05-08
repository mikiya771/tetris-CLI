package view

import (
	"github.com/nsf/termbox-go"

	a "github.com/tetris-CLI/action"
	config "github.com/tetris-CLI/config"
	debug "github.com/tetris-CLI/debug"
	dispatcher "github.com/tetris-CLI/dispatcher"
	st "github.com/tetris-CLI/stage"
	s "github.com/tetris-CLI/store"
	tm "github.com/tetris-CLI/tetrimino"
)

//UpdateView Tetrisのプレイ画面を描画する
func UpdateView() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x, rune := range []rune("Press ESC to exit.") {
		termbox.SetCell(x, 0, rune, termbox.ColorDefault, termbox.ColorDefault)
	}

	drawStage(s.Store.GetStage())
	drawTetrimino(s.Store.GetTetrimino())

	if config.Debug {
		drawDebugLogs(debug.GetDebugLogs())
	}

	termbox.Flush()
}

func drawStage(stage st.Stage) {
	for y, line := range stage.Lines {
		termbox.SetCell(0, y+1, []rune("|")[0], termbox.ColorDefault, termbox.ColorDefault)

		for x, cell := range line.Cells {
			if cell.IsFilled {
				termbox.SetCell(x+1, y+1, []rune("x")[0], termbox.ColorDefault, termbox.ColorDefault)
			} else {
				termbox.SetCell(x+1, y+1, []rune("_")[0], termbox.ColorDefault, termbox.ColorDefault)
			}
		}

		termbox.SetCell(len(line.Cells)+1, y+1, []rune("|")[0], termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawTetrimino(tetrimino tm.Tetrimino) {
	for _, mino := range tetrimino.Minos {
		termbox.SetCell(mino.X+1, mino.Y+1, []rune("x")[0], termbox.ColorDefault, termbox.ColorDefault)
	}
}

func drawDebugLogs(debugLogs []string) {
	for y, log := range debugLogs {
		for x, rune := range []rune(log) {
			termbox.SetCell(x+config.StageWidth+10, y+1, rune, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func init() {
	dispatcher.Register(a.UpdateViewAction, UpdateView)
}
