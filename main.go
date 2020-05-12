package main

import (
	"os"

	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	d "github.com/tetris-CLI/dispatcher"
	_ "github.com/tetris-CLI/reducer"
	_ "github.com/tetris-CLI/view"
)

func init() {
	d.Dispatcher.On(a.ExitGameAction, func() {
		termbox.Close()
		os.Exit(0)
	})
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	d.Dispatcher.Emit(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				d.Dispatcher.Emit(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				d.Dispatcher.Emit(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				d.Dispatcher.Emit(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				d.Dispatcher.Emit(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				d.Dispatcher.Emit(a.RotateTetriminoToRightAction)
			case termbox.KeySpace:
				d.Dispatcher.Emit(a.HardDropTetriminoAction)
			default:
			}
		default:
		}
	}
}
