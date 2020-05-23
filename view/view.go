package view

import (
	"errors"

	"github.com/nsf/termbox-go"
	"github.com/tetris-CLI/action"
	config "github.com/tetris-CLI/config"
	debug "github.com/tetris-CLI/debug"
	s "github.com/tetris-CLI/store"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
)

//View テトリスのゲーム画面を描画するための構造体
type View struct {
	store          *s.Store
	hasWatchTarget bool
}

//NewView View構造体を初期化して返す
func NewView() View {
	return View{
		hasWatchTarget: false,
	}
}

//Watch 変更を監視するStoreを設定する
func (view *View) Watch(store *s.Store) error {
	if view.hasWatchTarget {
		return errors.New("view already has watching target")
	}

	view.hasWatchTarget = true
	view.store = store
	store.UpdateNotifier.On(action.UpdateViewAction, view.updateView)
	return nil
}

//UpdateView Tetrisのプレイ画面を描画する
func (view View) updateView() {
	if view.hasWatchTarget == false {
		return
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawStage(view.store.GetStage())
	drawTetriminoDropPreview(view.store.GetStage(), view.store.GetTetrimino())
	drawTetrimino(view.store.GetTetrimino())
	drawTetriminoQueue(view.store.GetTetriminoQueue())

	if config.Debug {
		drawDebugLogs(debug.GetDebugLogs())
	}

	termbox.Flush()
}

const (
	stageTopLeftX = 1
	stageTopLeftY = 1

	queueTopLeftX = stageTopLeftX+config.StageWidth
	queueTopLeftY = stageTopLeftY+1

	debugLogTopLeftX = queueTopLeftX+8
	debugLogTopLeftY = stageTopLeftY
)

func drawStage(stage st.Stage) {
	for y := config.InvisibleStageHeight; y < config.StageHeight; y++ {
		termbox.SetCell(stageTopLeftX, stageTopLeftY+y-config.InvisibleStageHeight, '|', termbox.ColorDefault, termbox.ColorDefault)
		line := stage.Lines[y]

		for x, cell := range line.Cells {
			if cell.IsFilled {
				termbox.SetCell(stageTopLeftX+x+1, stageTopLeftY+y-config.InvisibleStageHeight, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(stageTopLeftX+x+1, stageTopLeftY+y-config.InvisibleStageHeight, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}

		termbox.SetCell(stageTopLeftX+len(line.Cells)+1, stageTopLeftY+y-config.InvisibleStageHeight, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCell(stageTopLeftX, stageTopLeftY+config.StageHeight-config.InvisibleStageHeight, '+', termbox.ColorDefault, termbox.ColorDefault)
	for x := 0; x < config.StageWidth; x++ {
		termbox.SetCell(stageTopLeftX+x+1, stageTopLeftY+config.StageHeight-config.InvisibleStageHeight, '-', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(stageTopLeftX+config.StageWidth+1, stageTopLeftY+config.StageHeight-config.InvisibleStageHeight, '+', termbox.ColorDefault, termbox.ColorDefault)
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
		if mino.Y < config.InvisibleStageHeight {
			continue
		}
		termbox.SetCell(stageTopLeftX+mino.X+1, stageTopLeftY+mino.Y-config.InvisibleStageHeight, ' ', termbox.ColorDefault, tetriminoColor[tetrimino.Shape])
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
		termbox.SetCell(stageTopLeftX+mino.X+1, stageTopLeftY+mino.Y-config.InvisibleStageHeight, '·', tetriminoColor[tetrimino.Shape], termbox.ColorDefault)
	}
}

func drawTetriminoQueue(tetriminoQueue []tm.ShapeType) {
	for i := 0; i < 5; i++ {
		tetrimino := tm.NewTetrimino(tetriminoQueue[i])
		for _, mino := range tetrimino.Minos {
			termbox.SetCell(queueTopLeftX+mino.X, queueTopLeftY+mino.Y+i*3, ' ', termbox.ColorDefault, tetriminoColor[tetrimino.Shape])
		}
	}
}

func drawDebugLogs(debugLogs []string) {
	for y, log := range debugLogs {
		for x, rune := range []rune(log) {
			termbox.SetCell(debugLogTopLeftX+x, debugLogTopLeftY+y, rune, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
