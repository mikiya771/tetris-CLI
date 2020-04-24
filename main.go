package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "TetrisCLI"
	app.Usage = "This app serves tetris on CLI"
	app.Version = "0.0.1"
	app.Action = func(context *cli.Context) error {
		fmt.Println(context.Args().Get(0))
		return nil
	}
	app.Run(os.Args)
}
