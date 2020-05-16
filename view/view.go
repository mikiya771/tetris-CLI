package view

import (
	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	config "github.com/tetris-CLI/config"
	debug "github.com/tetris-CLI/debug"
	s "github.com/tetris-CLI/store"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
	vc "github.com/tetris-CLI/viewController"
)

//UpdateView Tetrisのプレイ画面を描画する
func UpdateView() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x, rune := range []rune("Press ESC to exit.") {
		termbox.SetCell(x, 0, rune, termbox.ColorDefault, termbox.ColorDefault)
	}

	drawStage(s.Store.GetStage())
	drawTetrimino(s.Store.GetTetrimino())
	drawTetriminoDropPreview(s.Store.GetStage(), s.Store.GetTetrimino())

	if config.Debug {
		drawDebugLogs(debug.GetDebugLogs())
	}

	termbox.Flush()
}

func drawStage(stage st.Stage) {
	for y, line := range stage.Lines {
		termbox.SetCell(0, y+1, '|', termbox.ColorDefault, termbox.ColorDefault)

		for x, cell := range line.Cells {
			if cell.IsFilled {
				termbox.SetCell(x+1, y+1, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(x+1, y+1, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}

		termbox.SetCell(len(line.Cells)+1, y+1, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCell(0, config.StageHeight+1, '+', termbox.ColorDefault, termbox.ColorDefault)
	for i := 0; i < config.StageWidth; i++ {
		termbox.SetCell(i+1, config.StageHeight+1, '-', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(config.StageWidth+1, config.StageHeight+1, '+', termbox.ColorDefault, termbox.ColorDefault)
}

var tetriminoColor = map[tm.ShapeType]termbox.Attribute{
	tm.IShape: termbox.ColorCyan,
	tm.LShape: termbox.ColorMagenta,
	tm.JShape: termbox.ColorBlue,
	tm.OShape: termbox.ColorYellow,
	tm.TShape: termbox.ColorBlack,
	tm.SShape: termbox.ColorGreen,
	tm.ZShape: termbox.ColorRed,
}

func drawTetrimino(tetrimino tm.Tetrimino) {
	for _, mino := range tetrimino.Minos {
		termbox.SetCell(mino.X+1, mino.Y+1, ' ', termbox.ColorDefault, tetriminoColor[tetrimino.Shape])
	}
}

func drawTetriminoDropPreview(stage st.Stage, tetrimino tm.Tetrimino) {
	clone := tetrimino.Clone()

	for {
		clone.MoveBy(0, 1)
		if stage.IsConflictedWith(clone) {
			clone.MoveBy(0, -1)
			break
		}
	}

	for _, mino := range clone.Minos {
		termbox.SetCell(mino.X+1, mino.Y+1, '·', tetriminoColor[tetrimino.Shape], termbox.ColorDefault)
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
	vc.ViewEventManager.On(a.UpdateViewAction, UpdateView)
}
