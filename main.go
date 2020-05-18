package main

import (
	"os"

	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	store "github.com/tetris-CLI/store"
	_ "github.com/tetris-CLI/reducer"
	_ "github.com/tetris-CLI/view"
)

func init() {
	store.Store.Dispatcher.On(a.ExitGameAction, func() {
		termbox.Close()
		os.Exit(0)
	})
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	store.Store.Dispatcher.Emit(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				store.Store.Dispatcher.Emit(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				store.Store.Dispatcher.Emit(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				store.Store.Dispatcher.Emit(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				store.Store.Dispatcher.Emit(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				store.Store.Dispatcher.Emit(a.RotateTetriminoToRightAction)
			case termbox.KeySpace:
				store.Store.Dispatcher.Emit(a.HardDropTetriminoAction)
			default:
			}
		default:
		}
	}
}
