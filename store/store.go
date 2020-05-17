package store

import (
	"math/rand"
	"time"

	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
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

func (store *storeType) popTetriminoQueue() tm.ShapeType {
	if len(store.tetriminoQueue) <= 5 {
		shapes := []tm.ShapeType{tm.IShape, tm.LShape, tm.JShape, tm.OShape, tm.TShape, tm.SShape, tm.ZShape}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(shapes), func(i, j int) { shapes[i], shapes[j] = shapes[j], shapes[i] })
		store.tetriminoQueue = append(store.tetriminoQueue, shapes...)
	}

	shape := store.tetriminoQueue[0]
	store.tetriminoQueue = store.tetriminoQueue[1:]
	return shape
}

func (store *storeType) GetTetriminoQueue() []tm.ShapeType {
	return store.tetriminoQueue
}

func (store *storeType) GetTetrimino() tm.Tetrimino {
	return store.tetrimino
}

func (store *storeType) SetNextTetrimino() {
	tetrimino := tm.NewTetrimino(store.popTetriminoQueue())

	for i := 0; i < config.InvisibleStageHeight; i++ {
		tetrimino.MoveBy(0, 1)
		if store.stage.IsConflictedWith(tetrimino) {
			tetrimino.MoveBy(0, -1)
			break
		}
	}

	store.tetrimino = tetrimino
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
