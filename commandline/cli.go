package commandline

import (

	// "github.com/tesh254/thrusta-cli/commandline"
	"sort"

	"github.com/urfave/cli/v2"
)

// SetupCLI handles setting up of CLI commands
func SetupCLI() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "login",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					email := c.Args().Get(0)
					commands := CLI{}

					commands.LoginAction(email)
					return nil
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run a test background process",
				Action: func(c *cli.Context) error {
					InitWebsocketClient()
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
