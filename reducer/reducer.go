package reducer

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	dispatcher "github.com/tetris-CLI/dispatcher"
	store "github.com/tetris-CLI/reducer/store"
	c "github.com/tetris-CLI/reducer/store/cell"
	st "github.com/tetris-CLI/reducer/store/stage"
	tm "github.com/tetris-CLI/reducer/store/tetrimino"
)

func initializeGame() {
	store.Store.SetStage(st.NewStage())
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func setNewTetrimino() {
	store.Store.SetTetrimino(tm.NewTetrimino(tm.IShape))
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func rotateTetriminoToLeft() {
	//TODO: implement rotate left behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func rotateTetriminoToRight() {
	//TODO: implement rotate right behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func moveTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func softDropTetrimino() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func hardDropTetrimino() {
	//TODO: implement hard drop behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	dispatcher.Dispatch(a.FixTetriminoToStageAction)
}

func updateTetrimino() {
	for _, mino := range store.Store.GetTetrimino().Minos {
		if mino.Y+1 >= config.StageHeight {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		} else if store.Store.GetStage().Lines[mino.Y+1].Cells[mino.X].IsFilled {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		}
	}
}

func fixTetriminoToStage() {
	tetrimino := store.Store.GetTetrimino()
	for _, mino := range tetrimino.Minos {
		store.Store.SetStageCell(mino.X, mino.Y, c.Cell{IsFilled: true})
	}

	dispatcher.Dispatch(a.RefreshStageAction)

	stage := store.Store.GetStage()
	if stage.IsGameOver() {
		dispatcher.Dispatch(a.ExitGameAction)
	} else {
		dispatcher.Dispatch(a.SetNewTetriminoAction)
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

	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func init() {
	dispatcher.Register(a.InitializeGameAction, initializeGame)
	dispatcher.Register(a.SetNewTetriminoAction, setNewTetrimino)
	// dispatcher.Register(a.RotateTetriminoToLeftAction, rotateTetriminoToLeft)
	// dispatcher.Register(a.RotateTetriminoToRightAction, rotateTetriminoToRight)
	dispatcher.Register(a.MoveTetriminoToLeftAction, moveTetriminoToLeft)
	dispatcher.Register(a.MoveTetriminoToRightAction, moveTetriminoToRight)
	dispatcher.Register(a.SoftDropTetriminoAction, softDropTetrimino)
	// dispatcher.Register(a.HardDropTetriminoAction, hardDropTetrimino)
	dispatcher.Register(a.UpdateTetriminoAction, updateTetrimino)
	dispatcher.Register(a.FixTetriminoToStageAction, fixTetriminoToStage)
	dispatcher.Register(a.RefreshStageAction, refreshStage)
}
