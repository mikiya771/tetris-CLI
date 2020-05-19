package main

import (
	"os"

	g "github.com/tetris-CLI/game"
)

func main() {
	game := g.NewGame()
	game.Run()
	os.Exit(0)
}
