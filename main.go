package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "brizoctl"
	app.Description = "CLI tool for Brizo"

	app.Run(os.Args)
}
