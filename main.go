package main

import (
	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	dispatcher "github.com/tetris-CLI/dispatcher"

	// st "github.com/tetris-CLI/stage"
	s "github.com/tetris-CLI/store"
	// tm "github.com/tetris-CLI/tetrimino"
)

func tetris() {
	// tetrimino := tm.NewTetrimino(tm.IShape)
	// var stage st.Stage
	// for {
	// 	if stage.IsGameOver() {
	// 		break
	// 	}
	// 	if EvaluateTermination(tetrimino, stage) {
	// 		for _, mino := range tetrimino.Minos {
	// 			stage.SetMino(mino)
	// 		}
	// 		tetrimino = tm.NewTetrimino(tm.IShape)
	// 	}
	// }
}

// // EvaluateTermination ゲームオーバーか否かを判定する
// func EvaluateTermination(tetrimino tm.Tetrimino, stage st.Stage) bool {
// 	for _, mino := range tetrimino.Minos {
// 		if mino.Y >= 19 {
// 			return true
// 		}
// 		if stage.Lines[mino.Y+1].Cells[mino.X].IsFilled == true {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	s.NewStore()

	dispatcher.Dispatch(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				dispatcher.Dispatch(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				dispatcher.Dispatch(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				dispatcher.Dispatch(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				dispatcher.Dispatch(a.RotateTetriminoToLeftAction)
			default:
			}
		default:
		}
	}
}
