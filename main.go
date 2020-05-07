package main

import (
	"os"

	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	dispatcher "github.com/tetris-CLI/dispatcher"
	_ "github.com/tetris-CLI/store"
	_ "github.com/tetris-CLI/view"
)

func init() {
	dispatcher.Register(a.ExitGameAction, func() {
		termbox.Close()
		os.Exit(0)
	})
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	dispatcher.Dispatch(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				dispatcher.Dispatch(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				dispatcher.Dispatch(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				dispatcher.Dispatch(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				dispatcher.Dispatch(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				dispatcher.Dispatch(a.RotateTetriminoToLeftAction)
			case termbox.KeySpace:
				dispatcher.Dispatch(a.HardDropTetriminoAction)
			default:
			}
		default:
		}
	}
}
