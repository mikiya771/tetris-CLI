package reducer

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	d "github.com/tetris-CLI/dispatcher"
	store "github.com/tetris-CLI/store"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
)

func initializeGame() {
	store.Store.SetStage(st.NewStage())
	go d.Dispatcher.Trigger(a.SetNewTetriminoAction)
}

func setNextTetrimino() {
	store.Store.SetNextTetrimino()
	go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
}

func rotateTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToLeft()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
	}
}

func rotateTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToRight()

	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToLeft() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
	}
}

func moveTetriminoToRight() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
	}
}

func softDropTetrimino() {
	stage := store.Store.GetStage()
	tetrimino := store.Store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !stage.IsConflictedWith(clone) {
		store.Store.SetTetrimino(clone)
		go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
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
	go d.Dispatcher.Trigger(a.UpdateTetriminoAction)
}

func updateTetrimino() {
	for _, mino := range store.Store.GetTetrimino().Minos {
		if mino.Y+1 >= config.StageHeight {
			go d.Dispatcher.Trigger(a.FixTetriminoToStageAction)
			break
		} else if store.Store.GetStage().Lines[mino.Y+1].Cells[mino.X].IsFilled {
			go d.Dispatcher.Trigger(a.FixTetriminoToStageAction)
			break
		}
	}
}

func fixTetriminoToStage() {
	tetrimino := store.Store.GetTetrimino()
	for _, mino := range tetrimino.Minos {
		store.Store.SetStageCell(mino.X, mino.Y, c.Cell{IsFilled: true})
	}

	go d.Dispatcher.Trigger(a.RefreshStageAction)

	stage := store.Store.GetStage()
	if stage.IsGameOver() {
		go d.Dispatcher.Trigger(a.ExitGameAction)
	} else {
		go d.Dispatcher.Trigger(a.SetNewTetriminoAction)
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
	d.Dispatcher.On(a.InitializeGameAction, initializeGame)
	d.Dispatcher.On(a.SetNewTetriminoAction, setNextTetrimino)
	d.Dispatcher.On(a.RotateTetriminoToLeftAction, rotateTetriminoToLeft)
	d.Dispatcher.On(a.RotateTetriminoToRightAction, rotateTetriminoToRight)
	d.Dispatcher.On(a.MoveTetriminoToLeftAction, moveTetriminoToLeft)
	d.Dispatcher.On(a.MoveTetriminoToRightAction, moveTetriminoToRight)
	d.Dispatcher.On(a.SoftDropTetriminoAction, softDropTetrimino)
	d.Dispatcher.On(a.HardDropTetriminoAction, hardDropTetrimino)
	d.Dispatcher.On(a.UpdateTetriminoAction, updateTetrimino)
	d.Dispatcher.On(a.FixTetriminoToStageAction, fixTetriminoToStage)
	d.Dispatcher.On(a.RefreshStageAction, refreshStage)

}
