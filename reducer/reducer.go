package reducer

import (
	"time"

	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	e "github.com/tetris-CLI/emitter"
	s "github.com/tetris-CLI/store"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
)

//Reducer dispatcherによって発火されたActionに対する変更をstoreに施す
type Reducer struct {
	store      *s.Store
	timer      *time.Timer
	dispatcher e.Emitter
}

//NewReducer Reducer構造体を初期化して返す
func NewReducer(store *s.Store) Reducer {
	return Reducer{
		store: store,
	}
}

//Register Reducerにdispatcherを登録する
func (reducer *Reducer) Register(emitter e.Emitter) {
	reducer.dispatcher = emitter
	emitter.On(a.InitializeGameAction, reducer.initializeGame)
	emitter.On(a.StartTimerAction, reducer.startTimer)
	emitter.On(a.StopTimerAction, reducer.stopTimer)
	emitter.On(a.ResetTimerAction, reducer.resetTimer)
	emitter.On(a.SetNewTetriminoAction, reducer.setNextTetrimino)
	emitter.On(a.RotateTetriminoToLeftAction, reducer.rotateTetriminoToLeft)
	emitter.On(a.RotateTetriminoToRightAction, reducer.rotateTetriminoToRight)
	emitter.On(a.MoveTetriminoToLeftAction, reducer.moveTetriminoToLeft)
	emitter.On(a.MoveTetriminoToRightAction, reducer.moveTetriminoToRight)
	emitter.On(a.SoftDropTetriminoAction, reducer.softDropTetrimino)
	emitter.On(a.HardDropTetriminoAction, reducer.hardDropTetrimino)
	emitter.On(a.FixTetriminoToStageAction, reducer.fixTetriminoToStage)
	emitter.On(a.RefreshStageAction, reducer.refreshStage)
}

func (reducer Reducer) initializeGame() {
	reducer.store.SetStage(st.NewStage())
	reducer.dispatcher.Emit(a.SetNewTetriminoAction)
	reducer.dispatcher.Emit(a.StartTimerAction)
}

func (reducer *Reducer) startTimer() {
	reducer.timer = time.NewTimer(config.AutoDropIntervalTime)
	go func() {
		<-reducer.timer.C
		reducer.dispatcher.Emit(a.SoftDropTetriminoAction)
	}()
}

func (reducer *Reducer) stopTimer() {
	if !reducer.timer.Stop() {
		<-reducer.timer.C
	}
}

func (reducer *Reducer) resetTimer() {
	if reducer.timer == nil {
		return
	}
	reducer.timer.Stop()
	reducer.timer.Reset(config.AutoDropIntervalTime)
	go func() {
		<-reducer.timer.C
		reducer.dispatcher.Emit(a.SoftDropTetriminoAction)
	}()
}

func (reducer Reducer) setNextTetrimino() {
	reducer.store.SetNextTetrimino()

	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()

	if stage.IsConflictedWith(tetrimino) {
		reducer.dispatcher.Emit(a.ExitGameAction)
	} else {
		reducer.dispatcher.Emit(a.ResetTimerAction)
	}
}

func (reducer Reducer) rotateTetriminoToLeft() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToLeft()

	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.ResetTimerAction)
	}
}

func (reducer Reducer) rotateTetriminoToRight() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.RotateToRight()

	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.ResetTimerAction)
	}
}

func (reducer Reducer) moveTetriminoToLeft() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.ResetTimerAction)
	}
}

func (reducer Reducer) moveTetriminoToRight() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !stage.IsConflictedWith(clone) {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.ResetTimerAction)
	}
}

func (reducer Reducer) softDropTetrimino() {
	stage := reducer.store.GetStage()
	tetrimino := reducer.store.GetTetrimino()
	clone := tetrimino.Clone()
	clone.MoveBy(0, 1)
	if stage.IsConflictedWith(clone) {
		reducer.dispatcher.Emit(a.FixTetriminoToStageAction)
	} else {
		reducer.store.SetTetrimino(clone)
		reducer.dispatcher.Emit(a.ResetTimerAction)
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
	reducer.dispatcher.Emit(a.FixTetriminoToStageAction)
	reducer.dispatcher.Emit(a.ResetTimerAction)
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
