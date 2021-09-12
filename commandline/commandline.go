package commandline

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	Backend "thrusta/backend"
	LocalTCP "thrusta/tcp"

	"github.com/urfave/cli/v2"
)

type CLI struct {
}

type FmtOperator func(output string) (n int, err error)

type Console struct {
	Success FmtOperator
	Error   FmtOperator
	Info    FmtOperator
}

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

var console Console = Console{
	Success: func(output string) (n int, err error) { return fmt.Println(string(colorGreen), output) },
	Error:   func(output string) (n int, err error) { return fmt.Println(string(colorRed), output) },
	Info:    func(output string) (n int, err error) { return fmt.Println(string(colorBlue), output) },
}

func (c *CLI) SetupCLI() *cli.App {
	app := &cli.App{
		Name:  "Thrusta",
		Usage: "Automate your scripts remotely",
		Commands: []*cli.Command{
			{
				Name:    "st",
				Aliases: []string{"s"},
				Usage:   "Start the thrusta server status process and send stats",
				Action: func(c *cli.Context) error {
					console.Info("Starting Thrusta")

					ticker := time.NewTicker(5 * time.Second)

					quit := make(chan struct{})

					// go func() {
					for {
						select {
						case <-ticker.C:
							LocalTCP.ConnectToBackend()
						case <-quit:
							ticker.Stop()
							// return
						}
					}
					// }()

					// return nil
				},
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Setup Thrusta",
				Action: func(c *cli.Context) error {
					fmt.Println("Setting up")

					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test backend connection",
				Action: func(c *cli.Context) error {
					fmt.Println("Testing connection")

					LocalTCP.ConnectToBackend()

					return nil
				},
			},
			{
				Name:    "login",
				Aliases: []string{"l"},
				Usage:   "Generate access token",
				Action: func(c *cli.Context) error {
					fmt.Println("Processing")

					if c.NArg() > 0 {
						apiKey := c.Args().Get(0)
						nodeUUID := c.Args().Get(1)

						is_okay, err := Backend.GenerateLoginToken(apiKey, nodeUUID)

						if !is_okay {
							console.Error("Status: " + strconv.Itoa(err.Status))
							console.Error("Message: " + err.Message)
						} else {
							console.Success("Status: " + strconv.Itoa(err.Status))
							console.Success("Message: Everything is fine")
						}
					}

					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
