package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "brizoctl"
	app.Usage = "CLI client for Brizo"
	app.Description = "brizoctl enables users to manage Brizo deployments from the command line"
	app.Version = "0.1.0"

	// startup
	app.Before = func(c *cli.Context) error {
		err := parseConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	}

	// global flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Value:       "~/.brizo.json",
			Usage:       "Path to brizo config file",
			Destination: &ConfigPath,
		},
	}

	// commands
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "list",
			UsageText:   "list [TYPE]",
			Usage:       "retrieve list of specified resources",
			Subcommands: listCommands,
		},
		cli.Command{
			Name:        "get",
			UsageText:   "get [TYPE] [UUID]",
			Usage:       "retrieve details of specified resource",
			Subcommands: getCommands,
		},
	}

	// start your engines!
	app.Run(os.Args)
}
