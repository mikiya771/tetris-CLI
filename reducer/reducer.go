package reducer

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	e "github.com/tetris-CLI/emitter"
	s "github.com/tetris-CLI/store"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
)

type Reducer struct {
	store      *s.Store
	dispatcher e.Emitter
}

func NewReducer(store *s.Store) Reducer {
	return Reducer{
		store: store,
	}
}

func (reducer *Reducer) Register(emitter e.Emitter) {
	reducer.dispatcher = emitter
	emitter.On(a.InitializeGameAction, reducer.initializeGame)
	emitter.On(a.SetNewTetriminoAction, reducer.setNextTetrimino)
	emitter.On(a.RotateTetriminoToLeftAction, reducer.rotateTetriminoToLeft)
	emitter.On(a.RotateTetriminoToRightAction, reducer.rotateTetriminoToRight)
	emitter.On(a.MoveTetriminoToLeftAction, reducer.moveTetriminoToLeft)
	emitter.On(a.MoveTetriminoToRightAction, reducer.moveTetriminoToRight)
	emitter.On(a.SoftDropTetriminoAction, reducer.softDropTetrimino)
	emitter.On(a.HardDropTetriminoAction, reducer.hardDropTetrimino)
	emitter.On(a.UpdateTetriminoAction, reducer.updateTetrimino)
	emitter.On(a.FixTetriminoToStageAction, reducer.fixTetriminoToStage)
	emitter.On(a.RefreshStageAction, reducer.refreshStage)
}

func (reducer Reducer) initializeGame() {
	reducer.store.SetStage(st.NewStage())
	reducer.dispatcher.Emit(a.SetNewTetriminoAction)
}

func (reducer Reducer) setNextTetrimino() {
	reducer.store.SetNextTetrimino()

	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()

	if stage.IsConflictedWith(tetrimino) {
		reducer.dispatcher.Emit(a.ExitGameAction)
	} else {
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) rotateTetriminoToLeft() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToLeft()

	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) rotateTetriminoToRight() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToRight()

	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) moveTetriminoToLeft() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) moveTetriminoToRight() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) softDropTetrimino() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.UpdateTetriminoAction)
	}
}

func (reducer Reducer) hardDropTetrimino() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()

	for {
		clone.MoveBy(0, 1)
		if stage.IsConflictedWith(clone) {
			clone.MoveBy(0, -1)
			break
		}
	}
	reducer.store.SetTetrimino(clone)
	reducer.dispatcher.Emit(a.UpdateTetriminoAction)
}

func (reducer Reducer) updateTetrimino() {
	for _, mino := range reducer.store.GetTetrimino().Minos {
		if mino.Y+1 >= config.StageHeight {
			reducer.dispatcher.Emit(a.FixTetriminoToStageAction)
			break
		} else if reducer.store.GetStage().Lines[mino.Y+1].Cells[mino.X].IsFilled {
			reducer.dispatcher.Emit(a.FixTetriminoToStageAction)
			break
		}
	}
}

func (reducer Reducer) fixTetriminoToStage() {
	tetrimino := reducer.store.GetTetrimino()

	putPositionY := 0
	for _, mino := range tetrimino.Minos {
		if putPositionY < mino.Y {
			putPositionY = mino.Y
		}
	}

	if putPositionY < config.InvisibleStageHeight {
		reducer.dispatcher.Emit(a.ExitGameAction)
		return
	}

	for _, mino := range tetrimino.Minos {
		reducer.store.SetStageCell(mino.X, mino.Y, c.Cell{IsFilled: true})
	}

	reducer.dispatcher.Emit(a.RefreshStageAction)
	reducer.dispatcher.Emit(a.SetNewTetriminoAction)
}

func (reducer Reducer) refreshStage() {
	stage := reducer.store.GetStage()
	refreshed := st.NewStage()
	index := config.StageHeight - 1
	for i := len(stage.Lines) - 1; i >= 0; i-- {
		line := stage.Lines[i]
		if !line.IsFilledLine() {
			refreshed.Lines[index].Cells = line.Cells
			index--
		}
	}
	reducer.store.SetStage(refreshed)
}
