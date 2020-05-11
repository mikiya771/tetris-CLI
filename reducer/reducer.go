package reducer

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	d "github.com/tetris-CLI/dispatcher"
	store "github.com/tetris-CLI/store"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
)

func initializeGame() {
	store.Store.SetStage(st.NewStage())
	d.Dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func setNewTetrimino() {
	store.Store.SetTetrimino(tm.NewTetrimino(store.Store.PopRoundOfShape()))
	d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func rotateTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToLeft()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func rotateTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToRight()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func softDropTetrimino() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func hardDropTetrimino() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()

	for {
		clone.MoveBy(0, 1)
		if stage.IsConflictedWith(clone) {
			clone.MoveBy(0, -1)
			break
		}
	}
	store.Store.SetTetrimino(clone)
	d.Dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func updateTetrimino() {
	for _, mino := range store.Store.GetTetrimino().Minos {
		if mino.Y+1 >= config.StageHeight {
			d.Dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		} else if store.Store.GetStage().Lines[mino.Y+1].Cells[mino.X].IsFilled {
			d.Dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		}
	}
}

func fixTetriminoToStage() {
	tetrimino := store.Store.GetTetrimino()
	for _, mino := range tetrimino.Minos {
		store.Store.SetStageCell(mino.X, mino.Y, c.Cell{IsFilled: true})
	}

	d.Dispatcher.Dispatch(a.RefreshStageAction)

	stage := store.Store.GetStage()
	if stage.IsGameOver() {
		d.Dispatcher.Dispatch(a.ExitGameAction)
	} else {
		d.Dispatcher.Dispatch(a.SetNewTetriminoAction)
	}
}

func refreshStage() {
	stage := store.Store.GetStage()
	refreshed := st.NewStage()
	index := config.StageHeight - 1
	for i := len(stage.Lines) - 1; i >= 0; i-- {
		line := stage.Lines[i]
		if !line.IsFilledLine() {
			refreshed.Lines[index].Cells = line.Cells
			index--
		}
	}
	store.Store.SetStage(refreshed)
}

func init() {
	d.Dispatcher.Register(a.InitializeGameAction, initializeGame)
	d.Dispatcher.Register(a.SetNewTetriminoAction, setNewTetrimino)
	d.Dispatcher.Register(a.RotateTetriminoToLeftAction, rotateTetriminoToLeft)
	d.Dispatcher.Register(a.RotateTetriminoToRightAction, rotateTetriminoToRight)
	d.Dispatcher.Register(a.MoveTetriminoToLeftAction, moveTetriminoToLeft)
	d.Dispatcher.Register(a.MoveTetriminoToRightAction, moveTetriminoToRight)
	d.Dispatcher.Register(a.SoftDropTetriminoAction, softDropTetrimino)
	d.Dispatcher.Register(a.HardDropTetriminoAction, hardDropTetrimino)
	d.Dispatcher.Register(a.UpdateTetriminoAction, updateTetrimino)
	d.Dispatcher.Register(a.FixTetriminoToStageAction, fixTetriminoToStage)
	d.Dispatcher.Register(a.RefreshStageAction, refreshStage)
}
