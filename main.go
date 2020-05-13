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
	go d.Dispatcher.On(a.ExitGameAction, func() {
		termbox.Close()
		os.Exit(0)
	})
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	go d.Dispatcher.Trigger(a.InitializeGameAction)
	pollKeyEvent()
}

func pollKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				go d.Dispatcher.Trigger(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				go d.Dispatcher.Trigger(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				go d.Dispatcher.Trigger(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				go d.Dispatcher.Trigger(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				go d.Dispatcher.Trigger(a.RotateTetriminoToRightAction)
			case termbox.KeySpace:
				go d.Dispatcher.Trigger(a.HardDropTetriminoAction)
			default:
			}
		default:
		}
	}
}
