package store

import (
	a "github.com/tetris-CLI/action"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
	vc "github.com/tetris-CLI/viewController"
)

//storeType Tetrisのstateを保持する型
type storeType struct {
	tetrimino tm.Tetrimino
	stage     st.Stage
}

func (store *storeType) GetTetrimino() tm.Tetrimino {
	return store.tetrimino
}

func (store *storeType) SetTetrimino(tetrimino tm.Tetrimino) {
	store.tetrimino = tetrimino
	vc.Dispatch(a.UpdateViewAction)
}

func (store *storeType) GetStage() st.Stage {
	return store.stage
}

func (store *storeType) SetStage(stage st.Stage) {
	store.stage = stage
	vc.Dispatch(a.UpdateViewAction)
}

func (store *storeType) GetStageCell(x, y int) c.Cell {
	return store.stage.Lines[y].Cells[x]
}

func (store *storeType) SetStageCell(x, y int, cell c.Cell) {
	store.stage.Lines[y].Cells[x] = cell
}

//Store プレイしているゲームに関するデータを保持するインスタンス
var Store storeType

func init() {
	Store = storeType{}
}
