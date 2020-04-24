package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/urfave/cli"
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}
func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawLine(0, 0, "press ESC to exit.")
	drawLine(2, 1, fmt.Sprintf("date: %v", time.Now()))
	termbox.Flush()
}
func tetris() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			}
		default:
			draw()
		}
	}
}
func main() {
	app := cli.NewApp()
	app.Name = "TetrisCLI"
	app.Usage = "This app serves tetris on CLI"
	app.Version = "0.0.1"
	app.Action = func(context *cli.Context) error {
		fmt.Println(context.Args().Get(0))
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()
		tetris()
		return nil
	}
	app.Run(os.Args)
}
