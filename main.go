package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var pizza = []string{"Enjoy your pizza with some delicious"}

var version string = "1.0.0"

// Info shows CLI info
func Info() {
	app.Name = "Thrusta"
	app.Usage = "Ping me from anywhere"
	app.Author = "Wachira"
	app.Version = version
}

// Commands defines CLI commands
func Commands() {
	app.Commands = []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"--version", "-v"},
			Usage:   "Check version",
			Action: func(c *cli.Context) {
				fmt.Println("Thrusta version " + version)
			},
		},
	}
}

func main() {
	Info()
	Commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
