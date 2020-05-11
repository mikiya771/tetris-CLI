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
	d.Dispatcher.Register(a.ExitGameAction, func() {
		termbox.Close()
		os.Exit(0)
	})
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	d.Dispatcher.Dispatch(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				d.Dispatcher.Dispatch(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				d.Dispatcher.Dispatch(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				d.Dispatcher.Dispatch(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				d.Dispatcher.Dispatch(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				d.Dispatcher.Dispatch(a.RotateTetriminoToRightAction)
			case termbox.KeySpace:
				d.Dispatcher.Dispatch(a.HardDropTetriminoAction)
			default:
			}
		default:
		}
	}
}
