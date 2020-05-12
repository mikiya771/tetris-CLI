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
	tetrimino     tm.Tetrimino
	roundOfShapes []tm.ShapeType
	stage         st.Stage
}

func (store *storeType) setRoundOfShapes() []tm.ShapeType {
	shapes := []tm.ShapeType{tm.IShape, tm.LShape}
	rand.Shuffle(len(shapes), func(i, j int) { shapes[i], shapes[j] = shapes[j], shapes[i] })
	return shapes
}

func (store *storeType) PopRoundOfShape() tm.ShapeType {
	var shape tm.ShapeType
	storeLength := len(store.roundOfShapes)
	switch {
	case storeLength == 0:
		store.roundOfShapes = store.setRoundOfShapes()
		shape = store.roundOfShapes[0]
		store.roundOfShapes = store.roundOfShapes[1:]
	case storeLength == 1:
		shape = store.roundOfShapes[0]
		store.roundOfShapes = store.setRoundOfShapes()
	default:
		shape = store.roundOfShapes[0]
		store.roundOfShapes = store.roundOfShapes[1:]
	}
	return shape
}
func (store *storeType) GetTetrimino() tm.Tetrimino {
	return store.tetrimino
}

func (store *storeType) SetTetrimino(tetrimino tm.Tetrimino) {
	store.tetrimino = tetrimino
	vc.ViewEventManager.Emit(a.UpdateViewAction)
}

func (store *storeType) GetStage() st.Stage {
	return store.stage
}

func (store *storeType) SetStage(stage st.Stage) {
	store.stage = stage
	vc.ViewEventManager.Emit(a.UpdateViewAction)
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
