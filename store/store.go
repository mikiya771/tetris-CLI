package store

import (
	"math/rand"
	"time"

	a "github.com/tetris-CLI/action"
	"github.com/tetris-CLI/config"
	e "github.com/tetris-CLI/emitter"
	c "github.com/tetris-CLI/store/cell"
	st "github.com/tetris-CLI/store/stage"
	tm "github.com/tetris-CLI/store/tetrimino"
)

//Store Tetrisのstateを保持する型
type Store struct {
	tetrimino      tm.Tetrimino
	tetriminoQueue []tm.ShapeType
	stage          st.Stage
	UpdateNotifier e.Emitter
}

func (store *Store) popTetriminoQueue() tm.ShapeType {
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

//GetTetriminoQueue Store.tetriminoQueueを返す
func (store *Store) GetTetriminoQueue() []tm.ShapeType {
	return store.tetriminoQueue
}

//GetTetrimino Store.tetriminoを返す
func (store *Store) GetTetrimino() tm.Tetrimino {
	return store.tetrimino
}

//SetNextTetrimino Store.tetriminoQueueから1つ取り出し、tetriminoにセットする
func (store *Store) SetNextTetrimino() {
	tetrimino := tm.NewTetrimino(store.popTetriminoQueue())

	for i := 0; i < config.InvisibleStageHeight; i++ {
		tetrimino.MoveBy(0, 1)
		if store.stage.IsConflictedWith(tetrimino) {
			tetrimino.MoveBy(0, -1)
			break
		}
	}

	store.tetrimino = tetrimino
	store.UpdateNotifier.Emit(a.UpdateViewAction)
}

//SetTetrimino 引数をStore.tetriminoとしてセットする
func (store *Store) SetTetrimino(tetrimino tm.Tetrimino) {
	store.tetrimino = tetrimino
	go store.UpdateNotifier.Emit(a.UpdateViewAction)
}

//GetStage Store.stageを返す
func (store *Store) GetStage() st.Stage {
	return store.stage
}

//SetStage 引数をStore.stageとしてセットする
func (store *Store) SetStage(stage st.Stage) {
	store.stage = stage
	go store.UpdateNotifier.Emit(a.UpdateViewAction)
}

//GetStageCell Store.stage内の(x, y)の一のCellを返す
func (store *Store) GetStageCell(x, y int) c.Cell {
	return store.stage.Lines[y].Cells[x]
}

//SetStageCell cellをStore.stageの(x, y)にセットする
func (store *Store) SetStageCell(x, y int, cell c.Cell) {
	store.stage.Lines[y].Cells[x] = cell
}

//NewStore Store構造体を初期化して返す
func NewStore() Store {
	shapes := []tm.ShapeType{tm.IShape, tm.LShape, tm.JShape, tm.OShape, tm.TShape, tm.SShape, tm.ZShape}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shapes), func(i, j int) { shapes[i], shapes[j] = shapes[j], shapes[i] })
	return Store{
		tetriminoQueue: shapes,
		UpdateNotifier: e.NewEmitter(),
	}
}
