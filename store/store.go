package store

import (
	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/dispatcher"
	st "github.com/tetris-CLI/stage"
	tm "github.com/tetris-CLI/tetrimino"

	// v "github.com/tetris-CLI/view"
	config "github.com/tetris-CLI/config"
)

//Store Tetrisのstateを保持する型
type Store struct {
	Tetrimino tm.Tetrimino
	Stage     st.Stage
}

//NewStore Storeインスタンスを初期化して返す
func NewStore() Store {
	store := Store{
		Stage: st.NewStage(),
	}

	dispatcher.Register(a.InitializeGameAction, store.initializeGame)
	dispatcher.Register(a.SetNewTetriminoAction, store.setNewTetrimino)
	// dispatcher.Register(a.RotateTetriminoToLeftAction, store.rotateTetriminoToLeft)
	// dispatcher.Register(a.RotateTetriminoToRightAction, store.rotateTetriminoToRight)
	dispatcher.Register(a.MoveTetriminoToLeftAction, store.moveTetriminoToLeft)
	dispatcher.Register(a.MoveTetriminoToRightAction, store.moveTetriminoToRight)
	dispatcher.Register(a.SoftDropTetriminoAction, store.softDropTetrimino)
	// dispatcher.Register(a.HardDropTetriminoAction, store.hardDropTetrimino)
	dispatcher.Register(a.UpdateTetriminoAction, store.updateTetrimino)
	dispatcher.Register(a.FixTetriminoToStageAction, store.fixTetriminoToStage)

	return store
}

func (store *Store) initializeGame() {
	store.Stage = st.NewStage()
	store.updateView()
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func (store *Store) setNewTetrimino() {
	store.Tetrimino = tm.NewTetrimino(tm.IShape)
	store.updateView()
}

func (store *Store) rotateTetriminoToLeft() {
	//TODO: implement rotate left behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	store.updateView()
}

func (store *Store) rotateTetriminoToRight() {
	//TODO: implement rotate right behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	store.updateView()
}

func (store *Store) moveTetriminoToLeft() {
	for i := 0; i < len(store.Tetrimino.Minos); i++ {
		store.Tetrimino.Minos[i].X = store.Tetrimino.Minos[i].X - 1
	}
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	store.updateView()
}

func (store *Store) moveTetriminoToRight() {
	for i := 0; i < len(store.Tetrimino.Minos); i++ {
		store.Tetrimino.Minos[i].X = store.Tetrimino.Minos[i].X + 1
	}
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	store.updateView()
}

func (store *Store) softDropTetrimino() {
	for i := 0; i < len(store.Tetrimino.Minos); i++ {
		store.Tetrimino.Minos[i].Y = store.Tetrimino.Minos[i].Y + 1
	}
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	store.updateView()
}

func (store *Store) hardDropTetrimino() {
	//TODO: implement hard drop behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	dispatcher.Dispatch(a.FixTetriminoToStageAction)
}

func (store *Store) updateTetrimino() {
	store.Tetrimino.Update()

	for _, mino := range store.Tetrimino.Minos {
		if mino.Y + 1 >= config.StageHeight {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		} else if store.Stage.Lines[mino.Y+1].Cells[mino.X].IsFilled {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		}
	}

	store.updateView()
}

func (store *Store) fixTetriminoToStage() {
	for _, mino := range store.Tetrimino.Minos {
		store.Stage.SetMino(mino)
	}
	store.updateView()
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func (store *Store) refreshStage() {
	store.Stage.RefreshLines()
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func (store *Store) updateView() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x, rune := range []rune("Press ESC to exit.") {
		termbox.SetCell(x, 0, rune, termbox.ColorDefault, termbox.ColorDefault)
	}

	drawStage(store.Stage)
	drawTetrimino(store.Tetrimino)

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
