package store

import (
	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/dispatcher"
	st "github.com/tetris-CLI/stage"
	tm "github.com/tetris-CLI/tetrimino"

	config "github.com/tetris-CLI/config"
)

//storeType Tetrisのstateを保持する型
type storeType struct {
	Tetrimino tm.Tetrimino
	Stage     st.Stage
}

func (store *storeType) initializeGame() {
	store.Stage = st.NewStage()
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

func (store *storeType) setNewTetrimino() {
	store.Tetrimino = tm.NewTetrimino(tm.IShape)
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func (store *storeType) rotateTetriminoToLeft() {
	//TODO: implement rotate left behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func (store *storeType) rotateTetriminoToRight() {
	//TODO: implement rotate right behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
}

func (store *storeType) moveTetriminoToLeft() {
	clone := store.Tetrimino.Clone()
	clone.MoveBy(-1, 0)
	if !store.Stage.IsConflictedWith(clone) {
		store.Tetrimino.MoveBy(-1, 0)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func (store *storeType) moveTetriminoToRight() {
	clone := store.Tetrimino.Clone()
	clone.MoveBy(1, 0)
	if !store.Stage.IsConflictedWith(clone) {
		store.Tetrimino.MoveBy(1, 0)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func (store *storeType) softDropTetrimino() {
	clone := store.Tetrimino.Clone()
	clone.MoveBy(0, 1)
	if !store.Stage.IsConflictedWith(clone) {
		store.Tetrimino.MoveBy(0, 1)
		dispatcher.Dispatch(a.UpdateTetriminoAction)
	}
}

func (store *storeType) hardDropTetrimino() {
	//TODO: implement hard drop behavior
	dispatcher.Dispatch(a.UpdateTetriminoAction)
	dispatcher.Dispatch(a.FixTetriminoToStageAction)
}

func (store *storeType) updateTetrimino() {
	for _, mino := range store.Tetrimino.Minos {
		if mino.Y+1 >= config.StageHeight {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		} else if store.Stage.Lines[mino.Y+1].Cells[mino.X].IsFilled {
			dispatcher.Dispatch(a.FixTetriminoToStageAction)
			break
		}
	}

	dispatcher.Dispatch(a.UpdateViewAction)
}

func (store *storeType) fixTetriminoToStage() {
	for _, mino := range store.Tetrimino.Minos {
		store.Stage.SetMino(mino)
	}

	dispatcher.Dispatch(a.RefreshStageAction)

	if store.Stage.IsGameOver() {
		dispatcher.Dispatch(a.ExitGameAction)
	} else {
		dispatcher.Dispatch(a.SetNewTetriminoAction)
	}
}

func (store *storeType) refreshStage() {
	store.Stage.RefreshLines()
	dispatcher.Dispatch(a.SetNewTetriminoAction)
}

//Store プレイしているゲームに関するデータを保持するインスタンス
var Store storeType

func init() {
	Store = storeType{}

	dispatcher.Register(a.InitializeGameAction, Store.initializeGame)
	dispatcher.Register(a.SetNewTetriminoAction, Store.setNewTetrimino)
	// dispatcher.Register(a.RotateTetriminoToLeftAction, Store.rotateTetriminoToLeft)
	// dispatcher.Register(a.RotateTetriminoToRightAction, Store.rotateTetriminoToRight)
	dispatcher.Register(a.MoveTetriminoToLeftAction, Store.moveTetriminoToLeft)
	dispatcher.Register(a.MoveTetriminoToRightAction, Store.moveTetriminoToRight)
	dispatcher.Register(a.SoftDropTetriminoAction, Store.softDropTetrimino)
	// dispatcher.Register(a.HardDropTetriminoAction, Store.hardDropTetrimino)
	dispatcher.Register(a.UpdateTetriminoAction, Store.updateTetrimino)
	dispatcher.Register(a.FixTetriminoToStageAction, Store.fixTetriminoToStage)
	dispatcher.Register(a.RefreshStageAction, Store.refreshStage)
}
