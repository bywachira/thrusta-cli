// go: generate goversioninfo -icon = thrusta.ico
package main

import (
	"log"
	"os"

	"github.com/tesh254/thrusta-cli/commandline"
)

func main() {
	app := commandline.SetupCLI()

	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
	}
}
