package main

import (
	"log"
	"os"
	"sort"

	"github.com/tesh254/thrusta/commandline"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "login",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					email := c.Args().Get(0)
					commands := commandline.CLI{}

					commands.LoginAction(email)
					return nil
				},
			},
			// {
			// 	Name: "run",
			// 	Aliases: []string{"r"},
			// 	Usage: "run a test background process",
			// 	Action: func(c *cli.Context) error {
			// 		cron := cron.CronMethods{}

			// 		go cron.Test()
			// 		return nil
			// 	},
			// },
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
