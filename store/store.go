package store

import (
	"math/rand"

	a "github.com/tetris-CLI/action"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
	vc "github.com/tetris-CLI/viewController"
)

//storeType Tetrisのstateを保持する型
type storeType struct {
	tetrimino      tm.Tetrimino
	tetriminoQueue []tm.ShapeType
	stage          st.Stage
}

func (store *storeType) setNewTetriminoQueue() []tm.ShapeType {
	shapes := []tm.ShapeType{tm.IShape, tm.LShape}
	rand.Shuffle(len(shapes), func(i, j int) { shapes[i], shapes[j] = shapes[j], shapes[i] })
	return shapes
}

func (store *storeType) popTetriminoQueue() tm.ShapeType {
	var shape tm.ShapeType
	switch len(store.tetriminoQueue) {
	case 0:
		store.tetriminoQueue = store.setNewTetriminoQueue()
		shape = store.tetriminoQueue[0]
		store.tetriminoQueue = store.tetriminoQueue[1:]
	case 1:
		shape = store.tetriminoQueue[0]
		store.tetriminoQueue = store.setNewTetriminoQueue()
	default:
		shape = store.tetriminoQueue[0]
		store.tetriminoQueue = store.tetriminoQueue[1:]
	}
	return shape
}
func (store *storeType) GetTetrimino() tm.Tetrimino {
	return store.tetrimino
}

func (store *storeType) SetNextTetrimino() {
	store.tetrimino = tm.NewTetrimino(store.popTetriminoQueue())
	vc.ViewEventManager.Trigger(a.UpdateViewAction)
}

func (store *storeType) SetTetrimino(tetrimino tm.Tetrimino) {
	store.tetrimino = tetrimino
	go vc.ViewEventManager.Trigger(a.UpdateViewAction)
}
func (store *storeType) GetStage() st.Stage {
	return store.stage
}

func (store *storeType) SetStage(stage st.Stage) {
	store.stage = stage
	go vc.ViewEventManager.Trigger(a.UpdateViewAction)
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
