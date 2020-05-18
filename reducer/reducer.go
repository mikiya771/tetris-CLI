package reducer

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	store "github.com/tetris-CLI/store"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
)

func initializeGame() {
	store.Store.SetStage(st.NewStage())
	store.Store.Dispatcher.Emit(a.SetNewTetriminoAction)
}

func setNextTetrimino() {
	store.Store.SetNextTetrimino()

	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()

	if (stage.IsConflictedWith(tetrimino)) {
		store.Store.Dispatcher.Emit(a.ExitGameAction)
	} else {
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func rotateTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToLeft()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func rotateTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToRight()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func softDropTetrimino() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
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
	store.Store.Dispatcher.Emit(a.UpdateTetriminoAction)
}

func updateTetrimino() {
	for _, mino := range store.Store.GetTetrimino().Minos {
		if mino.Y+1 >= config.StageHeight {
			store.Store.Dispatcher.Emit(a.FixTetriminoToStageAction)
			break
		} else if store.Store.GetStage().Lines[mino.Y+1].Cells[mino.X].IsFilled {
			store.Store.Dispatcher.Emit(a.FixTetriminoToStageAction)
			break
		}
	}
}

func fixTetriminoToStage() {
	tetrimino := store.Store.GetTetrimino()

	putPositionY := 0
	for _, mino := range tetrimino.Minos {
		if putPositionY < mino.Y {
			putPositionY =  mino.Y
		}
	}

	if putPositionY < config.InvisibleStageHeight {
		store.Store.Dispatcher.Emit(a.ExitGameAction)
		return
	}

	for _, mino := range tetrimino.Minos {
		store.Store.SetStageCell(mino.X, mino.Y, c.Cell{IsFilled: true})
	}

	store.Store.Dispatcher.Emit(a.RefreshStageAction)
	store.Store.Dispatcher.Emit(a.SetNewTetriminoAction)
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
	store.Store.Dispatcher.On(a.InitializeGameAction, initializeGame)
	store.Store.Dispatcher.On(a.SetNewTetriminoAction, setNextTetrimino)
	store.Store.Dispatcher.On(a.RotateTetriminoToLeftAction, rotateTetriminoToLeft)
	store.Store.Dispatcher.On(a.RotateTetriminoToRightAction, rotateTetriminoToRight)
	store.Store.Dispatcher.On(a.MoveTetriminoToLeftAction, moveTetriminoToLeft)
	store.Store.Dispatcher.On(a.MoveTetriminoToRightAction, moveTetriminoToRight)
	store.Store.Dispatcher.On(a.SoftDropTetriminoAction, softDropTetrimino)
	store.Store.Dispatcher.On(a.HardDropTetriminoAction, hardDropTetrimino)
	store.Store.Dispatcher.On(a.UpdateTetriminoAction, updateTetrimino)
	store.Store.Dispatcher.On(a.FixTetriminoToStageAction, fixTetriminoToStage)
	store.Store.Dispatcher.On(a.RefreshStageAction, refreshStage)
}
